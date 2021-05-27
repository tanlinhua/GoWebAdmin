package middleware

import (
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/pkg/response"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

// xss/sql注入/csrf 等
func Safe() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			trace.Error("c.GetRawData.error:" + err.Error())
		}

		xData := FilterXSSInject(string(data)) // 问题:string(data),会把<转成ascii码%3C

		if ok := FilterSQLInject(xData); ok {
			response.New(c).Error(404, "fail")
			c.Abort()
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(xData))) // 很关键,把读过的字节流重新放到body
	}
}

// 过滤xss风险字符
func FilterXSSInject(buf string) string {
	fmt.Println("buf", buf)
	res := html.EscapeString(buf)
	fmt.Println("res", res)
	return res
}

// 过滤sql注入风险字符
func FilterSQLInject(to_match_str string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		trace.Error("FilterSQLInject.err=" + err.Error())
	}
	return re.MatchString(to_match_str)
}
