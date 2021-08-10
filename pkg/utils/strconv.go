package utils

import (
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
