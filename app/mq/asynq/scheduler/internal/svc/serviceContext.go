package svc

import (
	"github.com/hibiken/asynq"
	"quora/app/mq/asynq/scheduler/internal/config"
	asynqModel "quora/common/asynq"
)

type ServiceContext struct {
	Config config.Config

	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: asynqModel.NewScheduler(c.RedisDB.Addr, c.RedisDB.Pass),
	}
}
