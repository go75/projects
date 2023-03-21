package dao

import (
	"chat/global"
	"chat/model"

	"gorm.io/gorm"
)

func CreateUserIdentity(identity *model.UserIdentity) *gorm.DB {
	return global.DB.Create(&identity)
}

func GetUserIdentity(identity *model.UserIdentity) *gorm.DB {
	return global.DB.Where("mobile = ?", identity.Mobile).First(identity)
}