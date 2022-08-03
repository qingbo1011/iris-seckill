package test

import (
	"fmt"
	"iris-seckill/conf"
	"iris-seckill/db/mysql"
	"iris-seckill/model"
	"iris-seckill/service"
	"testing"

	logging "github.com/sirupsen/logrus"
)

func init() {
	conf.Init("../conf/config.ini")
	mysql.Init()
}

func TestGetProductByID(t *testing.T) {
	var product model.Product
	mysql.MysqlDB.Where("id = ?", 1).First(&product)
	fmt.Println(product)
}

func TestGetAllProduct(t *testing.T) {
	var res []*model.Product
	mysql.MysqlDB.Find(&res)
	fmt.Println(res)
}

func TestDeleteProductByID(t *testing.T) {
	mysql.MysqlDB.Where("id = ?", 1).Delete(&model.Product{})
}

func TestInsertProduct(t *testing.T) {
	p := model.Product{
		ProductName:     "test6",
		ProductNum:      22,
		ProductImageUrl: "test6",
		ProductUrl:      "test6",
	}
	err := mysql.MysqlDB.Create(&p).Error
	if err != nil {
		logging.Info(err)
	}
}

func TestUpdateProduct(t *testing.T) {
	newProduct := model.Product{
		ProductName:     "new name",
		ProductNum:      111,
		ProductImageUrl: "new img",
		ProductUrl:      "new url",
	}
	err := mysql.MysqlDB.Model(&model.Product{}).Where("id = ?", 2).Updates(newProduct).Error
	if err != nil {
		logging.Info(err)
	}
}

func TestInsertOrder(t *testing.T) {
	order := model.Order{
		UserID:      1001,
		ProductID:   1001,
		OrderStatus: 2,
	}
	err := mysql.MysqlDB.Create(&order).Error
	if err != nil {
		logging.Info(err)
	}
}

func TestSelectUserByName(t *testing.T) {
	var user model.User
	err := mysql.MysqlDB.Where("user_name = ?", "张三").First(&user).Error
	if err != nil {
		logging.Info(err)
	}
	fmt.Println(user)
}

func TestProductSubNumberOne(t *testing.T) {
	productService := service.NewProductService()
	err := productService.SubNumberOne(8)
	if err != nil {
		logging.Info(err)
	}
}
