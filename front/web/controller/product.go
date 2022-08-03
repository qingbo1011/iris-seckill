package controller

import (
	"iris-seckill/model"
	"iris-seckill/mq/rabbit"
	"iris-seckill/service"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

var (
	//生成的Html保存目录
	htmlOutPath = "./front/web/htmlProductShow/"
	//静态文件模版目录
	templatePath = "./front/web/views/template/"
)

// GetGenerateHtml 生成前端静态页面
func (p *ProductController) GetGenerateHtml() {
	productString := p.Ctx.URLParam("productID")
	productID, err := strconv.Atoi(productString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	//1.获取模版
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//2.获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	//3.获取模版渲染数据
	product, err := p.ProductService.GetProductByID(uint(productID))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//4.生成静态文件
	generateStaticHtml(p.Ctx, contenstTmp, fileName, product)
}

// 生成html静态文件
func generateStaticHtml(ctx iris.Context, template *template.Template, fileName string, product *model.Product) {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}
	}
	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

// 判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

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

// GetOrder 下订单
func (p *ProductController) GetOrder() []byte {
	productString := p.Ctx.URLParam("productID")
	userString := p.Ctx.GetCookie("uid")
	productID, err := strconv.Atoi(productString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	userID, err := strconv.ParseUint(userString, 10, 0)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	rabbit.NewMessage(uint(userID), uint(productID)) // 创建消息体

	return []byte("true")

	/* 引入RabbitMQ的下单逻辑 */

	/* 旧的下单逻辑 */
	//product, err := p.ProductService.GetProductByID(uint(productID))
	//if err != nil {
	//	p.Ctx.Application().Logger().Debug(err)
	//}
	//
	//var orderID uint
	//showMessage := "抢购失败！"
	//// 判断商品数量是否满足需求
	//if product.ProductNum > 0 {
	//	// 扣除商品数量
	//	product.ProductNum = product.ProductNum - 1
	//	err := p.ProductService.UpdateProduct(product)
	//	if err != nil {
	//		p.Ctx.Application().Logger().Debug(err)
	//	}
	//	// 创建订单
	//	userID, err := strconv.Atoi(userString)
	//	if err != nil {
	//		p.Ctx.Application().Logger().Debug(err)
	//	}
	//	order := &model.Order{
	//		UserID:      uint(userID),
	//		ProductID:   uint(productID),
	//		OrderStatus: model.OrderSuccess,
	//	}
	//	// 新建订单
	//	orderID, err = p.OrderService.InsertOrder(order)
	//	if err != nil {
	//		p.Ctx.Application().Logger().Debug(err)
	//	} else {
	//		showMessage = "抢购成功！"
	//	}
	//}
	//return mvc.View{
	//	Layout: "shared/productLayout.html",
	//	Name:   "product/result.html",
	//	Data: iris.Map{
	//		"orderID":     orderID,
	//		"showMessage": showMessage,
	//	},
	//}
}
