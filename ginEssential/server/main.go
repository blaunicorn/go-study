package main

import (
	"fmt"
	"os"
	"time"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//	 read config
	InitConfig()

	// db := InitDB()
	common.InitDB()

	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	// r.POST("/api/auth/register", controller.Register)

	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
