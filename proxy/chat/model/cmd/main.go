package main

import (
	"proxy/chat/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
  db, err := gorm.Open(sqlite.Open("../../test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(&model.UserBasic{})

  // Create
  user := &model.UserBasic{
	Name: "peer",
  }
  db.Create(user)

  // Read
  var temp model.UserBasic
  db.First(&temp, 1) // 根据整型主键查找
//   db.First(&temp, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

  // Update - 将 product 的 price 更新为 200
  db.Model(&user).Update("Pwd", "1234")
}