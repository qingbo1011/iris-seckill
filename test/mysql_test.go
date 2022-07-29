package test

import (
	"iris-seckill/conf"
	"iris-seckill/db/mysql"
	"iris-seckill/model"
	"testing"

	logging "github.com/sirupsen/logrus"
)

func init() {
	conf.Init("../conf/config.ini")
	mysql.Init()
}

func TestInsertProduct(t *testing.T) {
	p := model.Product{
		ProductID:       000000,
		ProductName:     "test",
		ProductNum:      1,
		ProductImageUrl: "test",
		ProductUrl:      "test",
	}
	err := mysql.MysqlDB.Create(&p).Error
	if err != nil {
		logging.Info(err)
	}
}
