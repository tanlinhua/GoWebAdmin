package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode          string
	AppHost          string
	AdminPort        string
	AdminId          int
	AdminGoogleAuth  int
	ApiPort          string
	DbType           string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPwd            string
	DbName           string
	RedisHost        string
	RedisPWD         string
	RedisDB          int
	JwtKey           string
	LangSwitch       bool
	LangDefault      string
	SwaggerOpen      int
	TelegramHost     string
	TelegramBotToken string
	TelegramChatId   string
)

//配置文件
func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Panic("配置文件[config.ini]读取失败，请检查!", err)
		return
	}

	AppMode = file.Section("App").Key("Mode").MustString("release")
	AppHost = file.Section("App").Key("Host").MustString("localhost")

	AdminPort = file.Section("Admin").Key("Port").MustString("1991")
	AdminId = file.Section("Admin").Key("Id").MustInt(1)
	AdminGoogleAuth = file.Section("Admin").Key("GoogleAuth").MustInt(0)

	ApiPort = file.Section("Api").Key("Port").MustString("2014")

	DbType = file.Section("Db").Key("Type").MustString("")
	DbHost = file.Section("Db").Key("Host").MustString("")
	DbName = file.Section("Db").Key("Name").MustString("")
	DbPort = file.Section("Db").Key("Port").MustString("")
	DbUser = file.Section("Db").Key("User").MustString("")
	DbPwd = file.Section("Db").Key("Pwd").MustString("")

	RedisHost = file.Section("Redis").Key("Host").MustString("127.0.0.1:6379")
	RedisPWD = file.Section("Redis").Key("Pwd").MustString("")
	RedisDB = file.Section("Redis").Key("Db").MustInt(0)

	JwtKey = file.Section("Jwt").Key("Key").MustString("")

	LangSwitch = file.Section("Lang").Key("Switch").MustBool(false)
	LangDefault = file.Section("Lang").Key("Default").MustString("zh-cn")

	SwaggerOpen = file.Section("Swagger").Key("Open").MustInt(0)

	TelegramHost = file.Section("Telegram").Key("Host").MustString("https://api.telegram.org/")
	TelegramBotToken = file.Section("Telegram").Key("BotToken").MustString("")
	TelegramChatId = file.Section("Telegram").Key("ChatId").MustString("")
}
