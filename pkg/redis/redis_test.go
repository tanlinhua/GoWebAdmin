package redis

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {
	ret0 := Redis.SSet("test", "Testvalue", 30)
	t.Log(ret0)
	val1 := Redis.SGet("test")
	t.Log(val1)
	t.Log("==============")
	ret1 := Redis.ListAdd("Task", "test1")
	t.Log(ret1)
	ret2 := Redis.ListAdd("Task", "test2")
	t.Log(ret2)
	val2 := Redis.ListGet("Task")
	t.Log(val2)
	val3 := Redis.ListClear("Task")
	t.Log(val3)
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
