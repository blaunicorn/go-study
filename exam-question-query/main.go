package main

import (
	"exam-question-query/cache"
	"exam-question-query/model"
	"exam-question-query/routes"
	"exam-question-query/utils"
)

// @title ToDoList API
// @version 0.0.1
// @description This is a sample Server pets
// @name FanOne
// @BasePath /api/v1
func main() {
	// http://localhost:3000/swagger/index.html
	//从配置文件读入配置
	utils.Init()
	model.InitDB()
	cache.Redis()
	//转载路由 swag init -g common.go
	routes.InitRouter()

}
