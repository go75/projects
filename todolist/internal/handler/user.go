package handler

import (
	"context"
	"net/http"
	"todolist/internal/repository"
	"todolist/internal/service"
)

type UserService struct {

}

func NewUserService() *UserService {
    return &UserService{}
}

// 用户登录
func (*UserService) UserLogin(c context.Context, req *service.UserRequest) (res *service.UserDetailResponse, err error) {
	user := repository.User{}
	res = &service.UserDetailResponse{
		Code: http.StatusOK,
	}
	err = user.ShowInfo(req)
	if err!=nil {
		res.Code = http.StatusBadRequest
		return
	}
	res.UserDetail = user.Build()
	return
}

// 用户注册
func (*UserService) UserRegist(c context.Context, req *service.UserRequest) (res *service.UserDetailResponse, err error) {
	user := repository.User{}
	res = &service.UserDetailResponse{
		Code: http.StatusOK,
	}
	err = user.Create(req)
	if err!=nil {
        res.Code = http.StatusBadRequest
        return
    }
	return
}