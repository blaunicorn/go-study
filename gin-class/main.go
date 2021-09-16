package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	// 7. gorm使用
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// 9. jwt
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}

func (u User) TableName() string {
	return "qm_users"
	// if (u.Name == "aaa") {
	// 	return "admin_users"
	// } else {
	// 	return "users"
	// }
}

func NewUser(id int, username string, password string) *User {
	user := new(User)
	// user.id = id
	// user.username = username
	// user.password = password
	return user
}

type PostParams struct {
	Name string `json:"name" uri:"name" binding:"required"`
	Age  int    `json:"age" uri:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" uri:"sex" binding:"required"`
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
		c.Next() // 是否往下走
		fmt.Println("second")
	}
}

type HelloWorld struct {
	gorm.Model // 底层规范
	Name       string
	Sex        bool
	Age        int
}

// 学生 、老师 、班级 、身份证卡
type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
	Teachers  []Teacher
}
type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers;"` // 通过中间表关联
	// TeacherID uint
}
type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}
type Teacher struct {
	gorm.Model
	TeacherName string
	// StudentID uint
	Students []Student `gorm:"many2many:student_teachers;"`
}

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func main() {
	r := gin.Default() // 携带基础中间件启动

	// 创建记录日记的文件
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.POST("/login", func(c *gin.Context) {
		// 加密
		claims := MyClaims{
			UserName: "qimiao",
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix() - 60,
				ExpiresAt: time.Now().Unix() + 60*60*2,
				Issuer:    "hahah",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		fmt.Println(token)
		mySigningKey := []byte("woshiwcyhhahhah")
		s, e := token.SignedString(mySigningKey)
		if e != nil {
			fmt.Println(e)
			return
		}
		fmt.Println(s)

		// 解密
		untoken, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(untoken)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    "hello",
			"code":    c.Writer.Status(),
		})
	})

	r.POST("/student", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "wcy_school_app:dGZfEfFZA7zb5kw4@tcp(81.70.77.41:3306)/wcy_school_app?charset=utf8mb4&parseTime=True&loc=Local")
		db.AutoMigrate(&Teacher{}, &Student{}, &Class{}, &IDCard{}) // 实例的指针地址
		// db.AutoMigrate(&Teacher{}, &Student{}, &Class{}, &IDCard{}) // 实例的指针地址
		db.SingularTable(false) // 不修改表名
		defer db.Close()
		if err != nil {
			panic(err)
		}
		// 创建学生数组
		// i := IDCard{
		// 	Num: 123445,
		// }

		// t := Teacher{
		// 	TeacherName: "zhangsan",
		// }
		// s := Student{
		// 	StudentName: "qm",
		// 	IDCard:      i,
		// 	Teachers:    []Teacher{t},
		// }
		// class := Class{
		// 	ClassName: "一班",
		// 	Students:  []Student{s},
		// }

		// // create
		// _ = db.Create(&i).Error
		// _ = db.Create(&t).Error
		// _ = db.Create(&s).Error
		// _ = db.Create(&class).Error
		// 创建学生数组结束main

		// 用postman创建学生数据
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
		// db.Create(&HelloWorld{
		// 	Name: "chengdu",
		// 	Sex:  false,
		// 	Age:  20,
		// })

		// find
		// var hello HelloWorld
		// db.First(&hello, "name=?", "chengdu")
		// var hello []HelloWorld
		// db.Find(&hello) // 接收的是切片的地址
		// db.Find(&hello, "age<?", 19) // 也可以如下链式调用main
		// db.Where("age<?", 19).Find(&hello)

		// update
		// db.Where("id in (?)", []int{1, 2}).Find(&hello).Updates(map[string]interface{}{
		// 	"Name": "qimiaoshua",
		// 	"Age":  22,
		// 	"Sex":  false,
		// }) // 批量 &[]HelloWorld{} or &hello
		// db.Where("id = ?", 1).Find(&hello).Updates(map[string]interface{}{
		// 	"Name": "qimiaoshua",
		// 	"Age":  22,
		// 	"Sex":  false,
		// })
		// db.Where("age<?", 19).Find(&hello).Updates(HelloWorld{
		// 	Name:"qimiaoshua",
		// 	Age: 22,
		// })
		// db.Where("age<?", 19).Find(&hello).Update("name", "qimiao")

		// delete
		// db.Delete(&hello, "id=?", 1) // Unscoped() 物理删除
		// fmt.Println(hello)
		// fmt.Println(hello)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    "hello",
			"code":    c.Writer.Status(),
		})
	})
	r.GET("/student/:ID", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "wcy_school_app:dGZfEfFZA7zb5kw4@tcp(81.70.77.41:3306)/wcy_school_app?charset=utf8mb4&parseTime=True&loc=Local")
		db.AutoMigrate(&Student{}) // 实例的指针地址
		// db.AutoMigrate(&Teacher{}, &Student{}, &Class{}, &IDCard{}) // 实例的指针地址
		db.SingularTable(false) // 不修改表名
		defer db.Close()
		if err != nil {
			panic(err)
		}

		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").First(&student, "id=?", id)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    student,
			"code":    c.Writer.Status(),
		})
	})
	r.GET("/class/:ID", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "wcy_school_app:dGZfEfFZA7zb5kw4@tcp(81.70.77.41:3306)/wcy_school_app?charset=utf8mb4&parseTime=True&loc=Local")
		// db.AutoMigrate(&Class{}) // 实例的指针地址
		defer db.Close()
		if err != nil {
			panic(err)
		}

		id := c.Param("ID")
		var class Class
		_ = c.BindJSON(&class)
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").Find(&class, "id=?", id)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    class,
			"code":    c.Writer.Status(),
		})
	})

	r.POST("/postDb", func(c *gin.Context) {
		// dsn := "root:Aa@6447985@tcp(demo.gin-vue-admin.com:3306)/ginclass?chartset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open("mysql", "wcy_school_app:dGZfEfFZA7zb5kw4@tcp(81.70.77.41:3306)/wcy_school_app?charset=utf8mb4&parseTime=True&loc=Local")
		// dsn := "root:Aa@6447985@tcp(demo.gin-vue-admin.com:3306)/ginclass?chartset=utf8mb4&parseTime=True&loc=Local"
		// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&HelloWorld{}) // 实例的指针地址
		db.SingularTable(false)       // 不修改表名

		// create
		// db.Create(&HelloWorld{
		// 	Name: "chengdu",
		// 	Sex:  false,
		// 	Age:  20,
		// })

		// find
		// var hello HelloWorld
		// db.First(&hello, "name=?", "chengdu")
		var hello []HelloWorld
		// db.Find(&hello) // 接收的是切片的地址
		// db.Find(&hello, "age<?", 19) // 也可以如下链式调用main
		// db.Where("age<?", 19).Find(&hello)

		// update
		// db.Where("id in (?)", []int{1, 2}).Find(&hello).Updates(map[string]interface{}{
		// 	"Name": "qimiaoshua",
		// 	"Age":  22,
		// 	"Sex":  false,
		// }) // 批量 &[]HelloWorld{} or &hello
		// db.Where("id = ?", 1).Find(&hello).Updates(map[string]interface{}{
		// 	"Name": "qimiaoshua",
		// 	"Age":  22,
		// 	"Sex":  false,
		// })
		// db.Where("age<?", 19).Find(&hello).Updates(HelloWorld{
		// 	Name:"qimiaoshua",
		// 	Age: 22,
		// })
		// db.Where("age<?", 19).Find(&hello).Update("name", "qimiao")

		// delete
		db.Delete(&hello, "id=?", 1) // Unscoped() 物理删除
		fmt.Println(hello)
		// fmt.Println(hello)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    hello,
			"code":    c.Writer.Status(),
		})
		defer db.Close()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello pong",
		})
	})
	r.GET("/path/:id", func(c *gin.Context) {

		// id := c.Param("id")
		id, _ := strconv.Atoi(c.Param("id"))
		username := c.DefaultQuery("username", "ooo")
		// username := c.Query("username")
		password := c.Query("password")
		data := NewUser(id, username, password)
		c.JSON(200, gin.H{
			"message":  "hello pong1",
			"data":     data,
			"id":       id,
			"username": username,
			"success":  true,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123")
		c.JSON(200, gin.H{
			"message":  "hello pong1",
			"username": username,
			"password": password,
			"success":  true,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		c.JSON(200, gin.H{
			"message": "hello pong",
			"success": true,
			"id":      id,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "123")
		c.JSON(200, gin.H{
			"message":  "hello pong put",
			"username": username,
			"password": password,
			"success":  true,
		})
	})
	r.POST("/testBind", func(c *gin.Context) {
		// 验证规则
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}
		var p PostParams
		err := c.ShouldBindJSON(&p)
		// err := c.ShouldBind(&p)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200, gin.H{
				"message": "参数错误。" + err.Error(),
				"success": false,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "hello testBind",
			"data":    p,
			"success": true,
		})

	})

	// http://localhost:1010/testUpload
	// header Content-Type multipart/form-data
	// body form-data file file类型
	r.POST("/testUpload", func(c *gin.Context) {
		name := c.PostForm("name")
		//多文件上传
		form, _ := c.MultipartForm()
		file, _ := form.File["file"]
		for _, f := range file {
			fmt.Println(f.Filename)
			c.SaveUploadedFile(f, "./uploads/"+f.Filename)
		}
		// 单文件上传
		// file, _ := c.FormFile("file")
		// fmt.Println((file.Filename))
		// // gin 函数上传
		// // c.SaveUploadedFile(file, "././uploads/"+file.Filename)
		// // 自定义实现
		// dst := "./uploads/" + file.Filename
		// src, _ := file.Open()
		// defer src.Close()
		// out, _ := os.Create(dst)
		// defer out.Close()
		// io.Copy(out, src)
		// // 文件回传给客户端
		// c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
		// c.File("./uploads/" + file.Filename)
		c.JSON(c.Writer.Status(), gin.H{
			"message": "hello testBind",
			"data":    file,
			"name":    name,
			"success": true,
			"code":    c.Writer.Status(),
		})
	})

	// 用户分组
	v1 := r.Group("v1").Use(middel())
	// v1 := r.Group("v1")
	v1.GET("/test", func(c *gin.Context) {
		fmt.Println("我再分组方法内部")
		c.JSON(c.Writer.Status(), gin.H{
			"success": true,
			"code":    c.Writer.Status(),
		})
	})
	// r.Run() // listen and server on 8080
	r.Run(":1010")
}
