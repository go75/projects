package model

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Name string
	Pwd string
	Phone string
	Sex byte
	Status byte
	Introduce string
	Brithday int64
	Email string
	Identity string
	ClientAddr string
	LoginTime uint64
	HeartbeatTime uint64
	LogoutTime uint64
	IsOnline bool
	DeviceInfo string
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}

