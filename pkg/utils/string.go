package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 变量替换 替换字符串content中所有的char字符
// content = utils.ReplaceContent(content, "#")
// content = utils.ReplaceContent(content, "@")
func ReplaceContent(content string, char string) string {
	count := strings.Count(content, char)
	if count <= 0 {
		return content
	}

	var chars string
	for i := 0; i < count; i++ {
		chars += char
	}
	rand := RandString(count)

	new := strings.ReplaceAll(content, chars, rand)
	return new
}

// 生成随机字符串
func RandString(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成随机数验证码
func GenValidateCode(len int) string {
	numbers := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < len; i++ {
		fmt.Fprintf(&sb, "%d", numbers[rand.Intn(10)])
	}
	return sb.String()
}

// 截取字符串
// haystack -> 源字符串
// needle -> 截取字符串
// before_needle -> true返回needle之前,false返回needle之后部分
func Strstr(haystack string, needle string, before_needle bool) string {
	idx := strings.Index(haystack, needle)
	if idx == -1 || needle == "" {
		return haystack
	}
	if before_needle {
		return haystack[0:idx]
	} else {
		return haystack[idx+len([]byte(needle))-1:]
	}
}

// 截取字符串
// source		-> 源字符串
// flagStart	-> 起始字符串
// flagEnd		-> 结束字符串
// return		-> flagStart到flagEnd的中间字符串
func SubStrByFlag(source, flagStart, flagEnd string) string {
	start := strings.Index(source, flagStart) + len(flagStart)
	end := strings.Index(source[start:], flagEnd) + start
	if end == start-1 {
		end = len(source)
	}
	return source[start:end]
}

// 获取字符串str中start到end中间的字符串
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		return ""
	}
	n += len(start)
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
