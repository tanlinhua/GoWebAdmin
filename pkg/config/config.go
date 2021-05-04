package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	APIPort    string
	AdminPort  string
	AdminId    int
	AdminName  string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

//配置文件
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Print("config/config.ini 配置文件读取失败，请检查", err)
	}
	loadServer(file)
	loadDatabase(file)
	loadAdmin(file)
}

func loadAdmin(file *ini.File) {
	AdminId = file.Section("admin").Key("id").MustInt(1)
	AdminName = file.Section("admin").Key("name").MustString("admin")
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	APIPort = file.Section("server").Key("APIPort").MustString("3030")
	AdminPort = file.Section("server").Key("AdminPort").MustString("3031")
	JwtKey = file.Section("jwt").Key("JwtKey").MustString("89js82js72")
}

func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}