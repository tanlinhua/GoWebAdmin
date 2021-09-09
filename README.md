# go web项目通用模板

## 一些记录

### VSCode插件安装报错解决方案：
```
开启代理设置，Go 1.13 以上默认启用，可跳过此步
go env -w GO111MODULE=on
设置代理
go env -w GOPROXY=https://goproxy.io,direct
```

### 交叉编译
```shell
SET CGO_ENABLED=0		#交叉编译不支持 CGO 所以要禁用它
SET GOOS=linux			#目标平台的操作系统 (darwin freebsd linux windows)
SET GOARCH=amd64		#目标平台的体系架构 (386 amd64 arm)
go build -o main main.go
```

### go mod
```
go get -u								工程目录下执行,更新所有依赖包,
go get -u github.com/gin-gonic/gin		只更新这一个依赖包
go mod why -m all						分析所有依赖的依赖链
go mod tidy								整理依赖
```

### 部署
1. nohup
```
nohup ./main >> /www/wwwroot/nohup.output.log 2>&1 &
ps -ef|grep main
kill -9 pid
```
2. [Supervisor](docs/一些笔记/Supervisor.md)

3. [nodejs pm2](https://cloud.tencent.com/developer/article/1677403)

### [WEB安全](docs/一些笔记/Web安全.md)

- [Web安全学习笔记](https://github.com/LyleMi/Learn-Web-Hacking)
- [在线阅读](https://websec.readthedocs.io/zh/latest/)

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

> [wxpay](https://github.com/smartwalle/wxpay)

> [QrCode](https://github.com/skip2/go-qrcode)

> [psutil for golang](https://github.com/shirou/gopsutil)

> [Json解析.gjson](https://github.com/tidwall/gjson)

> [Json解析.fastjson](https://github.com/valyala/fastjson)

> [FCM](https://github.com/maddevsio/fcm)

> [命令行.cobra](https://github.com/spf13/cobra)

> [命令行.urfave/cli](https://github.com/urfave/cli)

> [任务调度.Gron](https://github.com/roylee0704/gron)

> [任务调度.JobRunner](https://github.com/bamzi/jobrunner)

> [spf13/hugo](https://www.cnblogs.com/landv/p/11959097.html)

> [日志.zap](https://github.com/uber-go/zap)

> [微服务.go-zero](https://github.com/tal-tech/go-zero)

> [微服务.rpcx](https://github.com/smallnest/rpcx)

> [gnet网络框架](https://github.com/panjf2000/gnet)

> [协程池](https://github.com/panjf2000/ants)

> [websocket](github.com/gorilla/websocket)

> [telegram api](https://github.com/go-telegram-bot-api/telegram-bot-api)

> [key/value database](https://github.com/etcd-io/bbolt)

> [异步任务框架](https://github.com/RichardKnop/machinery)

> [(APNs)Apple Push Notification Service |](https://github.com/sideshow/apns2)
> [| APNs Demo](https://github.com/Finb/bark-server/tree/master/apns)

> [简单易用的各种数据结构](https://github.com/emirpasic/gods)

> [爬虫框架](https://github.com/gocolly/colly)

## 一些值得学习的项目

> 阅读Gin,Gorm,ants的源码

> [设计模式 Golang实现](https://github.com/senghoo/golang-design-pattern)

> [Go 语言实现的快速、稳定、内嵌的 k-v 数据库。](https://github.com/roseduan/rosedb)

> [NSQ](https://github.com/nsqio/nsq)

> [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)

> [Gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

## TODO

### casbin
- 百度&b站: golang casbin
- 参考gin-vue-admin
- https://www.bilibili.com/video/BV1qz4y167XP
- https://www.bilibili.com/video/BV13r4y1M7AC
- go get github.com/casbin/casbin/v2
- go get github.com/casbin/gorm-adapter/v3

### i18n
- https://blog.csdn.net/qq_42041981/article/details/105233384 (layui国际化)
- https://www.cnblogs.com/smallbo/p/15042203.html
- https://xuanwo.io/2019/12/11/golang-i18n/
- https://github.com/nicksnyder/go-i18n
- https://goframe.org/pages/viewpage.action?pageId=7301652
- https://github.com/suisrc/gin-i18n

### [go gin zap](https://www.liwenzhou.com/posts/Go/use_zap_in_gin/)