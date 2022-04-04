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
