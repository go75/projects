package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	//用户ID
	Id       uint      `gorm:"primary_key" json:"id"`
	Mobile   string `gorm:"type:varchar(20);not null;unique" json:"mobile"`
	Avatar   string    `gorm:"type:varchar(150)" json:"avatar"`
	Sex      string    `gorm:"type:varchar(2)" json:"sex"`       // 什么角色
	Nickname string    `gorm:"type:varchar(20)" json:"nickname"` // 什么角色
	Online   int       `gorm:"type:int(10)" json:"online"`       //是否在线
	Memo     string    `gorm:"type:varchar(140)" json:"memo"`    // 什么角色
	Createat   time.Time	`gorm:"type:datetime" json:"createat"`
}

func (user *User) TableName() string {
    return "user"
}