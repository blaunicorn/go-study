package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
路由分组

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.post("/login", login)
	v1.post("submit", submit)
*/
func main() {
	// 创建路由器
	router := gin.Default()
	v1 := router.Group("/api").Use(middel(), middel2())
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
