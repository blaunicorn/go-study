package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	// 7. gorm使用
	_ "github.com/go-sql-driver/mysql"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name" binding:"required"`
	Age  int    `json:"age" uri:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex" `
}

func mustBig(fl validator.FieldLevel) bool {
	// 反射断言
	fmt.Println(fl.Field().Interface().(int))
	if fl.Field().Interface().(int) <= 18 {
		return false
	}
	return true
}

// 创建中间件
func middel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first")
		c.Next()
		fmt.Println("next")
	}
}

func main() {

	// 数值转换
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)
	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	// 将常量保存为float32类型
	var c float64 = math.Pi
	fmt.Println(c)
	// f := 100.12345678901234567890123456789
	// 转换为int类型, 浮点发生精度丢失
	var d string = strconv.FormatFloat(c, 'f', -1, 64)
	fmt.Println(d)
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))

	// gin orm应用
	// dsn := "root:Aa@6447985@tcp(demo.gin-vue-admin.com:3306)/ginclass?chartset=utf8mb4&parseTime=True&loc=Local"
	// // db, err := gorm.Open("mysql", "root:Aa@6447985@/ginclass?chartset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// defer db.Close()
	// if err != nil {
	// 	panic(err)
	// }
	//Default返回一个默认的路由引擎
	r := gin.Default() // 携带基础中间件启动

	// 路由器分组
	v1 := r.Group("v1")
	// 分组增加中间件
	// v1 := r.Group("v1").Use(middel())
	v1.Use(middel())
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("我再分组方法内部")
		c.JSON(c.Writer.Status(), gin.H{
			"success": true,
		})
	})

	// 上传单\多文件
	// http://localhost:1010/testUpload
	// header Content-Type multipart/form-data
	// body form-data file file类型
	r.POST("/testUpload", func(c *gin.Context) {
		// name := c.PostForm("name")
		//  多文件
		form, _ := c.MultipartForm()
		file, _ := form.File["file"]
		fmt.Println(file)
		for _, f := range file {
			log.Println(f.Filename)
		}
		// 单文件
		// file, _ := c.FormFile("file")
		// fmt.Println(file)
		// gin 函数保存文件
		// c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		// c.SaveUploadedFile(file, "./"+file.Filename)
		// 手动保存文件
		// dst := "./uploads/" + file.Filename
		// src, _ := file.Open()

		// defer src.Close()

		// out, _ := os.Create(dst)

		// defer out.Close()

		// io.Copy(out, src)

		// // 文件回传给客户端
		// c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
		// c.File("/uploads/" + file.Filename)
		// c.JSON(c.Writer.Status(), gin.H{
		// 	"msg":  file,
		// 	"name": name,
		// 	"code": c.Writer.Status(),
		// })
	})

	// 使用结构体 bind绑定参数和参数验证
	r.POST("/testBindUri/", func(c *gin.Context) {
		// 验证规则
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}
		var p PostParams // 实例化结构体
		// http://localhost:1010/testBindUri?name=函数与&age=12233&sex=true
		err := c.ShouldBindQuery(&p)
		fmt.Println(err.Error())
		// /testBindUri/:name/:age/:sex
		// http://localhost:1010/testBindUri/函数与/1223/true
		// err := c.ShouldBindUri(&p)
		println("报错信息", gin.H{"data": 1})
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200, gin.H{
				"msg":  "报错",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功",
				"data": p,
			})
		}
	})
	// 使用结构体 bind绑定参数和参数验证
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams // 实例化结构体
		err := c.ShouldBindJSON(&p)
		println("报错信息", gin.H{"data": 1})
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "报错",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功",
				"data": p,
			})
		}
	})

	r.GET("/path/:id", func(c *gin.Context) {
		user := c.DefaultQuery("user", "wcy")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			user:      user,
			pwd:       pwd,
			"message": "哈哈哈",
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.PostForm("user")

		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user":    user,
			"pwd":     pwd,
			"message": "哈哈哈",
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "哈哈哈",
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.PostForm("user")

		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user":    user,
			"pwd":     pwd,
			"message": "哈哈哈",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "哈哈哈",
		})
	})
	r.Run(":1010") // listen and serve on 0.0.0.0:8080  Run(":1010")
	// r.Run() // listen and serve on 0.0.0.0:8080  Run(":1010")
	// get  参数挂载url中url  put form body  delete put
}
