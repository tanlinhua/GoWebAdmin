# go web项目通用模板

## 目录拆分
```
└── GoWebAdmin
	├── pkg			(功能性扩展包)
	└── service		(业务层服务包)
```

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

### socket: too many open files
```
vim /etc/security/limits.conf
在最后加入
* soft nofile 65535
* hard nofile 65535

* soft nproc 65535
* hard nproc 65535

tips↓
* 表示所有用户
soft/hard 软硬限制
nproc 最大线程数 / nofile 最大文件数
```
```
临时修改
cat /proc/39977/limits
prlimit --nofile=65536:65536 --pid 39977

查看当前系统打开的文件数量,代码如下:
lsof | wc -l
watch "lsof | wc -l"

查看某一进程的打开文件数量,代码如下:
lsof -p 3490 | wc -l
```
```
supervisor 控制的程序
CentOS 上使用系统自带的 supervisor，使用 systemd 启动 supervisord 的服务。被 supervisor 管理的程序，
继承的是 systemd 对应的限制，如果需要修改的话，就需要在启动.service 文件里面修改对应的限制

方法1:
vi /usr/lib/systemd/system/supervisord.service

[Service]
Type=forking
LimitNOFILE=102400
LimitNPROC=102400
方法2:
vi /etc/supervisord.conf
minfds=102400
```
[或者通过此方案限制并发数](pkg/gpool/docs/demo.md)


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

## 初始化Vue管理后台HTTP服务
```go
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

## 一些库👇

> [合集1](https://learnku.com/articles/56078)
> [合集2](https://learnku.com/articles/41230)

> [Go 开发者路线图](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/zh-CN/ReadMe-zh-CN.md)

> [alipay](https://github.com/smartwalle/alipay)

> [QrCode](https://github.com/skip2/go-qrcode)

> [psutil for golang](https://github.com/shirou/gopsutil)

> [Json解析.gjson](https://github.com/tidwall/gjson)

> [Json解析.fastjson](https://github.com/valyala/fastjson)

> [FCM](https://github.com/maddevsio/fcm)

> [命令行.cobra](https://github.com/spf13/cobra)

> [命令行.urfave/cli](https://github.com/urfave/cli)

> [任务调度.Gron](https://github.com/roylee0704/gron)

> [任务调度.JobRunner](https://github.com/bamzi/jobrunner)

> [github.com/spf13/hugo](https://www.cnblogs.com/landv/p/11959097.html)

> [日志.zap](https://github.com/uber-go/zap)

> [微服务.go-zero](https://github.com/tal-tech/go-zero)

> [微服务.rpcx](https://github.com/smallnest/rpcx)

> [gnet网络框架](https://github.com/panjf2000/gnet)

> [协程池](https://github.com/panjf2000/ants)

> [设计模式](https://github.com/tanlinhua/golang-design-pattern)

> [pprof](https://github.com/gin-contrib/pprof)

> [websocket](github.com/gorilla/websocket)

> [telegram api](https://github.com/go-telegram-bot-api/telegram-bot-api)

> [key/value database](https://github.com/etcd-io/bbolt)

> [csrf](https://github.com/gorilla/csrf)


## 一些值得学习的项目

> [设计模式 Golang实现](https://github.com/senghoo/golang-design-pattern)

> [Go 语言实现的快速、稳定、内嵌的 k-v 数据库。](https://github.com/roseduan/rosedb)

> [NSQ](https://github.com/nsqio/nsq)

> [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)

> [Gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

## TODO

### i18n
- https://xuanwo.io/2019/12/11/golang-i18n/
- https://github.com/nicksnyder/go-i18n

### 升级Grom到2.0

### [go gin zap](https://www.liwenzhou.com/posts/Go/use_zap_in_gin/)