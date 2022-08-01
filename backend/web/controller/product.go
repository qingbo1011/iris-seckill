package controller

import (
	"iris-seckill/model"
	"iris-seckill/request"
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

// GetAdd 显示添加产品的页面
func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

// PostAdd 添加商品（POST请求去向数据库加数据）
func (p *ProductController) PostAdd() {
	productReq := request.ProductReq{}
	err2 := p.Ctx.ReadForm(&productReq)
	if err2 != nil {
		p.Ctx.StopWithError(iris.StatusBadRequest, err2)
		return
	}
	product := model.Product{
		ProductName:     productReq.ProductName,
		ProductNum:      productReq.ProductNum,
		ProductImageUrl: productReq.ProductImageUrl,
		ProductUrl:      productReq.ProductUrl,
	}
	_, err := p.ProductService.InsertProduct(&product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}
