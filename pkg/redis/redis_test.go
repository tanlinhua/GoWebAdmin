package redis

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestRedisKV(t *testing.T) {
	Redis.SSet("test", "pibigstar", 30)
	t.Log(Redis.SGet("test"))
}

func TestRedisListPush(t *testing.T) {
	Redis.ListAdd("Task", "test1")
	Redis.ListAdd("Task", "test2")
	ret := Redis.ListAdd("Task", "test3")
	t.Log(ret)
}

func TestRedisListPop(t *testing.T) {
	ret := Redis.ListGet("Task")
	t.Log(ret)
}

func TestReidsListClean(t *testing.T) {
	ret := Redis.ListClear("Task")
	t.Log(ret)
}

var lua = `
local key1 = tostring(KEYS[1])
local key2 = tostring(KEYS[2])
local args1 = tonumber(ARGV[1])
local args2 = tonumber(ARGV[2])

if key1 == "user"
then
	redis.call('SET',key1,args1)
	return 1
else
	redis.call('SET',key2,args2)
	return 2
end
return 0
`

func TestLua(t *testing.T) {
	client, err := GetRedisClient()
	if err != nil {
		t.Error(err)
	}
	script := redis.NewScript(lua)

	cmd := script.Run(client, []string{"user", "test"}, 1, 2)
	t.Log(cmd.Result())
}
