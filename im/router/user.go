package router

import (
	"im/handler"

	"github.com/gin-gonic/gin"
)

func userRouter(r gin.IRouter) {
	// 用户登录界面
	r.GET("/login", handler.UserLoginPage)
	// 用户注册界面
	r.GET("/regist", handler.UserRegistPage)
	// 用户头像
	r.GET("/head", handler.UserHead)
	// 用户登录
	r.POST("/login", handler.UserLogin)
	// 注册用户
	r.POST("/regist", handler.UserRegist)
	// 删除用户
	r.DELETE("/delete", handler.DeleteUser)
	// 修改用户
	r.POST("/update", handler.UpdateUser)
	// 查询用户
	r.GET("/query", handler.QueryUserByName)
}