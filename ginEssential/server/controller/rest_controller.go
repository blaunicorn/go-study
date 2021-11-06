package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// func (c PostController) Create(ctx *gin.Context) {

// }

// func (c PostController) Update(ctx *gin.Context) {

// }

// func (c PostController) Show(ctx *gin.Context) {

// }

// func (c PostController) Delete(ctx *gin.Context) {

// }
