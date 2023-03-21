package router

import (
	"im/handler"

	"github.com/gin-gonic/gin"
)

func friendRouter(r gin.IRouter) {
	// 好友列表
	r.GET("/list", handler.GetUserFriends)
	// 添加好友
	r.POST("/add", handler.AddFriend)
	// 删除好友
	r.DELETE("/delete", handler.DeleteFriend)
	// 获取好友会话
	r.GET("/session", handler.Session)
	// 新好友消息
	r.GET("/new", handler.NewFriendsInfo)
}