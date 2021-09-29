package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(msg, url string, c *gin.Context) {
	c.HTML(http.StatusOK, "main/jump.html", gin.H{"code": 1, "wait": 3, "url": url, "msg": msg})
}

func Error(msg, url string, c *gin.Context) {
	c.HTML(http.StatusOK, "main/jump.html", gin.H{"code": 0, "wait": 3, "url": url, "msg": msg})
}
