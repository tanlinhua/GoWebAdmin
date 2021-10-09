# go web项目通用模板

## 一些记录

### VSCode插件安装报错解决方案：
```shell
# 开启Go module，Go 1.13 以上默认启用，可跳过此步
go env -w GO111MODULE=on
# 设置代理
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
```shell
go get -u							#工程目录下执行,更新所有依赖包
go get -u github.com/gin-gonic/gin	#只更新这一个依赖包
go mod why -m all					#分析所有依赖的依赖链
go mod tidy							#整理依赖
```

### 部署
1. nohup
```shell
nohup ./main >> /www/wwwroot/log/main.nohup.log 2>&1 &
ps -ef|grep main
kill -9 pid
```
2. [Supervisor](docs/一些笔记/部署/Supervisor.md)

3. [nodejs pm2](https://cloud.tencent.com/developer/article/1677403)

### BT破解
```shell
sed -i "s|if (bind_user == 'True') {|if (bind_user == 'REMOVED') {|g" /www/server/panel/BTPanel/static/js/index.js
rm -rf /www/server/panel/data/bind.pl
```

### [WEB安全](docs/一些笔记/学习笔记/Web安全.md)

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

## Layui
```
Layui重要公告

所有对 layui 为之热爱、鞭策、奉献，和支持过的开发者：
请接受我用意念和字节传达的深深歉意。这是一个无力、无奈，甚至无助的决定：

layui 官网将于 2021年10月13日 进行下线。
届时，包括新版下载、文档和示例在内的所有框架日常维护工作，将全部迁移到 Github 和 Gitee。
此后，layui 仍会在代码托管平台所活跃，且 2.7 正式版也将在其间首发。而 layui 官网将不复存在。
这不是终结，只是重归到开源的纯粹中来。

再者，对于 layuiAdmin 和 layim 专区，将会迁移到新站进行保留，以便老用户还能下载使用，且此二者不再面向新用户。

过去五年，layui 有幸被应用在不计其数的 Web 平台，在前端工程化迅速席来的浪潮中，我们仍然感受到一丝来自于 jQuery 的余晖，这是一种带有热量的冰冷（反之亦可。
使命已达，便纵有万般遗憾，更与何人说？！

最后，请大家怀揣对 Web 前端技术的热忱，去拥抱 Vue.js、拥抱 Element UI、拥抱更好的新时代，
以及，所有那些值得去追求的美好事物。

—— 贤心
```
[github](https://github.com/sentsin/layui)
[gitee](https://gitee.com/sentsin/layui)

## Vue后台框架

> [GinVueAdmin.Web](https://github.com/flipped-aurora/gin-vue-admin)

> [naive-ui-admin](https://github.com/jekip/naive-ui-admin)

> [vue-element-admin](https://github.com/PanJiaChen)

> [vue-admin-beautiful-pro](https://github.com/chuzhixin/vue-admin-beautiful-pro)

> [2021，排名前 15 的 Vue 后台管理模板](https://mp.weixin.qq.com/s/4RVwmY8lOi4EmjR3iAW2nw)

> [Creative Tim](https://github.com/creativetimofficial)

## 一些库👇

> [合集1](https://learnku.com/articles/56078)
> [合集2](https://learnku.com/articles/41230)

> [Go 开发者路线图](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/zh-CN/ReadMe-zh-CN.md)

> [alipay](https://github.com/smartwalle/alipay)

> [wxpay](https://github.com/smartwalle/wxpay)

> [QrCode](https://github.com/skip2/go-qrcode)

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

> [getty网络框架](https://github.com/AlexStocks/getty-examples)

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

> [今日头条rpc框架 Kitex](https://www.cloudwego.io/zh/)

> [分布式事务管理器](https://github.com/yedf/dtm)

> [开源即时通讯框架](https://github.com/tinode/chat)

> [权限控制 casbin](https://github.com/casbin)

## 一些值得学习的项目

> 阅读Gin,Gorm,ants的源码

> [设计模式 Golang实现](https://github.com/senghoo/golang-design-pattern)

> [Go 语言实现的快速、稳定、内嵌的 k-v 数据库。](https://github.com/roseduan/rosedb)

> [NSQ](https://github.com/nsqio/nsq)

> [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)

> [Gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

> [数据结构与算法](https://github.com/TheAlgorithms)
