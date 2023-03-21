package model

// 群聊消息数据
type GroupMessage struct {
	ID uint `gorm:"primary_key"`
	// 群聊名称
	GroupId uint
	// 发送方名称
	SenderId uint
	// 消息类型 文字为0, 图片为1, 音频为2
	Type uint
	// 消息内容
	Content string
}

func (m *GroupMessage) TableName() string {
	return "group_message"
}