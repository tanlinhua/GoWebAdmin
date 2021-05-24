# go web项目通用模板

## 一些记录

### VSCode插件安装报错解决方案：
```
开启代理设置，Go 1.13 以上默认启用，可跳过此步
go env -w GO111MODULE=on
设置代理
go env -w GOPROXY=https://goproxy.io,direct
```

### go交叉编译
```
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

### 程序目录执行后台运行命令
```
nohup ./main >> /www/wwwroot/main.go.nohup.output.`date +%Y-%m-%d`.log 2>&1 &
ps -ef|grep main
->kill -9 pid
```

### centos 1W并发http.Get报错(socket: too many open files)
```
永久修改open files 方法
vim /etc/security/limits.conf  
在最后加入  
* soft nofile 65535
* hard nofile 65535
soft/hard前面的 * 表示所有用户
```
[或者通过此方案限制并发数](https://github.com/tanlinhua/GoTestDemo/blob/main/goroutine/semaphore/main.go)


### [redis常用命令](https://www.runoob.com/redis/redis-tutorial.html)
```
redis-cli 
auth "pwd"
ping
keys Task*
LLen Task_77
LRANGE Task_77 0 999
```

### jsdelivr+github cdn
```
https://cdn.jsdelivr.net/gh/用户名称/仓库名称@版本号/目录  
https://github.com/TurboWay/imgstore/blob/master/bigscreen/corp.jpg  
生成链接↓  
https://cdn.jsdelivr.net/gh/TurboWay/imgstore@master/bigscreen/corp.jpg 
```

### 一些库👇

> [Go 开发者路线图](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/zh-CN/ReadMe-zh-CN.md)

> https://github.com/smartwalle/alipay

> https://github.com/hashicorp/consul

> https://github.com/Shopify/sarama

> https://github.com/nicksnyder/go-i18n

> https://github.com/skip2/go-qrcode

> https://github.com/shirou/gopsutil

> https://github.com/tidwall/gjson

> https://github.com/maddevsio/fcm

---

## NOTES

### api模块增加(go gin swagger)
> https://blog.csdn.net/u013068184/article/details/106687646/

> https://github.com/tanlinhua/GoTestDemo/blob/main/swagger/main.go

---

## 初始化Vue管理后台HTTP服务
```
func InitVueAdminServer() {
	e := gin.New()
	e.Use(gin.Recovery())
    
	e.Static("js", "vue/admin/js")
	e.Static("css", "vue/admin/css")
	e.Static("fonts", "vue/admin/fonts")
	e.Static("img", "vue/admin/img")
	e.StaticFile("admin/favicon.ico", "vue/admin/favicon.ico")
	e.LoadHTMLGlob("vue/admin/index.html")
	e.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	e.Run(config.AdminPort)
}
```

## TODO

### xss/sql注入/csrf

- https://github.com/utrack/gin-csrf

- https://github.com/xinliangnote/PHP/blob/master/04-PHP%20WEB%20%E5%AE%89%E5%85%A8%E9%98%B2%E5%BE%A1.md

- https://www.iquanku.com/read/over-golang/04-Web%E7%BC%96%E7%A8%8B-10-Web%E5%AE%89%E5%85%A8.md

### go gin 项目 热更新/优雅的重启

- [Docker](https://blog.csdn.net/u010214802/article/details/90674343)

- https://github.com/gravityblast/fresh

- https://blog.csdn.net/qihoo_tech/article/details/104386331

### golang 消息队列

- [Kafka](kafka)
- [Nsq](https://blog.csdn.net/luolianxi/article/details/105279432)
- [Mqtt](mqtt)
- [RabbitMQ](rabbitmq)

### 浏览器关闭admin账号未退出
