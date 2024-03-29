# go web 项目通用模板

## TODO

## 一些记录

### VSCode 插件安装报错解决方案：

```shell
# 开启Go module，Go 1.13 以上默认启用，可跳过此步
go env -w GO111MODULE=on
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct
```

```
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
npm config set registry  https://registry.npmmirror.com
```

### 交叉编译

```shell
Windwos
SET CGO_ENABLED=0		#交叉编译不支持 CGO 所以要禁用它
SET GOOS=linux			#目标平台的操作系统 (darwin freebsd linux windows)
SET GOARCH=amd64		#目标平台的体系架构 (386 amd64 arm)
go build -o main main.go

Mac
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

### go mod

```shell
go get -u							#工程目录下执行,更新所有依赖包
go get -u github.com/gin-gonic/gin	#只更新这一个依赖包
go mod why -m all					#分析所有依赖的依赖链
go mod tidy							#整理依赖
```

### 漏洞检测

```
安装
$ go install golang.org/x/vuln/cmd/govulncheck@latest
使用
$ govulncheck ./...
官网
https://pkg.go.dev/vuln
https://go.dev/security/vuln/database
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

### Layui

[文档](https://layui.dev/docs/)

[github](https://github.com/sentsin/layui)

[gitee](https://gitee.com/sentsin/layui)

组件参考

[pearadmin](http://www.pearadmin.com/)
[LayuiMini](http://layuimini.99php.cn/)

快速构建表单

[Pear-Admin-Layui](https://gitee.com/pear-admin/Pear-Admin-Layui) 已 fork
-> VsCode Live Server -> 开发工具-表单构建

## [WEB 安全](docs/一些笔记/学习笔记/Web安全.md)

- [Web 安全学习笔记](https://github.com/LyleMi/Learn-Web-Hacking)
- [在线阅读](https://websec.readthedocs.io/zh/latest/)

## jsdelivr+github cdn

```
https://cdn.jsdelivr.net/gh/用户名称/仓库名称@版本号/目录
https://github.com/TurboWay/imgstore/blob/master/bigscreen/corp.jpg
生成链接↓
https://cdn.jsdelivr.net/gh/TurboWay/imgstore@master/bigscreen/corp.jpg
```

## Vue 后台框架

> [GinVueAdmin.Web](https://github.com/flipped-aurora/gin-vue-admin)

> [naive-ui-admin](https://github.com/jekip/naive-ui-admin)

> [vue-element-admin](https://github.com/PanJiaChen)

> [vue-admin-beautiful-pro](https://github.com/chuzhixin/vue-admin-beautiful-pro)

> [vue-pure-admin](https://www.bilibili.com/video/BV1534y1S7HV?p=1)

> [vue-admin-work/w](https://gitee.com/qingqingxuan/opend-vue-admin-work)

> [2021，排名前 15 的 Vue 后台管理模板](https://mp.weixin.qq.com/s/4RVwmY8lOi4EmjR3iAW2nw)

> [Creative Tim](https://github.com/creativetimofficial)

## 一些库 👇

> [合集 1](https://learnku.com/articles/56078) > [合集 2](https://learnku.com/articles/41230)

> [Go 开发者路线图](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/zh-CN/ReadMe-zh-CN.md)

> [alipay](https://github.com/smartwalle/alipay)

> [wxpay](https://github.com/smartwalle/wxpay)

> [QrCode](https://github.com/skip2/go-qrcode)

> [Json 解析.gjson](https://github.com/tidwall/gjson)

> [Json 解析.fastjson](https://github.com/valyala/fastjson)

> [FCM](https://github.com/maddevsio/fcm)

> [命令行.cobra](https://github.com/spf13/cobra)

> [命令行.urfave/cli](https://github.com/urfave/cli)

> [任务调度.Gron](https://github.com/roylee0704/gron)

> [任务调度.JobRunner](https://github.com/bamzi/jobrunner)

> [spf13/hugo](https://www.cnblogs.com/landv/p/11959097.html)

> [日志.zap](https://github.com/uber-go/zap)

> [日志.zerolog](https://github.com/rs/zerolog)

> [微服务.go-zero](https://github.com/tal-tech/go-zero)

> [微服务.rpcx](https://github.com/smallnest/rpcx)

> [getty 网络框架](https://github.com/AlexStocks/getty-examples)

> [gnet 网络框架](https://github.com/panjf2000/gnet)

> [协程池](https://github.com/panjf2000/ants)

> [websocket](github.com/gorilla/websocket)

> [telegram api](https://github.com/go-telegram-bot-api/telegram-bot-api)

> [key/value database](https://github.com/etcd-io/bbolt)

> [异步任务框架](https://github.com/RichardKnop/machinery)

> [(APNs)Apple Push Notification Service |](https://github.com/sideshow/apns2) > [| APNs Demo](https://github.com/Finb/bark-server/tree/master/apns)

> [简单易用的各种数据结构](https://github.com/emirpasic/gods)

> [爬虫框架](https://github.com/gocolly/colly)

> [chromedp](https://github.com/chromedp/chromedp)

> [今日头条 rpc 框架 Kitex](https://www.cloudwego.io/zh/)

> [分布式事务管理器](https://github.com/yedf/dtm)

> [开源即时通讯框架](https://github.com/tinode/chat)

> [权限控制 casbin](https://github.com/casbin)

> [hyperf/gotask](https://github.com/hyperf/gotask)

> [Python 资源大全中文版](https://blog.csdn.net/u013128262/article/details/79483998)

> [直播后台-GO](https://github.com/gwuhaolin/livego)

> [直播前端-JS](https://github.com/bilibili/flv.js)

> [Open-IM](https://www.rentsoft.cn/)

> [excel](https://github.com/qax-os/excelize)

> [carbon-时间处理库](https://github.com/golang-module/carbon)

> [dongle-编码解码&加密解密库](https://github.com/golang-module/dongle)

> [lancet-工具函数库](https://github.com/duke-git/lancet)

> [2D 游戏引擎](github.com/hajimehoshi/ebiten)

> [Rice-将静态文件打包到 Go 应用程序中](https://github.com/GeertJohan/go.rice)

## 一些值得学习的项目

> 阅读 Gin,Gorm,ants 的源码

> [设计模式 Golang 实现](https://github.com/senghoo/golang-design-pattern)

> [Go 语言实现的快速、稳定、内嵌的 k-v 数据库。](https://github.com/roseduan/rosedb)

> [NSQ](https://github.com/nsqio/nsq)

> [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)

> [Gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)

> [数据结构与算法](https://github.com/TheAlgorithms)

> [goploy](https://docs.goploy.icu/#/start/index)
