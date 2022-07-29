package controller

import (
	"iris-seckill/model"
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

// GetAll 获取全部产品
func (p *ProductController) GetAll() (mvc.View, error) {
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

// PostUpdate 修改产品
func (p *ProductController) PostUpdate() {
	product := &model.Product{}
	p.Ctx.Request().ParseForm()
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}
