package utils

import "strings"

// 数组解析成字符串
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// 字符串解析成数组
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

// 检查数组[hystack]中是否存在某个值[needle]
func In_array(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// 比较数组，返回两个数组的差集
func ArrayDiff(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if !v {
			result = append(result, k)
		}
	}
	return result
}

// 比较数组，返回两个数组的交集
func ArrayIntersect(array1 []string, arrayOthers ...[]string) []string {
	c := make(map[string]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]string, 0)
	for k, v := range c {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// array_push — 将一个或多个单元压入数组的末尾（入栈）
func ArrayPush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}

// array_pop — 弹出数组最后一个单元（出栈）
func ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// array_merge — 合并一个或多个数组
func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}
