package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

func ToString(i interface{}) (string, error) {
	return cast.ToStringE(i)
}

func ToInt(i interface{}) (int, error) {
	return cast.ToIntE(i)
}

func ToInt64(i interface{}) (int64, error) {
	return cast.ToInt64E(i)
}

func ToTime(i interface{}) (time.Time, error) {
	return cast.ToTimeE(i)
}

//任意类型转换int64
func ToInt64s(str interface{}) (r int64) {
	item := fmt.Sprintf("%v", str)
	if item == "" {
		return
	}
	r, err := strconv.ParseInt(item, 10, 64)
	if err != nil {
		return
	}
	return
}

//任意类型转换boll
func ToBool(str interface{}) (r bool) {
	item := fmt.Sprintf("%v", str)
	if item == "" {
		return
	}
	b, err := strconv.ParseBool(item)
	if err != nil {
		return
	}
	return b
}

//任意类型转换string
func ToString2(str interface{}) (r string) {
	if str == nil {
		return
	}
	item := fmt.Sprintf("%v", str)
	if item == "" {
		return
	}
	r = item
	return
}
