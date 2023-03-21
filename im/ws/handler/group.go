package handler

import (
	"im/dao"
	"im/log"
	"im/model"
	"im/utils"
	"im/ws/channel"
	"im/ws/entity"
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
	ch := utils.Channel(channel.Group, msg.GroupId)
	err := utils.Publish(ch, r.ID, entity.Text, r.ProcessId, msg)
	if err != nil {
		log.Warn.Println("消息发布失败: ", err)
		return
	}
}

// 获取群聊列表
func GetGroupList(r *entity.Request) {

	ls, _ := dao.QueryGroupSessionsByUserId(r.ID)
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

}

// 同意新群友请求
func AgreeNewGroup(r *entity.Request) {

}

// 拒绝新群友请求
func RefuseNewGroup(r *entity.Request) {

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