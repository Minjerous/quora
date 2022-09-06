package mq

import (
	"context"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type Config struct {
	Coon    *amqp.Connection
	Channel *amqp.Channel
}

func Connect(Addr string) *amqp.Channel {
	defer func() {
		if err := recover(); err != nil {
			logx.WithContext(context.Background()).Errorf("连接失败！")
		}
	}()
	conn, err := amqp.Dial(Addr)
	if err != nil {
		logx.WithContext(context.Background()).Errorf(err.Error())
	}
	chanel, err := conn.Channel()
	if err != nil {
		logx.WithContext(context.Background()).Errorf(err.Error())
	}
	return chanel
}

func (t *Config) Publish(message []byte, retry int, Addr string, key string) error {

	if t.Channel == nil {
		retry += 1
		if retry <= 2 {

			logx.WithContext(context.Background()).Errorf(fmt.Sprintf("3秒后尝试第%v次重连", retry))
			time.Sleep(3 * time.Second)
			t.Channel = Connect(Addr)
			err := t.Publish(message, retry, Addr, key)
			if err != nil {
				logx.WithContext(context.Background()).Errorf(err.Error())
				return errors.New("mq connection err")
			}
		}
		return errors.New("mq connection err")
	}
	err := t.Channel.Publish(
		"",
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "json/application",
			Body:        message,
		})

	return err
}
