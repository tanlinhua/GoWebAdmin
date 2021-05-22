package utils

import "testing"

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
