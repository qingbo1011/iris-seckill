package controller

import (
	"iris-seckill/model"
	"iris-seckill/service"
	"iris-seckill/util"
	"strconv"

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

// GetLogin 获取登录页面
func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

// PostLogin 用户登陆接口
func (c *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//2、验证账号密码正确
	user, ok := c.Service.IsPwdSuccess(userName, password)
	if !ok {
		return mvc.Response{
			Path: "/user/login",
		}
	}
	util.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(int64(user.ID), 10))
	c.Session.Set("userID", strconv.FormatInt(int64(user.ID), 10))
	return mvc.Response{
		Path: "/product/detail",
	}
}
