package dao

import (
	"chat/global"
	"chat/model"

	"gorm.io/gorm"
)

func CreateContact(contact *model.Contact) *gorm.DB {
	return global.DB.Create(contact)
}

func GetContactsByOwneridAndCate(contact *model.Contact) ([]*model.Contact, *gorm.DB) {
	list := make([]*model.Contact, 0)
	db := global.DB.Where("ownerid = ? and cate = ?", contact.Ownerid, contact.Cate).Find(&list)
	return list, db
}