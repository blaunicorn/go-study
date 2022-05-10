package main

import (
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}
