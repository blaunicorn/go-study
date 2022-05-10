package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// 原生
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

// 采用cors包
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowedMethods:  []string{"*"},
			AllowedHeaders:  []string{"*"},
			ExposedHeaders:  []string{"Content-Length", "Authorization"},
			// AllowCredentials: true,
			// AllowOriginFunc: func(origin string) bool {
			// 	return origin == "https://github.com"
			// },
			MaxAge: 12 * time.Hour,
		})
	}
}
