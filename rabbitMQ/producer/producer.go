/*
mq生产者例子
(1)新建一个mq实例，设置路由规则
(2)发送消息
消息生产者不需要声明队列
每个实例独立一个socket比多个实例共用一个socket要快，前提是socket够用情况下
*/
package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const MqServerAddr = "amqp://guest:guest@192.168.0.233:5672/"

// exchange属性，对应消费消息客户端的exchange属性，但不需要指明队列名称
var (
	Info    = &MqExchange{ExchangeName: "logs_direct", RoutingKey: "info", ExchangeType: "direct"}
	Warning = &MqExchange{ExchangeName: "logs_direct", RoutingKey: "warning", ExchangeType: "direct"}
	Errors  = &MqExchange{ExchangeName: "logs_direct", RoutingKey: "error", ExchangeType: "direct"}
)

// 生产者接口
type MqProducer interface {
	MqPublish(message []byte) error
	Close()
}

// rabbitMQ最小管理单元exchange
type MqExchange struct {
	ExchangeName string // exchange名称
	RoutingKey   string // 路由key
	ExchangeType string // exchange类型，只有三种"fanout"、"direct"、"topic"
}

// mq生产对象
type mqProduce struct {
	addr     string           // mq服务器地址
	conn     *amqp.Connection // 连接
	ch       *amqp.Channel    // 通道
	exchange *MqExchange      // exchange
}

// 新建对象
func NewMqProduce(addr string, exchange *MqExchange) (MqProducer, error) {
	mq := &mqProduce{addr: addr, exchange: exchange}

	var err error
	if err = mq.connect(); err != nil { // 连接mq服务器
		return nil, err
	}
	if err = mq.createChannel(); err != nil { // 创建一个通道
		return nil, err
	}
	if err = mq.createExchange(); err != nil { // 创建一个exchange
		return nil, err
	}
	return mq, nil
}

// 连接mq服务器
func (mq *mqProduce) connect() error {
	conn, err := amqp.Dial(mq.addr)
	if err != nil {
		return err
	}
	mq.conn = conn
	return nil
}

// 创建一个通道
func (mq *mqProduce) createChannel() error {
	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}
	mq.ch = ch
	return nil
}

// 创建一个exchange，设置相关exchange属性
func (mq *mqProduce) createExchange() error {
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

// 发布消息
func (mq *mqProduce) MqPublish(message []byte) error {
	err := mq.ch.Publish( // 根据路由规则发送消息到相应的队列
		mq.exchange.ExchangeName, // exchange name
		mq.exchange.RoutingKey,   // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}

	//	fmt.Printf("Sent: %s\n\n", string(message))
	return nil
}

// 释放资源
func (mq *mqProduce) Close() {
	mq.ch.Close()
	mq.conn.Close()
}

func main() {
	// 新建生产者
	info, err := NewMqProduce(MqServerAddr, Info)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer info.Close()

	warning, err := NewMqProduce(MqServerAddr, Warning)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer warning.Close()

	erro, err := NewMqProduce(MqServerAddr, Errors)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer erro.Close()

	// 待发送的消息
	infoMessage := []byte("---info---")
	waringMessage := []byte("---warning---")
	errorMessage := []byte("---error---")
	count := 100000
	t := time.Now()
	// 发送消息
	for i := 0; i < count; i++ {
		info.MqPublish(infoMessage)
		warning.MqPublish(waringMessage)
		erro.MqPublish(errorMessage)
	}
	fmt.Printf("send %d messages spend time: %fs\n", count*3, time.Now().Sub(t).Seconds())
}
