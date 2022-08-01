package service

import (
	"iris-seckill/db/mysql"
	"iris-seckill/model"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *model.User, isOk bool)
	AddUser(user *model.User) (userId uint, err error)
}

type UserService struct {
}

func NewService() IUserService {
	return &UserService{}
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (*model.User, bool) {
	var user model.User
	err := mysql.MysqlDB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, false
	}
	ok, err := user.CheckPassword(pwd)
	if err != nil {
		return nil, false
	}
	if !ok {
		return nil, false
	}
	return &user, true
}

// AddUser 用户注册
func (u *UserService) AddUser(user *model.User) (userId uint, err error) {
	err = mysql.MysqlDB.Create(user).Error
	return user.ID, err
}
