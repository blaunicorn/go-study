package main

import (
	"fmt"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	// db := InitDB()
	common.InitDB()

	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.POST("/api/auth/register", controller.Register)

	r = CollectRoute(r)

	panic(r.Run())

}
