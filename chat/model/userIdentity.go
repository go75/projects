package model

type UserIdentity struct {
	// 电话号
	Mobile   string `gorm:"type:varchar(20);not null;unique" json:"mobile"`
	// 密码
	Pwd string `gorm:"not null"`
	// 盐值
	Salt string `gorm:"not null"`
}

func (m *UserIdentity) TableName() string {
    return "user_identity"
}