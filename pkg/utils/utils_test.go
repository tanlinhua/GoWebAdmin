package utils

import (
	"testing"
)

func TestPassWord(t *testing.T) {
	password := "123456"
	hash, _ := PasswordHash(password)

	t.Log("密码:", password)
	t.Log("密码hash:", hash)

	match := PasswordVerify(password, hash)
	t.Log("校验结果:", match)
}

func TestShortUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(ShortUUID())
	}
}

//测试数组差集
func TestArrayDiff(t *testing.T) {
	a1 := []string{"test1", "test2", "test3"}
	a2 := []string{"test3", "test4"}
	r := ArrayDiff(a1, a2)
	t.Log(r)
}

//测试数组交集
func TestArrayIntersect(t *testing.T) {
	a1 := []string{"test1", "test2", "test3"}
	a2 := []string{"test3", "test4"}
	r := ArrayIntersect(a1, a2)
	t.Log(r)
}

//测试http get
func TestHttpGet(t *testing.T) {
	var params = make(map[string]string)
	params["a"] = "b"
	ok, resp := HttpGet("http://api.ipify.org", nil)
	t.Log(ok, resp)
}

//测试正则表达式
func TestRegxp(t *testing.T) {
	isMail := Is_Email("test@test.com")
	t.Log(isMail)
	isPhone := Is_Phone_China("13800138000")
	t.Log(isPhone)
}
