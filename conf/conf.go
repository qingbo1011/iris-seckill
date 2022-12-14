package conf

import (
	"time"

	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	IrisAddrBackend string
	IrisAddrFront   string

	MysqlHost            string
	MysqlPort            string
	MysqlUser            string
	MysqlPassword        string
	MysqlName            string
	MysqlIsLog           bool
	MysqlIsSingularTable bool
	MysqlMaxIdleConns    int
	MysqlMaxOpenConns    int
	MysqlConnMaxLifetime time.Duration

	RedisHost string

	RabbitMQHost         string
	RabbitMQPort         string
	RabbitMQUser         string
	RabbitMQPassword     string
	RabbitMQVirtualHosts string
)

func Init(path string) {
	file, err := ini.Load(path)
	if err != nil {
		logging.Fatalln(err)
	}

	loadService(file)
	loadMysql(file)
	loadRedis(file)
	loadRabbitMQ(file)
}

func loadService(file *ini.File) {
	IrisAddrBackend = file.Section("service").Key("IrisAddrBackend").MustString("127.0.0.1:8080")
	IrisAddrFront = file.Section("service").Key("IrisAddrFront").MustString("127.0.0.1:8082")
}

func loadMysql(file *ini.File) {
	section, err := file.GetSection("mysql")
	if err != nil {
		logging.Fatalln(err)
	}
	MysqlHost = section.Key("MysqlHost").String()
	MysqlPort = section.Key("MysqlPort").String()
	MysqlUser = section.Key("MysqlUser").String()
	MysqlPassword = section.Key("MysqlPassword").String()
	MysqlName = section.Key("MysqlName").String()
	MysqlIsLog = section.Key("MysqlIsLog").MustBool(true)
	MysqlIsSingularTable = section.Key("MysqlIsSingularTable").MustBool(true)
	MysqlMaxIdleConns = section.Key("MysqlMaxIdleConns").MustInt(20)
	MysqlMaxOpenConns = section.Key("MysqlMaxOpenConns").MustInt(100)
	MysqlConnMaxLifetime = time.Duration(section.Key("MysqlConnMaxLifetime").MustInt(30)) * time.Second
}

func loadRedis(file *ini.File) {
	section, err := file.GetSection("redis")
	if err != nil {
		logging.Fatalln(err)
	}
	RedisHost = section.Key("RedisHost").String()
}

func loadRabbitMQ(file *ini.File) {
	section, err := file.GetSection("rabbit")
	if err != nil {
		logging.Fatalln(err)
	}
	RabbitMQHost = section.Key("RabbitMQHost").String()
	RabbitMQPort = section.Key("RabbitMQPort").String()
	RabbitMQUser = section.Key("RabbitMQUser").String()
	RabbitMQPassword = section.Key("RabbitMQPassword").String()
	RabbitMQVirtualHosts = section.Key("RabbitMQVirtualHosts").String()
}
