package routers

import (
	"fmt"
	"net/http"
	"strings"
	"taxcas/routers/api/admin"
	"taxcas/routers/api/user"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "taxcas/docs"

	"taxcas/pkg/export"
	"taxcas/pkg/setting"
	"taxcas/pkg/upload"
	"taxcas/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	// 跨域
	r.Use(Cors())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// html
	r.LoadHTMLGlob("views/*")
	webadmin := r.Group("/admin")
	webadmin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"taxcas"})
	})
	webadmin.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title":"taxcas"})
	})

	r.StaticFS("/export", http.Dir(export.GetExportFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		// 获取字体
		apiv1.GET("/admin/fonts", admin.GetFonts)

		// 获取证书列表
		apiv1.GET("/admin/certs", admin.GetCertList)

		// 添加证书
		apiv1.POST("/admin/certs", admin.AddCert)

		// 上传图片
		apiv1.POST("/admin/images", admin.UploadImage)

		// 预览证书
		apiv1.GET("/admin/images/certs", admin.PreviewImage)

		// 获取申请状态
		apiv1.GET("/admin/applicants/certs/:certid", admin.GetApplicantList)

		// 导入审核结果
		apiv1.POST("/admin/applicants/certs/:certid", admin.ImportApplicants)

		// 导出申请状态
		apiv1.GET("/admin/files/applicants/certs/:certid", admin.ExportApplicants)
	}

	// 用户端接口
	{
		// 申请证书
		apiv1.POST("/weixin/applicants/users/:wxid", apply.Apply)

		// 查询用户
		apiv1.GET("/weixin/users/:wxid")

		// 查询证书
		apiv1.GET("/weixin/certs", admin.GetCertList)

		// 获取openid
		//apiv1.GET("/weixin/openid", weixin.WXGetOpenID)

		// 获取支付订单
		//apiv1.GET("/weixin/wxorder/:openid", weixin.WXPayUnifyOrderReq)

		// 微信服务端回调
		//apiv1.GET("weixin/wxnotify", weixin.WXPayCallback)

	}

	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("Content-Type, %s", headerStr)
		} else {
			headerStr = "Connection, Content-Length, Accept, User-Agent, Content-Type, Accept-Encoding, Origin, Referer, Accept-Language"
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			//c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}