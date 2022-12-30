package routes

import (
	"exam-question-query/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	// r.Use(middleware.Cors())
	// r.Use(middleware.Logger())
	// r.Use(gin.Recovery())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

	}

	r.Run(utils.HttpPort)
}
