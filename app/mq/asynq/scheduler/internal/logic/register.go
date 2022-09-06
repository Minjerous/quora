package logic

import (
	"context"
	"quora/app/mq/asynq/scheduler/internal/svc"
)

//注册任务

type MqScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMqScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqScheduler {
	return &MqScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqScheduler) Register() {
	l.HandleRegisterTasks()
}
