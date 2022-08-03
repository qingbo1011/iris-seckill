package main

import (
	"iris-seckill/mq/rabbit/simple"
	"iris-seckill/service"
)

func main() {
	productService := service.NewProductService() // 创建productService
	orderService := service.NewOrderService()     // 创建orderService
	rabbitMQSimple := simple.NewRabbitMQSimple("sec-kill")

	rabbitMQSimple.ConsumeSimple(orderService, productService)
}
