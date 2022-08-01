package controller

import (
	"iris-seckill/model"
	"iris-seckill/request"
	"iris-seckill/service"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

// GetAll 获取全部产品
func (p *ProductController) GetAll() mvc.View {
	products, err := p.ProductService.GetAllProduct()
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"products": products,
		},
	}
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
	err := p.Ctx.ReadForm(&productReq)
	if err != nil {
		p.Ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	product := model.Product{
		ProductName:     productReq.ProductName,
		ProductNum:      productReq.ProductNum,
		ProductImageUrl: productReq.ProductImageUrl,
		ProductUrl:      productReq.ProductUrl,
	}
	_, err = p.ProductService.InsertProduct(&product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// GetManager 获取修改产品页面
func (p *ProductController) GetManager() mvc.View {
	idStr := p.Ctx.URLParam("id")
	id, err := strconv.ParseUint(idStr, 10, 0) // string转10进制的uint
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(uint(id))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

// PostUpdate 修改产品
func (p *ProductController) PostUpdate() {
	productReq := request.ProductReq{}
	err := p.Ctx.ReadForm(&productReq)
	product := model.Product{
		ProductName:     productReq.ProductName,
		ProductNum:      productReq.ProductNum,
		ProductImageUrl: productReq.ProductImageUrl,
		ProductUrl:      productReq.ProductUrl,
	}
	err = p.ProductService.UpdateProduct(&product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// GetDelete 删除产品
func (p *ProductController) GetDelete() {
	idStr := p.Ctx.URLParam("id")
	id, err := strconv.ParseUint(idStr, 10, 0) // string转10进制的uint
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err = p.ProductService.DeleteProductByID(uint(id))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为：" + idStr)
	}
	p.Ctx.Application().Logger().Debug("删除商品成功，ID为：" + idStr)
	p.Ctx.Redirect("/product/all")
}
