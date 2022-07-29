package main

import (
	"context"
	"iris-seckill/backend/web/controller"
	"iris-seckill/conf"
	"iris-seckill/db/mysql"
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	logging "github.com/sirupsen/logrus"
)

func main() {
	// 创建iris实例
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 注册模板
	tmplate := iris.HTML("./backend/web/view", ".html").Layout("share/layout.html").Reload(true)
	app.RegisterView(tmplate)
	// 设置模板目标
	app.HandleDir("/assets", "./backend/web/assets")
	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("share/error.html")
	})

	// 注册控制器
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	productService := service.NewProductService()
	productParty := app.Party("/product")
	productApp := mvc.New(productParty)
	productApp.Register(ctx, productService)
	productApp.Handle(new(controller.ProductController))

	// 启动服务
	err := app.Run(
		iris.Addr(conf.IrisAddr),
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
