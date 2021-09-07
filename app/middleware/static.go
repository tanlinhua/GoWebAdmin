package middleware

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 浏览器缓存静态资源
func StaticFileHandler() gin.HandlerFunc {

	data := []byte(time.Now().String())
	maxAge := int((time.Hour * 24 * 30).Seconds())

	etag := fmt.Sprintf("W/%x", md5.Sum(data))       // ETag值
	ctl := fmt.Sprintf("public, max-age=%d", maxAge) // 浏览器缓存时间

	return func(c *gin.Context) {
		c.Header("Cache-Control", ctl)
		c.Header("ETag", etag)

		if match := c.GetHeader("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				c.Status(http.StatusNotModified)
				return
			}
		}
	}
}
