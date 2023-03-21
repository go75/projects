package handler

import (
	"im/common/res"
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"

	"github.com/gin-gonic/gin"
)

// GetUserFriends
// @Tags 好友模块
// @Summary 获取所有好友
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/list [get]
func GetUserFriends(c *gin.Context) {
	//id := c.GetUint("id")
	
	// sessions, _ := dao.GetUserFriend(id)
	//res.OkWithData(c, "获取成功", sessions)
}

// NewFriendsInfo
// @Tags 好友模块
// @Summary 新好友信息
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/new [get]
func NewFriendsInfo(c *gin.Context) {	
	id := c.GetUint("id")
	ids, _ := dao.GetAddUserSenderIdsByReveiverId(id)
	users := make([]model.User, 0)
	global.DB.Where("id in ?", ids).Find(&users)
	res.OkWithData(c, "好友查询成功", users)
}

// AddFriend
// @Tags 好友模块
// @Summary 添加好友
// @Param id formData string false "对方id"
// @Param content formData string false "请求消息内容"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/add [post]
func AddFriend(c *gin.Context) {
	senderId := c.GetUint("id")
	var receiverId uint
	err := utils.QuerySwitch(c.PostForm("id"), receiverId)
  	if err != nil {
		log.Warn.Println("对方id格式错误")
		res.Err(c, "对方id格式错误")
		return
	}
  	if senderId == receiverId {
  		log.Warn.Println("不能加自己为好友")
  		res.Err(c, "不能加自己为好友")
  		return
  	}
  
  	user := &model.User {
		ID: receiverId,
	}

  	dao.GetUserById(user)
  	if user.Name == "" {
		res.Err(c, "未查询到该用户")
  		return
	}

	msg := &model.AddUserMessage {
		SenderId: senderId,
		ReceiverId: receiverId,
	}

	dao.CreateAddUserMessage(msg)
	res.Ok(c, "添加成功")
}

// DeleteFriend
// @Tags 好友模块
// @Summary 删除好友
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/delete [post]
func DeleteFriend(c *gin.Context) {
	
}

// AgreeNewFriend
// @Tags 好友模块
// @Summary 同意新好友的请求
// @Param id formData string false "对方id"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/add/agree [post]
func AgreeNewFriend(c *gin.Context) {
	var senderId uint
	err := utils.QuerySwitch(c.PostForm("id"), senderId)
	if err != nil {
		res.Err(c, "对方id错误")
		return
	}
	receiverId := c.GetUint("id")
	var session *model.UserSession
	if senderId < receiverId {
		session = &model.UserSession{
			SmallId: senderId,
			BigId: receiverId,
		}
	} else {
		session = &model.UserSession{
			SmallId: receiverId,
			BigId: senderId,
		}
	}
	// 开启事务
	global.DB.Begin()
	// 删除新好友表的数据
	db := dao.DeleteAddUserMessage(senderId, receiverId)
	if db.RowsAffected != 1 {
		// 出现错误, 回滚数据
		res.Err(c, "对方数据有误")
		global.DB.Rollback()
		return
	}
	// 新增用户会话表的数据
	db = dao.CreateUserSession(session)
	if db.RowsAffected != 1 {
		// 出现错误, 回滚数据
		global.DB.Rollback()
		res.Err(c, "同意失败")
		return
	}

	// 数据库操作无误, 提交数据
	global.DB.Commit()
	res.Ok(c, "已同意请求")
}

// RefuseNewFriend
// @Tags 好友模块
// @Summary 同意新好友的请求
// @Param id formData string false "对方id"
// @Success 200 {string} json{"code", "msg", "data"}
// @Router /friend/add/refuse [post]
func RefuseNewFriend(c *gin.Context) {
	var senderId uint
	err := utils.QuerySwitch(c.PostForm("id"), senderId)
	if err != nil {
		res.Err(c, "对方id错误")
		return
	}
	receiverId := c.GetUint("id")
	// 删除新好友表的数据
	db := dao.DeleteAddUserMessage(senderId, receiverId)
	if db.RowsAffected != 1 {
		// 出现错误, 回滚数据
		res.Err(c, "对方数据有误")
		global.DB.Rollback()
		return
	}

	res.Ok(c, "已拒绝请求")
}