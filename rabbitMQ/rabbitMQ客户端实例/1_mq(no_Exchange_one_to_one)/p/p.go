package main

import (
	"fmt"
	"log"
//	"time"
	"github.com/streadway/amqp"
)

// 处理错误
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@192.168.0.222:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 新建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 把消息内容发送到队列
	body := "hello ...... "
	for {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
//		time.Sleep(time.Millisecond)
	}

}
