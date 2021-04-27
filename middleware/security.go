package middleware

import (
	"github.com/gin-gonic/gin"
)

// xss/sql注入/csrf 等
func Safe() gin.HandlerFunc {
	return func(c *gin.Context) {
		// trace.Debug("middleware.Safe.start")
		// tests/tests/html.go
	}
}
