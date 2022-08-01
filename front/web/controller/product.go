package controller

import (
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
	OrderService   service.IOrderService
	Session        *sessions.Session
}

// GetDetail 获取商品详情信息页面
func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductByID(2) // 这里product_id写死了
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}
