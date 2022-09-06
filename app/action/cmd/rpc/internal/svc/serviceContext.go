package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"quora/app/action/cmd/rpc/internal/config"
	asynqModel "quora/common/asynq"

	redisModel "quora/common/redis"
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Client
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Redis:       redisModel.InitRedis(c.RedisDB.Addr, c.RedisDB.Pass),
		AsynqServer: asynqModel.NewServer(c.RedisDB.Addr, c.RedisDB.Pass),
	}
}
