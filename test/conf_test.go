package test

import (
	"fmt"
	"iris-seckill/conf"
	"testing"
)

func TestConfig(t *testing.T) {
	conf.Init("../conf/config.ini")
	fmt.Println(conf.IrisAddrBackend)
	fmt.Println(conf.MysqlHost)
	fmt.Println(conf.RedisHost)
	fmt.Println(conf.RabbitMQUser)
	fmt.Println(conf.RabbitMQHost)
}
