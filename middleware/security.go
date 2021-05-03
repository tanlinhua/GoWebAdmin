package middleware

import (
	"github.com/gin-gonic/gin"
)

// xss/sql注入/csrf 等
func Safe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// trace.Debug("middleware.Safe.start")
	}
}

// func TestSafe() bool {
// 	test := html.EscapeString(`<script>alert("test")</script> select * from user`)
// 	fmt.Println(test)
// 	test2 := html.UnescapeString(test)
// 	fmt.Println(test2)

// 	if ok := FilteredSQLInject(test2); ok {
// 		fmt.Println("存在sql注入风险")
// 		return true
// 	}
// 	return false
// }

// // 正则过滤sql注入的方法，参数 : 要匹配的语句
// func FilteredSQLInject(to_match_str string) bool {
// 	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
// 	re, err := regexp.Compile(str)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return re.MatchString(to_match_str)
// }
