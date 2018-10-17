/*
mq消费者例子
(1)新建一个mq实例，设置路由规则
(2)一直等待接收并消费消息
注：每个实例使用一个独立的sock
*/
package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const MqServerAddr = "amqp://guest:guest@192.168.0.233:5672/"

// exchange属性，和消息生产者的exchange对应
var (
	// 注：如果多个路由的ExchangeName和QueueName都相同，RoutingKey无效，会轮流接收接收消息
	Info    = &Exchange{ExchangeName: "logs_direct", QueueName: "queue_info", RoutingKey: "info", ExchangeType: "direct", AutoAck: true}
	Warning = &Exchange{ExchangeName: "logs_direct", QueueName: "queue_warning", RoutingKey: "warning", ExchangeType: "direct", AutoAck: false}
	Erro    = &Exchange{ExchangeName: "logs_direct", QueueName: "queue_error", RoutingKey: "error", ExchangeType: "direct", AutoAck: false}
)

// 消费消息接口
type MqConsumer interface {
	MqConsume(hf HandleFunc)
}

// rabbitMQ最小管理单元exchange，每个exchange可以绑定多个队列，队列之间由RoutingKey区分
type Exchange struct {
	ExchangeName string // exchange名称
	QueueName    string // 队列名称
	RoutingKey   string // 路由key
	ExchangeType string // exchange类型，只有三种"fanout"、"direct"、"topic"
	AutoAck      bool   // true表示忽略，false表示手动ack
}

// mq消费对象
type mqConsume struct {
	addr     string           // mq服务器地址
	conn     *amqp.Connection // 连接
	ch       *amqp.Channel    // 通道
	exchange *Exchange        // mq路由规则
	stopChan chan bool        // 退出订阅
}

// 新建对象
func NewMqConsume(addr string, exchange *Exchange) MqConsumer {
	return &mqConsume{addr: addr, exchange: exchange}
}

// 连接mq服务器
func (mq *mqConsume) connect() error {
	conn, err := amqp.Dial(mq.addr)
	if err != nil {
		return err
	}
	mq.conn = conn
	return nil
}

// 创建一个通道
func (mq *mqConsume) createChannel() error {
	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}
	mq.ch = ch
	return nil
}

// 创建一个exchange，设置exchange相关属性
func (mq *mqConsume) createExchange() error {
	return mq.ch.ExchangeDeclare(
		mq.exchange.ExchangeName, // name
		mq.exchange.ExchangeType, // exchange类型"fanout"、"direct"、"topic"
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
}

// 创建一个queue，设置queue相关属性
func (mq *mqConsume) createQueue() error {
	_, err := mq.ch.QueueDeclare(
		mq.exchange.QueueName, // name
		true,  // durable 持久化
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

// 绑定exchange和queue
func (mq *mqConsume) bind() error {
	return mq.ch.QueueBind(
		mq.exchange.QueueName,    // queue name
		mq.exchange.RoutingKey,   // routing key
		mq.exchange.ExchangeName, // exchange
		false,
		nil,
	)
}

// 订阅消息
func (mq *mqConsume) subscribe(logic HandleFunc) error {
	fmt.Printf("RoutingKey=[%s], ExchangeType=[%s], Waiting for receive message.\n", mq.exchange.RoutingKey, mq.exchange.ExchangeType)
	msgs, err := mq.ch.Consume(
		mq.exchange.QueueName, // queue
		"",                  // consumer
		mq.exchange.AutoAck, // 手动还是自动回复，如果为手动回复，实际没有回复，该消息会mq服务器一直发送，直到有回复才会删除该消息
		false,               // exclusive
		false,               // no local
		false,               // no wait
		nil,                 // args
	)
	if err != nil {
		return err
	}

loop:
	for {
		select {
		case msg := <-msgs: // 从通道获取消息
			if !mq.exchange.AutoAck { // 判断是否需要手动回复
				if ok := logic(msg.Body); ok { // 正确逻辑处理，发出通知
					if err := mq.ch.Ack(msg.DeliveryTag, false); err != nil {
						fmt.Printf("DeliveryTag: %s", err.Error())
					} else {
						fmt.Println("已经手动通知RabbitMQ删除消息：", string(msg.Body))
					}
				}
			} else { // 忽略回复
				logic(msg.Body)
			}

		case <-mq.stopChan:
			break loop
		}
	}

	// 关闭通道和连接
	mq.ch.Close()
	mq.conn.Close()
	return nil
}

// 接收消息(通道阻塞)
func (mq *mqConsume) MqConsume(hf HandleFunc) {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("MqConsume:", err.Error())
		}
	}()

	if err = mq.connect(); err != nil { // 连接mq服务器
		return
	}
	if err = mq.createChannel(); err != nil { // 创建一个通道
		return
	}
	if err = mq.createExchange(); err != nil { // 创建一个exchange
		return
	}
	if err = mq.createQueue(); err != nil { // 创建一个队列
		return
	}
	if err = mq.bind(); err != nil { // 绑定exchange和queue
		return
	}
	if err = mq.subscribe(hf); err != nil { // 订阅消息
		return
	}
}

type HandleFunc func(message []byte) bool

// 逻辑处理函数
func HandleFuncLogic(message []byte) bool {
	fmt.Printf("Handle Message = %s\n", string(message))
	return true
}

func main() {
	exit := make(chan bool)
	println(" To exit press CTRL+C ")

	info := NewMqConsume(MqServerAddr, Info)
	go info.MqConsume(HandleFuncLogic)

	warning := NewMqConsume(MqServerAddr, Warning)
	go warning.MqConsume(HandleFuncLogic)

	erro := NewMqConsume(MqServerAddr, Erro)
	go erro.MqConsume(HandleFuncLogic)

	<-exit
}
