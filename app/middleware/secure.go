package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func Secure() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		BrowserXssFilter: true,
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}
	}
}

// 同类型中间件: https://github.com/gin-contrib/secure

/*
Golang/Gin框架添加对HTTPS的支持
HTTPS配置步骤:

1.首先在阿里云搞定ICP域名备案
2.添加一个子域名
3.给子域名申请免费 SSL 证书, 然后下载证书对应的 pem 和 key 文件.
用 GIN 框架添加一个 github.com/unrolled/secure 中间件就可以了.

下面是一个简单的示例代码:
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/unrolled/secure"
)

func main() {
    router := gin.Default()
    router.Use(TlsHandler())

    router.RunTLS(":8080", "ssl.pem", "ssl.key")
}

func TlsHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        secureMiddleware := secure.New(secure.Options{
            SSLRedirect: true,
            SSLHost:     "localhost:8080",
        })
        err := secureMiddleware.Process(c.Writer, c.Request)

        // If there was an error, do not continue.
        if err != nil {
            return
        }

        c.Next()
    }
}
上面代码直接在子域名前添加 HTTPS 就可以安全通讯了.
*/
