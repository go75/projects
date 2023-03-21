package dao

import (
	"chat/global"
	"chat/model"

	"gorm.io/gorm"
)

func CreateUser(user *model.User) *gorm.DB {
	return global.DB.Create(user)
}

func GetUserByMobile(user *model.User) *gorm.DB {
	return global.DB.First(user)
}