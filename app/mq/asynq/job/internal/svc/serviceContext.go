package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"quora/app/mq/asynq/job/internal/config"
	asynqModel "quora/common/asynq"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	Redis       *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      config.Config{},
		AsynqServer: asynqModel.NewServer(c.RedisDB.Addr, c.RedisDB.Pass),
	}
}
