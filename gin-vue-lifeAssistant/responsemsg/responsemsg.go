package responsemsg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	ERROR   = 500

	// code=1000 .. 用户模块错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	// code=2000 ...文章模块的错误

	// code=3000 ...分类模块的错误
)

var responseMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	// code=1000 .. 用户模块错误
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:    "token不存在",
	ERROR_TOKEN_RUNTIME:  "token已过期",
	ERROR_TOKEN_WRONG:    "token不正确",
	ERROR_TOKEN_TYPE:     "token格式错误",
}

func GetResponseMsg(code int) string {
	return responseMsg[code]
}

// {
// 	code: 20001,
// 	data: XXX,
// 	msg: XXX
// }

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// 业务成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//业务失败
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
