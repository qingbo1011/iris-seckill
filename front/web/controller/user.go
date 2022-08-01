package controller

import (
	"iris-seckill/model"
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Ctx     iris.Context
	Service service.IUserService
	Session *sessions.Session
}

// GetRegister 获取注册界面
func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

// PostRegister 用户注册（向数据库中新增数据）
func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	user := model.User{
		NickName: nickName,
		UserName: userName,
	}
	err := user.SetPassword(password)
	_, err = c.Service.AddUser(&user)
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}
