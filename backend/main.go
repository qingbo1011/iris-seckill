package main

import (
	"iris-seckill/conf"

	"github.com/kataras/iris/v12"
	logging "github.com/sirupsen/logrus"
)

func main() {
	// 创建iris实例
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 注册模板
	tmplate := iris.HTML("./backend/web/assets", ".html").Layout("share/layout").Reload(true)
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
	conf.Init("../conf/config.ini")
}
