package mq

import (
	"fmt"
	"log"
	"strings"
)

func (t *Config) SubScribe(message []byte, retry int, Addr, exchange, key string) error {

	//创建直连交换机
	err := t.Channel.ExchangeDeclare(
		"user-follow", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	//创建队列
	q, err := t.Channel.QueueDeclare(
		"uid:101100", // name
		false,        // durable
		false,        // delete when unused
		true,         // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	err = t.Channel.QueueBind(
		q.Name,              // queue name
		"",                  // routing key
		"uid:555555:follow", // exchange
		false,
		nil,
	)

	fmt.Println(q.Name)
	uid := strings.Split(q.Name, ":")
	msgs, err := t.Channel.Consume(
		q.Name, // queue
		uid[1], // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
	return err
}
