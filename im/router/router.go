package router

import (
	"im/handler"
	"im/middleware"
	"im/ws"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(middleware.RateLimiter, middleware.Jwt)
	r.GET("/", handler.Index)
	ws.Router("/upgrade", r)
	user := r.Group("user")
	userRouter(user)
	friendRouter(r.Group("friend"))
	group := r.Group("group")
	groupRouter(group)
}
