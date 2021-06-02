package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode   string
	APIPort   string
	AdminPort string

	AdminId   int
	AdminName string
	LoginAuth int

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisAddr string
	RedisPWD  string
	RedisDB   int

	JwtKey string

	SwaggerIsOpen int
)

//配置文件
func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Panic("配置文件[config.ini]读取失败，请检查!", err)
		return
	}
	loadServer(file)
	loadAdmin(file)
	loadDatabase(file)
	loadRedis(file)
	loadOther(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	APIPort = file.Section("server").Key("APIPort").MustString("3030")
	AdminPort = file.Section("server").Key("AdminPort").MustString("3031")
}

func loadAdmin(file *ini.File) {
	AdminId = file.Section("admin").Key("id").MustInt(1)
	AdminName = file.Section("admin").Key("name").MustString("admin")
	LoginAuth = file.Section("admin").Key("LoginAuth").MustInt(2)
}

func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func loadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("127.0.0.1:6379")
	RedisPWD = file.Section("redis").Key("RedisPWD").MustString("")
	RedisDB = file.Section("redis").Key("RedisDB").MustInt(0)
}

func loadOther(file *ini.File) {
	JwtKey = file.Section("jwt").Key("JwtKey").MustString("89js82js72")
	SwaggerIsOpen = file.Section("swagger").Key("IsOpen").MustInt(0)
}
