package middleware

import (
	ratelimiter "proxy/rateLimiter"
	"time"

	"github.com/gin-gonic/gin"
)


func RateLimiter() func(c *gin.Context){
	l := ratelimiter.NewLimiter(time.Second, 100)
	return func(c *gin.Context) {
		if !l.Allow() {
			c.Abort()
			return
		}
		c.Next()
	}
}