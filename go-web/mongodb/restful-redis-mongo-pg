package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []User{{1, "zhangsna"}, {2, "lisi"}, {3, "zhangwu"}}

func main() {
	r := gin.Default()
	r.GET("/user/", get)
	r.GET("/user/:id", getOne)
	r.POST("/user/", post)
	r.PATCH("/user/:id", patch)
	r.DELETE("/user/:id", deleteById)

	r.Run(":8080")
}

func get(c *gin.Context) {
	c.JSON(200, users)
}

func post(c *gin.Context) {
	var user User
	c.Bind(&user)
	if user.Id == 0 || user.Username == "" {
		c.JSON(200, "用户信息不全")
		return
	}
	users = append(users, User{Username: user.Username, Id: user.Id})
	c.JSON(200, users)
}

func patch(c *gin.Context) {
	var user User
	istrue := false
	c.Bind(&user)
	id := c.Param("id")
	for i, v := range users {
		if strconv.Itoa(v.Id) == id {
			istrue = true
			users[i].Username = user.Username
			user.Id = v.Id
		}
	}
	fmt.Println("users:", users)
	fmt.Println("type:", reflect.TypeOf(user.Id))
	// 增加用户id是否存在的判断
	if istrue == true {
		c.JSON(200, user)
	} else {
		c.JSON(200, "数据id未能检测到。")
	}

}

func deleteById(c *gin.Context) {
	id := c.Param("id")
	isTrue := false
	for i, v := range users {
		if strconv.Itoa(v.Id) == id {
			isTrue = true
			users = append(users[:i], users[i+1:]...)
		}
	}
	if isTrue == false {
		c.JSON(200, "数据id未能检测到。")
		return
	}
	c.JSON(200, users)
}

func getOne(c *gin.Context) {
	var user User
	isTrue := false
	id := c.Param("id")
	for _, v := range users {
		if strconv.Itoa(v.Id) == id {
			isTrue = true
			user = v
			break
		}
	}
	fmt.Println("users:", users)
	if isTrue == false {
		c.JSON(200, "数据id未能检测到。"+id)
		return
	}
	c.JSON(200, user)
}
