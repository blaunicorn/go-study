package main

import (
	"github.com/blaunicorn/oceanlearn.teach/ginessential/controller"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	// r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	// Protecting user information interface with middleware
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("/", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	// categoryRoutes.PATCH("/:id", categoryController.Delete)

	postRoutes := r.Group("/posts")
	// 路由组增加中间件
	// postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("/", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.DELETE("/list", postController.PageList)

	return r
}
