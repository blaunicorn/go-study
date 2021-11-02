package main

import (
	"github.com/blaunicorn/oceanlearn.teach/ginessential/controller"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	// Protecting user information interface with middleware
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
