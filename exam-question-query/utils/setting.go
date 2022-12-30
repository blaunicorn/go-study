package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	DbCharset  string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请坚持文件路径：", err)
	}
	LoadServer(file)
	LoadDataBase(file)
	LoadBucket(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8081")
	JwtKey = file.Section("server").Key("JwtKey").MustString("a_secret_crect")
}

func LoadDataBase(file *ini.File) {

	Db = file.Section("mysql").Key("Db").MustString("mysql")
	DbHost = file.Section("mysql").Key("DbHost").MustString("localhost")
	DbPort = file.Section("mysql").Key("DbPort").MustString("3306")
	DbUser = file.Section("mysql").Key("DbUser").MustString("gin_essential")
	DbPassWord = file.Section("mysql").Key("DbPassWord").MustString("rr4bsaDkbeSwj2tb")
	DbName = file.Section("mysql").Key("DbName").MustString("gin_essential")
	DbCharset = file.Section("mysql").Key("charset").MustString("utf8")
}

//图床
func LoadBucket(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("")
	QiniuServer = file.Section("qiniu").Key("QiniuServer").MustString("")

}

