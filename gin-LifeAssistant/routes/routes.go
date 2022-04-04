package routes

import (
	v1 "gin-vue-lifeassistant/api/v1"
	"gin-vue-lifeassistant/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "It works",
			})
		})

		// user模块的路由接口,增删改查
		userRoutes := router.Group("/user")
		userRoutes.POST("/", v1.NewUserController().Create)
		userRoutes.DELETE("/:id", v1.NewUserController().Delete)
		userRoutes.DELETE("/", v1.NewUserController().DeleteList)
		userRoutes.PUT("/:id", v1.NewUserController().Update)
		userRoutes.GET("/", v1.NewUserController().List)
		userRoutes.GET("/:id", v1.NewUserController().Show)

		// category模块的路由接口
		categoryRoutes := router.Group("/category")
		categoryRoutes.POST("/", v1.NewCategoryController().Create)
		categoryRoutes.DELETE("/:id", v1.NewCategoryController().Delete)
		categoryRoutes.DELETE("/", v1.NewCategoryController().DeleteList)
		categoryRoutes.PUT("/:id", v1.NewCategoryController().Update)
		categoryRoutes.GET("/", v1.NewCategoryController().List)
		categoryRoutes.GET("/:id", v1.NewCategoryController().Show)
		// article模块的路由接口
	}
	r.Run(utils.HttpPort)
}
