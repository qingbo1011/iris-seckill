package controller

import (
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService service.IOrderService
}

// Get 获取全部订单页面
func (o *OrderController) Get() mvc.View {
	orders, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orders,
		},
	}
}
