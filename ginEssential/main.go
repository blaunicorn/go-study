package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/auth/register", func(ctx *gin.Context) {
		// get parameters
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("teleForm")
		password := ctx.PostForm("password")

		// data validation
		// gin.H == map[string]interface{}
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "Mobile phone number must be 11 digits",
			})
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "The password cannot be less than 6 digits",
			})
		}
		//determine whether the mobile-phone nember exists

		// create user

		// return results
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	panic(r.Run())

}
