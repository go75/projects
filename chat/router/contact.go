package router

import (
	"chat/handler"

	"github.com/gin-gonic/gin"
)

func contact(r gin.IRouter) {
	r.POST("/friend", handler.LoadFriend)
	r.POST("/community", handler.LoadCommunity)
}