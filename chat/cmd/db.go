package main

import (
	"chat/config"
	"chat/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
  db, err := gorm.Open(mysql.Open(config.DbBsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // 迁移 schema
  db.AutoMigrate(
    &model.User{}, 
    &model.UserIdentity{},
    &model.Contact{},
    &model.Community{},
  )

  // // Create
  // user := &model.User{
	// Name: "peer",
  // }
  // db.Create(user)

  // // Read
  // var temp model.User
  // db.First(&temp, 1) // 根据整型主键查找
  // //   db.First(&temp, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

  // fmt.Println(temp)


  // // Update - 将 product 的 price 更新为 200
  // db.Model(&user).Update("Pwd", "1234")
}