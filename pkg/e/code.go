package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_CERT       = 10001
	ERROR_EXIST_CERT_FAIL  = 10002
	ERROR_NOT_EXIST_CERT   = 10003
	ERROR_GET_CERTS_FAIL   = 10004
	ERROR_COUNT_CERT_FAIL  = 10005
	ERROR_ADD_CERT_FAIL    = 10006
	ERROR_EDIT_CERT_FAIL   = 10007
	ERROR_DELETE_CERT_FAIL = 10008
	ERROR_EXPORT_FILE_FAIL = 10009
	ERROR_IMPORT_FILE_FAIL = 10010

	ERROR_NOT_EXIST_USER        = 10011
	ERROR_CHECK_EXIST_USER_FAIL = 10012
	ERROR_ADD_USER_FAIL         = 10013
	ERROR_DELETE_USER_FAIL      = 10014
	ERROR_EDIT_USER_FAIL        = 10015
	ERROR_COUNT_USER_FAIL       = 10016
	ERROR_GET_USER_FAIL         = 10017
	ERROR_GET_USERS_FAIL        = 10018
	ERROR_GEN_CERT_POSTER_FAIL  = 10019
	ERROR_CERT_APPLY_DISABLED   = 10020
	ERROR_EXIST_APPLY  			= 10021
	ERROR_CHECK_EXIST_APPLY_FAIL  = 10022
	ERROR_ADD_APPLY  			= 10023
	ERROR_EXIST_APPLY_PAY	    = 10024
	ERROR_EXIST_APPLY_PID  		= 10025

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_AUTH_CHECK_USRNAME_FAIL  = 20005
	ERROR_AUTH_CHECK_PASSWORD_FAIL = 20006
	ERROR_AUTH_CHANGE_PASSWORD_FAIL = 20007

	ERROR_UPLOAD_CREATE_IMAGE_FAIL  = 30000
	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
	ERROR_UPLOAD_CREATE_FILE_FAIL   = 30004
	ERROR_UPLOAD_SAVE_FILE_FAIL     = 30005
	ERROR_UPLOAD_CHECK_FILE_FAIL    = 30006
	ERROR_UPLOAD_CHECK_FILE_FORMAT  = 30007
	ERROR_GET_USER_CERT_IMAGES  	= 30008
	ERROR_GET_USER_CERT_FILES	  	= 30009
)
