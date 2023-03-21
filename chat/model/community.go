package model

import "time"

type Community struct {
	Id         uint      `gorm:"primary_key" json:"id"`
	//名称
	Name   string 		`gorm:"type:varchar(30)" json:"name"`
	//群主ID
	Ownerid       uint	`json:"ownerid"`   // 什么角色
	//群logo
	Icon	   string 		`gorm:"type:varchar(250)" json:"icon"`
	//como
	Cate      int	`gorm:"type:int(11)" json:"cate"`   // 什么角色
	//描述
	Memo    string	`gorm:"type:varchar(100)" json:"memo"`   // 什么角色
	//
	Createat   time.Time	`gorm:"type:datetime" json:"createat"`   // 什么角色
}
const (
		COMMUNITY_CATE_COM = 0x01
	)

func (community *Community) TableName() string {
	return "community"
}