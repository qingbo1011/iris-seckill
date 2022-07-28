package controller

import (
	"iris-seckill/service"

	"github.com/kataras/iris/v12"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}
