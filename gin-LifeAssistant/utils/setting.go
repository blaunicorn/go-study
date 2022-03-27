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
	charset    string
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
	Db = file.Section("server").Key("Db").MustString("mysql")
	DbHost = file.Section("server").Key("DbHost").MustString("localhost")
	DbPort = file.Section("server").Key("DbPort").MustString("3306")
	DbUser = file.Section("server").Key("DbUser").MustString("gin_essential")
	Dbpassword = file.Section("server").Key("Dbpassword").MustString("rr4bsaDkbeSwj2tb")
	DbName = file.Section("server").Key("DbName").MustString("gin_essential")
	charset = file.Section("server").Key("charset").MustString("utf8")
}
