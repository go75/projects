package model

import (
	"time"
)

//好友和群都存在这个表里面
//可根据具体业务做拆分
type Contact struct {
	Id         uint      `gorm:"primary_key" json:"id"`
	//谁的10000
	Ownerid       uint	 `json:"ownid"`  // 记录是谁的
	//对端,10001
	Dstid        uint	 `json:"dstid"`  // 对端信息
	//
	Cate      int	     `json:"cate"`   // 什么类型
	Memo    string		 `gorm:"varchar(100)" json:"memo"`   // 备注
	//
	Createat   time.Time	`gorm:"type:datetime" json:"createat"`	// 创建时间
}

const (
		CONCAT_CATE_USER = 0x01
	    CONCAT_CATE_COMUNITY = 0x02
	)

func (contact *Contact) TableName() string {
    return "contact"
}