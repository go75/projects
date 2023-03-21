package model

// 用户模块
type User struct {
	ID   uint `gorm:"primary_key"`
	// 名称
	Name string `gorm:"not null;unique"`
	// 个人介绍
	Introduce string
}

const (
	MAN = "M"
	WOMAN = "W"
	UNKNOW = "U"
)

func (u *User) TableName() string {
	return "user"
}