package asynqModel

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func NewServer(addr, passwd string) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr, Password: addr},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Info("asynq server exec task is failure  err : %+v \n", err)
				return true
			},
			Concurrency: 20,
		},
	)
}

func NewScheduler(addr, passwd string) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     addr,
			Password: passwd,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				logx.Info("Scheduler EnqueueErrorHandler  err : %+v , task : %+v", err, task)
			},
		})
}
