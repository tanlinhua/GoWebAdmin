package middleware

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultPrefix url prefix of pprof
	DefaultPrefix = "/debug/pprof"
)

func getPrefix(prefixOptions ...string) string {
	prefix := DefaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}
	return prefix
}

// Register the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func PprofRegister(r *gin.Engine, prefixOptions ...string) {
	RouteRegister(&(r.RouterGroup), prefixOptions...)
}

// RouteRegister the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.GrouterGroup. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func RouteRegister(rg *gin.RouterGroup, prefixOptions ...string) {
	prefix := getPrefix(prefixOptions...)

	prefixRouter := rg.Group(prefix)
	{
		prefixRouter.GET("/", gin.WrapF(pprof.Index))
		prefixRouter.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		prefixRouter.GET("/profile", gin.WrapF(pprof.Profile))
		prefixRouter.POST("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/symbol", gin.WrapF(pprof.Symbol))
		prefixRouter.GET("/trace", gin.WrapF(pprof.Trace))
		prefixRouter.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		prefixRouter.GET("/block", gin.WrapH(pprof.Handler("block")))
		prefixRouter.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		prefixRouter.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		prefixRouter.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		prefixRouter.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}

// https://github.com/gin-contrib/pprof

// 浏览器查看
// url = http://ip:prot/jason/pprof

// 交互式终端
// $ go tool pprof url/profile?seconds=60
// 可视化PDF: 首先进入交互式终端，然后输入pdf命令，其他文件格式执行pprof help查看命令说明。

// 可视化火焰图
// 启动可视化界面的命令格式：
// $ go tool pprof -http=":8081" [binary] [profile]
// binary：可执行程序的二进制文件，通过go build命名生成
// profile：protobuf格式的文件

// 协程分析
// 报告协程相关信息，可以用来查看有哪些协程正在运行、有多少协程在运行等。
// go tool pprof http://ip:port/debug/pprof/goroutine

// heap 分析
// heap：查看堆相关信息，包括一些GC的信息。
// go tool pprof http://ip:port/debug/pprof/heap

// block分析
// 报告协程阻塞的情况，可以用来分析和查找死锁等性能瓶颈，默认不开启， 需要调用开启。
// runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪
// go tool pprof url/pprof/block

// mutex 分析
// 查看互斥的争用情况，默认不开启， 需要调用需要在程序中调用
// runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
