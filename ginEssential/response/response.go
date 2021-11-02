package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
