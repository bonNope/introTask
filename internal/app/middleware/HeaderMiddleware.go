package middleware

import "github.com/gin-gonic/gin"

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("X-PING")
		if header == "ping" {
			c.Writer.Header().Set("X-PONG", "pong")
		}
		c.Next()
	}
}
