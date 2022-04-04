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
	db.AutoMigrate(&User{}, &Category{})
	// db.AutoMigrate(&User{}, &Category{}, &Article{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		println(err)
	}
	return db
}
func GetDB() *gorm.DB {
	return db
}
