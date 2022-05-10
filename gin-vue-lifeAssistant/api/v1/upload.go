package v1

import (
	"fmt"
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/responsemsg"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	fmt.Println(fileHeader.Size, fileHeader.Filename)
	url, err := model.UploadFile(file, fileSize)

	if url != "" {
		responsemsg.Success(ctx, gin.H{"data": gin.H{"url": url}}, "upload success")
		return
	}
	if err != nil {
		responsemsg.Fail(ctx, nil, err.Error())
		return
	}
}
