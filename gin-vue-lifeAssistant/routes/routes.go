package routes

import (
	v1 "gin-vue-lifeassistant/api/v1"
	"gin-vue-lifeassistant/middleware"
	"gin-vue-lifeassistant/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	router := r.Group("api/v1")
	{
		router.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It works",
			})
		})

		// 用户登录、登出、获取自身信息、注册、修改密码
		router.POST("/login", v1.Login)
		router.Use(middleware.TokenJwt())
		router.POST("/upload", v1.Upload)

		// user模块的路由接口,增删改查
		userRoutes := router.Group("/user")
		// userRoutes.Use(middleware.TokenJwt())
		// Protecting user information interface with middleware
		// r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
		userRoutes.POST("/", v1.NewUserController().Create)
		userRoutes.DELETE("/:id", v1.NewUserController().Delete)
		userRoutes.DELETE("/", v1.NewUserController().DeleteList)
		userRoutes.PUT("/:id", v1.NewUserController().Update)
		userRoutes.GET("/list", v1.NewUserController().List)
		userRoutes.GET("/:id", v1.NewUserController().Show)

		// category模块的路由接口
		categoryRoutes := router.Group("/category")
		userRoutes.Use(middleware.TokenJwt())
		categoryRoutes.POST("/", v1.NewCategoryController().Create)
		categoryRoutes.DELETE("/:id", v1.NewCategoryController().Delete)
		categoryRoutes.DELETE("/", v1.NewCategoryController().DeleteList)
		categoryRoutes.PUT("/:id", v1.NewCategoryController().Update)
		categoryRoutes.GET("/list", v1.NewCategoryController().List)
		categoryRoutes.GET("/:id", v1.NewCategoryController().Show)

		// article模块的路由接口
		articleRoutes := router.Group("/article")
		articleRoutes.POST("/", v1.NewArticleController().Create)
		articleRoutes.DELETE("/:id", v1.NewArticleController().Delete)
		articleRoutes.DELETE("/", v1.NewArticleController().DeleteList)
		articleRoutes.PUT("/:id", v1.NewArticleController().Update)
		articleRoutes.GET("/list", v1.NewArticleController().List)
		articleRoutes.GET("/:id", v1.NewArticleController().Show)
	}
	r.Run(utils.HttpPort)
}
