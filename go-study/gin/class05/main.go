package main

import (
	"class05/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
gin的日志工具
*/
func main() {
	// gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式

	// 创建路由器
	router := gin.Default()

	v1 := router.Group("/api").Use(middel(), middleware.LoggerToFile(), middel2())
	// v1 := router.Group("/api")
	v1.GET("/test", test)
	v1.POST("/login", login)
	v1.POST("/submit", submit)

	router.Run(":8081")
}

func test(c *gin.Context) {
	fmt.Println("It's working.")
	c.JSON(http.StatusOK, "it's working.")
}
func login(c *gin.Context) {

}

func submit(c *gin.Context) {

}

func middel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before")
		ctx.Next()
		fmt.Println("after")
	}
}

func middel2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before2")
		ctx.Next()
		fmt.Println("after2")
	}
}
