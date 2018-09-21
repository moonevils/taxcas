package weixin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/objcoding/wxpay"
	"net/http"
	//"taxcas/pkg/app"
	//"taxcas/pkg/e"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"sort"
	"strings"
	"io/ioutil"
	"encoding/json"
	"taxcas/pkg/config"
	"taxcas/models"
//	"errors"
)

func WXGetOpenID(c *gin.Context) {
	type Response struct {
                Access_token string  `json:"access_token"`
                Expires_in  int      `json:"expires_in"`
		Refresh_token string `json:"refresh_token"`
                Openid      string   `json:"openid"`
		Scope       string   `json:"scope"`
        }
	code := c.Param("code")

	url := strings.Join([]string{"https://api.weixin.qq.com/sns/oauth2/access_token",
		"?appid=", config.AppID,
		"&secret=", config.AppSecret, 
		"&code=", code,
		"&grant_type=authorization_code"}, "")

	resp, err := http.Get(url)
        if err != nil ||  resp.StatusCode != http.StatusOK{
                fmt.Println("response err!")
                fmt.Println(err)
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Println("io err!")
        }

              fmt.Println("+++++++++++++++++++++++++")
              fmt.Println(string(body))
              fmt.Println("+++++++++++++++++++++++++")

        result := Response{}
        err = json.Unmarshal(body, &result)
        if err != nil {
                fmt.Println(err)
        }

        openid := result.Openid
	fmt.Println("=======================")
	fmt.Println(result)
	fmt.Println("=======================")
	c.JSON(http.StatusOK, gin.H{"openid": openid, "token": "123456"})
}

func WXPayUnifyOrderReq(c *gin.Context) {
	//appG := app.Gin{c}
	ip := c.ClientIP()
	openid := c.Param("openid")
	certid := c.Param("certid")
	result := models.C_certs{}
	isExist,err :=  models.MgoFindOne("CertID", certid, "certs",&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "FAIL","extra":"查找数据库失败"})	
	}
	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{"msg": "FAIL","extra":"没有找到此证书"})
	}
	
	price := result.Price
	out_trade_no := UniqueId()
	fmt.Println(ip)

	client := wxpay.NewClient(wxpay.NewAccount(config.AppID, config.MchID, config.ApiKey, false))
	params := make(wxpay.Params)
	params.SetString("body", result.CertName).
		SetString("out_trade_no", out_trade_no).
		SetInt64("total_fee", int64(price)).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", config.Notify_url).
		SetString("openid", openid).
		SetString("trade_type", "JSAPI")

	p, err := client.UnifiedOrder(params)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "UnifyOder failed"})
	}

	prepay_id := p.GetString("prepay_id")
	//	appG.Response(http.StatusOK, true, e.SUCCESS, p)
	c.JSON(http.StatusOK, gin.H{"prepay_id": prepay_id})
}

type WXPayNotifyReq struct {
	Return_code    string `xml:"return_code"`
	Return_msg     string `xml:"return_msg"`
	Appid          string `xml:"appid"`
	Mch_id         string `xml:"mch_id"`
	Nonce          string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	Result_code    string `xml:"result_code"`
	Openid         string `xml:openid`
	Is_subscribe   string `xml:"is_subscribe"`
	Trade_type     string `xml:"trade_type"`
	Bank_type      string `xml:"bank_type"`
	Total_fee      string `xml:"total_fee"`
	Fee_type       string `xml:"fee_type"`
	Cash_fee       int    `xml:"cash_fee"`
	Cash_fee_Type  int    `xml:"cash_fee_type"`
	Transaction_id string `xml:"transaction_id"`
	Out_trade_no   string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	Time_end       string `xml:"time_end"`
}

type WXPayNotifyResp struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
}

func WXPayCallback(c *gin.Context) {
	//	appG := app.Gin{c}
	//	appG.Response(http.StatusOK, true, e.SUCCESS, "SUCCESS")
	var mr WXPayNotifyReq
	var ms WXPayNotifyResp
	err := c.Bind(&mr)
	if err != nil {
		ms.Return_code = "FAIL"
		ms.Return_msg = "failed to receive args"
		c.XML(http.StatusOK, ms)
		return
	}

	var reqMap map[string]interface{}
	reqMap = make(map[string]interface{}, 0)

	reqMap["return_code"] = mr.Return_code
	reqMap["return_msg"] = mr.Return_msg
	reqMap["appid"] = mr.Appid
	reqMap["mch_id"] = mr.Mch_id
	reqMap["nonce_str"] = mr.Nonce
	reqMap["result_code"] = mr.Result_code
	reqMap["openid"] = mr.Openid
	reqMap["is_subscribe"] = mr.Is_subscribe
	reqMap["trade_type"] = mr.Trade_type
	reqMap["bank_type"] = mr.Bank_type
	reqMap["total_fee"] = mr.Total_fee
	reqMap["fee_type"] = mr.Fee_type
	reqMap["cash_fee"] = mr.Cash_fee
	reqMap["cash_fee_type"] = mr.Cash_fee_Type
	reqMap["transaction_id"] = mr.Transaction_id
	reqMap["out_trade_no"] = mr.Out_trade_no
	reqMap["attach"] = mr.Attach
	reqMap["time_end"] = mr.Time_end

	if wxpayVerifySign(reqMap, mr.Sign) {
		//这里就可以更新我们的后台数据库了，其他业务逻辑同理。
		ms.Return_code = "SUCCESS"
		ms.Return_msg = "OK"
	} else {
		ms.Return_code = "FAIL"
		ms.Return_msg = "failed to verify sign, please retry!"
	}

	c.XML(http.StatusOK, ms)
}

func wxpayVerifySign(needVerifyM map[string]interface{}, sign string) bool {
	signCalc := wxpayCalcSign(needVerifyM, "API_KEY")

	//	slog.Debug("计算出来的sign: %v", signCalc)
	//	slog.Debug("微信异步通知sign: %v", sign)
	if sign == signCalc {
		fmt.Println("签名校验通过!")
		return true
	}

	fmt.Println("签名校验失败!")
	return false
}

func wxpayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}