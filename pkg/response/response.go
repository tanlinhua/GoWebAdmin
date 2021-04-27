package response

// const (
// 	SUCCSE = 1
// 	ERROR  = 0

// 	ERROR_USERNAME_USED  = 1001
// 	ERROR_PASSWORD_WRONG = 1002
// 	ERROR_USER_NOT_EXIST = 1003
// 	ERROR_TOKEN_EXIST    = 1004
// 	ERROR_TOKEN_ERROR    = 1005
// 	ERROR_USER_NO_RIGHT  = 1008
// )

// var statusMsg = map[int]string{
// 	SUCCSE:               "SUCCESS",
// 	ERROR:                "FAIL",
// 	ERROR_USERNAME_USED:  "用户名已存在！",
// 	ERROR_PASSWORD_WRONG: "密码错误",
// 	ERROR_USER_NOT_EXIST: "用户不存在",
// 	ERROR_TOKEN_EXIST:    "TOKEN不存在",
// 	ERROR_TOKEN_ERROR:    "TOKEN_ERROR",
// 	ERROR_USER_NO_RIGHT:  "权限不足",
// }

// func getMsgContent(status int) string {
// 	return statusMsg[status]
// }

// func ResponseMeta(status int) map[string]interface{} {
// 	var meta = make(map[string]interface{})
// 	meta["code"] = status
// 	meta["msg"] = getMsgContent(status)
// 	return meta
// }

func Success(msg string, count int, data interface{}) map[string]interface{} {
	var meta = make(map[string]interface{})
	meta["code"] = 0
	meta["msg"] = msg
	meta["count"] = count
	meta["data"] = data
	return meta
}

func Error(msg string, count int, data interface{}) map[string]interface{} {
	var meta = make(map[string]interface{})
	meta["code"] = 1
	meta["msg"] = msg
	meta["count"] = count
	meta["data"] = data
	return meta
}
