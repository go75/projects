package router

import (
	"chat/middleware"

	"github.com/gin-gonic/gin"
)

func page(r *gin.Engine) {
	r.LoadHTMLFiles("./view/chat/index.html", "./view/user/login.html", "./view/user/regist.html")
	r.Use(middleware.RateLimiter)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.GET("/regist", func(c *gin.Context) {
		c.HTML(200, "regist.html", nil)
	})
}