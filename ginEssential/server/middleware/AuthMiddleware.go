package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get authorization header object
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions1",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions2",
			})
			ctx.Abort()
			return
		}

		// After verification, obtain the userid in claim
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// User information judgment
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Insufficient permissions",
			})
			ctx.Abort()
			return
		}

		// If the user exists, write the user's information to the context
		ctx.Set("user", user)
		ctx.Next()
	}
}
