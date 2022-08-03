package simple

import (
	"fmt"
	"iris-seckill/conf"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

// Rabbit RabbitMQ结构体
type Rabbit struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机名称
	Key       string // bind Key 名称
	MqUrl     string // 连接信息
}

// NewRabbitMQ 创建Rabbit结构体实例
func NewRabbitMQ(queueName, exchange, key string) *Rabbit {
	// rmqURL := "amqp://qingbo:qingbo@43.138.57.192:5672/qingbo"
	var builder strings.Builder
	s := []string{"amqp://", conf.RabbitMQUser, ":", conf.RabbitMQPassword, "@", conf.RabbitMQHost, "/", conf.RabbitMQVirtualHosts}
	for _, str := range s {
		builder.WriteString(str)
	}
	rmqURL := builder.String()
	return &Rabbit{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqUrl:     rmqURL,
	}
}

// Destroy 断开channel和connection
func (r Rabbit) Destroy() error {
	err := r.channel.Close()
	err = r.conn.Close()
	return err
}

// 错误处理函数
func (r Rabbit) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

// NewRabbitMQSimple 创建Simple模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *Rabbit {
	rabbitMQ := NewRabbitMQ(queueName, "", "") // 创建RabbitMQ实例
	var err error
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.MqUrl) // 获取connection
	rabbitMQ.failOnErr(err, "failed to connect RabbitMQ")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel() // 获取channel
	rabbitMQ.failOnErr(err, "failed to open a channel")
	return rabbitMQ
}

// PublishSimple Simple模式 生产者
func (r Rabbit) PublishSimple(msg string) {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否具有排他性
		false,       // 是否阻塞处理
		nil,         // 其他额外的属性
	)
	if err != nil {
		log.Println(err)
	}
	// 2.调用channel 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // 如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false, // 如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		log.Println(err)
	}
}

// ConsumeSimple Simple模式 消费者
func (r Rabbit) ConsumeSimple() {
	// 1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	queue, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否具有排他性
		false,       // 是否阻塞处理
		nil,         // 额外的属性
	)
	if err != nil {
		log.Println(err)
	}
	// 2.消费消息
	msgs, err := r.channel.Consume(
		queue.Name, // 队列名称
		"",         // 用来区分多个消费者
		true,       // 是否自动应答
		false,      // 是否独有
		false,      // 设置为true，表示不能将同一个Connection中生产者发送的消息传递给这个Connection中的消费者
		false,      // 队列是否阻塞
		nil,        // 额外的属性
	)
	if err != nil {
		log.Println(err)
	}

	// 3.启用协程处理消息
	forever := make(chan bool) // 开个channel阻塞住，让开启的协程能一直跑着
	go func() {
		for delivery := range msgs {
			// 消息逻辑处理，可以自行设计逻辑
			fmt.Println("Received a message:", string(delivery.Body))
		}
	}()
	fmt.Println(" [*] Waiting for messages.")
	<-forever
}
