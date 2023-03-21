package dao

import (
	"im/global"
	"im/model"

	"gorm.io/gorm"
)

func CreateAddGroupMessage(msg *model.AddGroupMessage) *gorm.DB {
	return global.DB.Create(msg)
}

func GetAddGroupMessagesByGroupId(id uint) ([]*model.AddGroupMessage, *gorm.DB) {
	list := make([]*model.AddGroupMessage, 0)
	db := global.DB.Where("group_id", id).Find(&list)
	return list, db
}

func DeleteGroupMessage(msg *model.AddGroupMessage) *gorm.DB {
	return global.DB.Where("sender_id = ? and group_id = ?", msg.SenderId, msg.GroupId).Delete(msg)
}