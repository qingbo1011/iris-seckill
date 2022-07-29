package controller

import (
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

// GetAll 获取全部产品
func (p ProductController) GetAll() (mvc.View, error) {
	products, err := p.ProductService.GetAllProduct()
	if err != nil {
		return mvc.View{}, err
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"products": products,
		},
	}, nil
}
