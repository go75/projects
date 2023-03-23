package handler

import (
	"im/dao"
	"im/global"
	"im/log"
	"im/model"
	"im/utils"
	"im/ws/channel"
	"im/ws/entity"
	"im/ws/manager"
)

// 获取群聊会话
func GetGroupSession(r *entity.Request) {
	msgs, _ := dao.QueryGroupMessage(r.ProcessId)
	if len(msgs) == 0 {
		return
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, msgs)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}

	global.Rd.Subscribe()
}

// 获取新群聊信息
func GetNewGroup(r *entity.Request) {
	ls, _ := dao.GetAddGroupMessagesByGroupId(r.ProcessId)
	if len(ls) == 0 {
		return
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 发送群聊消息
func SendGroupMsg(r *entity.Request) {
	// 群聊消息持久化
	msg := &model.GroupMessage {
		GroupId: r.ProcessId,
		SenderId: r.SenderId,
		Type: r.Type,
		Content: r.Payload,
	}
	dao.CreateGroupMessage(msg)
	// 发布消息
	session := &model.GroupSession{
		GroupId: r.ProcessId,
	}
	ls, _ := dao.QueryGroupSessionsByGroupId(session)
	for _, session := range ls {
		if session.UserId == r.SenderId {
			continue
		}
		conn := manager.Get(session.UserId)
		if conn != nil {
			utils.Send(session.UserId, r.ID, entity.Text, 0, msg)
		}
	}
}

// 获取群聊列表
func GetGroupList(r *entity.Request) {

	ls, _ := dao.QueryGroupSessionsByUserId(r.SenderId)
	if len(ls) == 0 {
		return
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 添加群聊
func AddGroup(r *entity.Request) {
	// 新群友消息持久化
	msg := &model.AddGroupMessage{
		SenderId: r.SenderId,
		GroupId: r.ProcessId,
	}
	dao.CreateAddGroupMessage(msg)
	ch := utils.Channel(channel.AddGroup, r.ProcessId)
	err := utils.Publish(ch, r.ID, entity.Text, r.ProcessId, msg)
	if err != nil {
		log.Warn.Println("消息发布失败: ", err)
		return
	}
}

// 通过群聊名称模糊查询群聊
func GetFuzzyGroupByGroupName(r *entity.Request) {
	ls, _ := dao.GetFuzzyGroupByGroupName(r.Payload)
	err := utils.Send(r.SenderId, r.ID, entity.Text, 0, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 同意新群友请求
func AgreeNewGroup(r *entity.Request) {
	var userID uint
	err := utils.QuerySwitch(r.Payload, userID)
	if err != nil {
		log.Warn.Println("消息负载转化失败: " + err.Error())
		return
	}
	var session = &model.GroupSession{
		GroupId: r.SenderId,
		UserId: userID,
	}

	msg := &model.AddGroupMessage {
		SenderId: userID,
		GroupId:  r.ProcessId,
	}
	// 开启事务
	global.DB.Begin()
	
	db := dao.DeleteAddGroupMessage(msg)
	if db.RowsAffected != 1 {
		// 事务回滚
		global.DB.Rollback()
		return
	}
	db = dao.CreateGroupSession(session)
	if db.RowsAffected != 1 {
		// 事务回滚
		global.DB.Rollback()
		return
	}

	// 提交事务
	global.DB.Commit()

	group := &model.Group {
		ID: r.ProcessId,
	}
	dao.GetGroupById(group)
	if group.Name == "" {
		log.Warn.Printf("id为%d的用户信息获取失败\n", group.ID)
		return
	}
	// 返回新群聊的信息
	err = utils.Send(userID, r.ID, entity.Text, 0, group)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}

// 拒绝新群友请求
func RefuseNewGroup(r *entity.Request) {
	msg := &model.AddGroupMessage{
		SenderId: r.ProcessId,
		GroupId: r.SenderId,
	}
	dao.DeleteAddGroupMessage(msg)
}

// 获取群聊聊天记录
func GetGroupMsgs(r *entity.Request) {
	ls, _ := dao.QueryGroupMessage(r.ProcessId)
	if len(ls) == 0 {
		return
	}
	err := utils.Send(r.SenderId, r.ID, entity.Text, r.ProcessId, ls)
	if err != nil {
		log.Warn.Println("发送websocket消息失败: ", err)
		return
	}
}