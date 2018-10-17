package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// 错误处理
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// 连接RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@192.168.0.201:5672/")
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
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 消费队列
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	// 接收消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	/*
		msgs, err := ch.Consume(
		  q.Name, // queue
		  "",     // consumer
		  false,  // auto-ack 取消自动回复，在消息中主动回复
		  false,  // exclusive
		  false,  // no-local
		  false,  // no-wait
		  nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
		  for d := range msgs {
		    log.Printf("Received a message: %s", d.Body)
		    dot_count := bytes.Count(d.Body, []byte("."))
		    t := time.Duration(dot_count)
		    time.Sleep(t * time.Second)
		    log.Printf("Done")
		    d.Ack(false)
		  }
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	*/
}
