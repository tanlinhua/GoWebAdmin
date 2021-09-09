package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode       string
	APIPort       string
	AdminPort     string
	ExtIP         string
	AdminId       int
	AdminName     string
	GoogleAuth    int
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string
	RedisAddr     string
	RedisPWD      string
	RedisDB       int
	JwtKey        string
	LangSwitch    bool
	LangDefault   string
	SwaggerIsOpen int
	TgHost        string
	TgToken       string
	TgChatId      string
)

//配置文件
func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Panic("配置文件[config.ini]读取失败，请检查!", err)
		return
	}

	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	APIPort = file.Section("server").Key("APIPort").MustString("3030")
	AdminPort = file.Section("server").Key("AdminPort").MustString("3031")
	ExtIP = file.Section("server").Key("ExtIP").MustString("")

	AdminId = file.Section("admin").Key("id").MustInt(1)
	AdminName = file.Section("admin").Key("name").MustString("admin")
	GoogleAuth = file.Section("admin").Key("GoogleAuth").MustInt(2)

	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("127.0.0.1:6379")
	RedisPWD = file.Section("redis").Key("RedisPWD").MustString("")
	RedisDB = file.Section("redis").Key("RedisDB").MustInt(0)

	JwtKey = file.Section("jwt").Key("key").MustString("89js82js72")

	LangSwitch = file.Section("lang").Key("switch").MustBool(false)
	LangDefault = file.Section("lang").Key("default").MustString("zh-cn")

	SwaggerIsOpen = file.Section("swagger").Key("IsOpen").MustInt(0)

	TgHost = file.Section("notify").Key("TgHost").MustString("https://api.telegram.org/")
	TgToken = file.Section("notify").Key("TgToken").MustString("")
	TgChatId = file.Section("notify").Key("TgChatId").MustString("")
}
