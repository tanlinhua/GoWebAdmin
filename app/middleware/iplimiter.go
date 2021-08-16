package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	rdb "github.com/tanlinhua/go-web-admin/pkg/redis"
	"github.com/tanlinhua/go-web-admin/pkg/response"
)

// IP指定时间内请求次数限制器,依赖redis存储
func IpLimiter() gin.HandlerFunc {

	slidingWindow := 5 * time.Second // var slidingWindow time.Duration
	limit := 100

	key := "ipLimiter"

	if rdb.Handler == nil {
		fmt.Println("IpLimiter: redis未就绪")
		return func(c *gin.Context) {}
	}

	return func(c *gin.Context) {
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key)
		now := time.Now().UnixNano()
		max := fmt.Sprint(now - (slidingWindow.Nanoseconds()))

		rdb.Handler.ZRemRangeByScore(userCntKey, "0", max).Result()

		reqs, _ := rdb.Handler.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			// c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"status": http.StatusTooManyRequests, "message": "too many request"})
			c.Abort()
			response.New(c).Error(http.StatusTooManyRequests, "too many request")
			return
		}

		c.Next()
		rdb.Handler.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		rdb.Handler.Expire(userCntKey, slidingWindow)
	}
}

// https://github.com/imtoori/gin-redis-ip-limiter
// https://www.runoob.com/redis/redis-sorted-sets.html

/*
func TestIpLimiter() {
	for i := 0; i < 110; i++ {
		c := &http.Client{}

		resp, e := c.Get("http://localhost:1991/admin/login")
		if e != nil {
			fmt.Println("请求错误:", e.Error())
			return
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		fmt.Println("响应状态码:", resp.StatusCode, "& body=", string(body), " & i=", i)
	}
}
*/
