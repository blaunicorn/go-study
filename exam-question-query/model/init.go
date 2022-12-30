package model

import (
	"exam-question-query/utils"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// var err error

func InitDB() {
	// var ormLogger logger.Interface

	// if gin.Mode() == "debug" {
	// 	ormLogger = logger.Default.LogMode(logger.Info)
	// } else {
	// 	ormLogger = logger.Default
	// }

	// args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
	// 	utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName, utils.DbCharset)
	// println(args)
	path := strings.Join([]string{utils.DbUser, ":", utils.DbPassWord, "@tcp(", utils.DbHost, ":", utils.DbPort, ")/", utils.DbName, "?charset=utf8&parseTime=true"}, "")
	println(path)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       path,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
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
	// db.AutoMigrate(&User{})
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
		println("err")
	}
	DB = db
	migration()
}

//执行数据迁移
func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{})
	if err != nil {
		return
	}
	//DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}

/*
gorm:v1
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)   //默认不加复数s
	db.DB().SetMaxIdleConns(20)  //设置连接池，空闲
	db.DB().SetMaxOpenConns(100) //打开
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
*/
