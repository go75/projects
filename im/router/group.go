package router

import (
	"im/handler"

	"github.com/gin-gonic/gin"
)

func groupRouter(r gin.IRouter) {
	// 获取群聊列表
	r.GET("/list", handler.GroupList)
	// 创建群聊
	r.POST("/regist", handler.GroupRegist)
	// 获取群聊会话
	r.GET("/session", handler.GroupSession)
	// 新群友
	r.GET("/new", handler.NewGroupFriendsInfo)

}