package repository

import (
	"todolist/internal/service"
	"todolist/pkg/util"
)

type User struct {
	Id uint32 `gorm:"primarykey"`
	Name string `gorm:"unique"` 
	Pwd string
}

func (u *User)ShowInfo(req *service.UserRequest) error {
	err := DB.Where("name=?", req.Name).First(u).Error;
	if err!=nil {
		return err
	}
	return nil
}

func (u *User) Create(req *service.UserRequest) error {
	u.Name = req.Name
	u.Pwd = util.MakePwd(req.Pwd)
	return DB.Create(u).Error
}

func (u *User) Build() *service.UserModel {
	return &service.UserModel{
		ID: uint32(u.Id),
		Name: u.Name,
	}
}