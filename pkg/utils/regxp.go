package utils

import (
	"regexp"
)

// str中是否为纯数字
func Is_Number(str string) bool {
	return regexp.MustCompile("^[0-9]+$").MatchString(str)
}

// 帐号校验,字母开头,允许5-16字节,允许字母数字下划线
func Is_UserName(str string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{4,15}$`)
	return reg.MatchString(str)
}

// 邮箱
func Is_Email(email string) bool {
	pattern := `\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 身份证
func Is_IDCard(email string) bool {
	return regexp.MustCompile(`\d{17}[\d|x]|\d{15}`).MatchString(email)
}

// IP地址
func Is_IP(email string) bool {
	p := `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`
	return regexp.MustCompile(p).MatchString(email)
}

// 网址
func Is_Url(v string) (r bool) {
	reg := regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`)
	return reg.MatchString(v)
}

// 手机号-中国
func Is_Phone_China(mobileNum string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mobileNum)
}

// sql注入风险字符检查
func SQLInjectCheck(to_match_str string) bool {
	p := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	reg := regexp.MustCompile(p)
	return reg.MatchString(to_match_str)
}
