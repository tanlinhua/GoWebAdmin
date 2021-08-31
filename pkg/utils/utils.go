package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

// md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// password_hash
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// password_verify
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 转json
func Json_encode(data interface{}) (string, error) {
	// jsons, err := json.Marshal(data)
	// return string(jsons), err

	// 转义符 \u0026 , 设置json序列化不转义
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(data)
	return buffer.String(), err
}

// 解json
func Json_decode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

// 判断变量是否为空
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}
