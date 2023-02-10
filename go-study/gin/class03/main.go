package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

/*
读取文件： c.FormFile("filename")
本地写文件： os.create()   c.SaveUploadFile(file,dst)
*/
func main() {
	//     单文件上传接收
	r := gin.Default()
	r.POST("/api/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// log.Panicln(file.Filename)
		// 获取文件外的其他参数
		name := c.PostForm("name")
		fmt.Println("name:", name)
		// 上传文件到指定的路径

		// gin 保存
		c.SaveUploadedFile(file, "./uploadFile/"+file.Filename)

		// 原生保存. 打开文件，然后保存到本地
		src, _ := file.Open()
		defer src.Close()
		out, _ := os.Create("./uploadFile/" + "002.png")
		io.Copy(out, src)
		defer out.Close()
		// 返回json文件
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "success",
		// 	"data": gin.H{
		// 		"name":     name,
		// 		"filename": fmt.Sprintf("'%s' uploaded.", file.Filename),
		// 	},
		// })

		// 给前端返回文件
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", "文件名.png"))
		c.File("./uploadFile/" + file.Filename)
	})

	// 多文件上传
	r.POST("/api/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		fmt.Println(files)
		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./uploadFile/"+file.Filename)
		}

	})
	r.Run(":8081")
}
