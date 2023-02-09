package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 动态参数路由 :id
	r.GET("/api/uri/:id", func(c *gin.Context) {
		id := c.Param("id")
		// user := c.Query("user")
		user := c.DefaultQuery("user", "default_user")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"message": "success",
			"id":      id,
			"user":    user,
			"pwd":     pwd,
		})
	})

	// post  form方式传参
	r.POST("/api/path_for_form", func(c *gin.Context) {

		id := c.PostForm("id")
		user := c.DefaultPostForm("user", "default_user")
		password := c.PostForm("password")
		fmt.Println("id:", id)
		c.JSON(200, gin.H{
			"message":  "success",
			"id":       id,
			"user":     user,
			"password": password,
		})
	})

	// post  json方式接收参数 方式一
	r.POST("/api/path_for_json", func(c *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"id":       json["id"],
			"user":     json["user"],
			"password": json["password"],
		})

	})

	// post  json方式接收参数 方式二 bind模式MustBind : Bind BindJSON BindXML BindQuery BindYAML;  ShouldBind ...（表单验证，自定义验证）
	type User struct {
		Id       int    `json:"id" form:"id" uri:"id_uri" form:"id_form" binding:"required`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	r.POST("/api/path_for_json2", func(c *gin.Context) {
		json := User{}

		// 方式一
		// c.BindJSON(&json)
		//方式二
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "fail",
				"data":    err.Error(),
			})
			return
		}

		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			// "data": gin.H{
			// "id":       json.Id,
			// "name":     json.Name,
			// "password": json.Password,

			// },
			"data": json,
		})
	})

	// 默认端口8080
	r.Run(":8081")
}

/*
修改端口号
r.Run(":8081")
GET
POST
DELETE
PUT

*/
