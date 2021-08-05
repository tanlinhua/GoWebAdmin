package cache

import (
	"testing"
	"time"
)

type Permission struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	GroupName   string `json:"group_name"`
}

func TestCache(t *testing.T) {
	SetCache("k1", "v1", 5*time.Second)
	t.Log(GetCache("k1"))
	time.Sleep(6 * time.Second)
	t.Log(GetCache("k1"))

	var permissions = []Permission{
		{
			Name:        "A1",
			DisplayName: "查看员工列表",
			GroupName:   "A:员工管理",
		},
		{
			Name:        "A2",
			DisplayName: "添加员工",
			GroupName:   "A:员工管理",
		},
	}
	SetCache("go_cache_add_struct", permissions, 60*time.Minute)
	t.Log(GetCache("go_cache_add_struct"))
}
