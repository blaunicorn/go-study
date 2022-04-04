# gin+vue+uniapp 全栈博客+微信小程序-生活助手

## 初始化项目
```
go mod init gin-vue-lifeassistant 
go mod tidy     
```
## 运行项目
```
go run main.go
```
### ps:如果vscode 没有自动更新补全go语言,运行以下命令
```
// ctrl + shift + p 
go:install/update tools
```

## 建立目录结构
```
|-- api 接口
    |-- user.go
    |-- category.go
    └─ article.go
|-- config
    |-- config.ini  配置文件
|-- middleware  中间件：jwt等
|-- model 数据模型,映射数据库表结构
    |-- Article.go
    |-- Category.go
    |-- user.go
    |-- db.go  连接数据库
|-- routes 路由
|-- upload 上传文件目录
|-- utils 工具
    |-- setting.go 读取ini配置文件，生成参数
|-- main.go 入口
```

## 1、读取参数，启动路由
```
// utils\setting.go
package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	Dbpassword string
	DbName     string
	DbCharset  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请坚持文件路径：", err)
	}
	LoadServer(file)
	LoadDataBase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8081")
}

func LoadDataBase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("gin_essential")
	Dbpassword = file.Section("database").Key("Dbpassword").MustString("rr4bsaDkbeSwj2tb")
	DbName = file.Section("database").Key("DbName").MustString("gin_essential")
	DbCharset = file.Section("database").Key("charset").MustString("utf8")
}

```
```
// routes\routes.go
package routes

import (
	"gin-vue-lifeassistant/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)
}

```

```
main.go 入口引入配置文件和路由
package main

import (
	"gin-vue-lifeassistant/model"
	"gin-vue-lifeassistant/routes"
)

func main() {
	model.InitDB()
	routes.InitRouter()
}

```

### 2.连接数据库，建立模型
```
// model\db.go  并在入口main.go文件中引入 	model.InitDB()
package model

import (
	"fmt"
	"gin-vue-lifeassistant/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

var err error

func InitDB() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		utils.DbUser, utils.Dbpassword, utils.DbHost, utils.DbPort, utils.DbName, utils.DbCharset)
	println(args)
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			// NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	// migration schema
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
func GetDB() *gorm.DB {
	return db
}

```
```
// model\User.go 用户模型
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" json:"username"`
	Telephone string `gorm:"varchar(11);not null;unique" json:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Role      int    `gorm:"type:int" json:"role"`
}

```

## 3.建立api逻辑错误处理模块和api接口