package main

import (
	"testing"
	"im/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
	
	// 迁移 schema
	db.AutoMigrate(
	  &model.User{}, 
	  &model.UserIdentity{},
	  &model.Group{}, 
	  &model.UserSession{}, 
	  &model.GroupSession{}, 
	  &model.AddUserMessage{}, 
	  &model.AddGroupMessage{}, 
	  &model.UserMessage{}, 
	  &model.GroupMessage{},
	)
}