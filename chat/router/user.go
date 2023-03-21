package router

import (
	"chat/handler"

	"github.com/gin-gonic/gin"
)

func user(r gin.IRouter) {
	r.POST("/login", handler.UserLogin)
	r.POST("/regist", handler.UserRegist)
	r.POST("/addfriend", handler.AddFriend)
}