package main

import (
	"context"
	"iris-seckill/conf"
	"iris-seckill/db/mysql"
	"iris-seckill/front/web/controller"
	"iris-seckill/service"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	logging "github.com/sirupsen/logrus"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./front/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板
	app.HandleDir("/public", "./front/web/public")
	//访问生成好的html静态文件
	app.HandleDir("/html", "./front/web/htmlProductShow")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	sess := sessions.New(sessions.Config{
		Cookie:  "AdminCookie",
		Expires: 600 * time.Minute,
	})

	// 注册控制器
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userService := service.NewService()
	userApp := mvc.New(app.Party("/user"))
	userApp.Register(userService, ctx, sess.Start)
	userApp.Handle(new(controller.UserController))

	productService := service.NewProductService()
	orderService := service.NewOrderService()
	productParty := app.Party("/product")
	productApp := mvc.New(productParty)
	productApp.Register(productService, orderService)
	productApp.Handle(new(controller.ProductController))

	// 启动服务
	err := app.Run(
		iris.Addr(conf.IrisAddrFront),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	if err != nil {
		logging.Info(err)
	}
}

func init() {
	conf.Init("./conf/config.ini")
	mysql.Init()
}
