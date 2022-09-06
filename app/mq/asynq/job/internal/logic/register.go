package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"quora/app/mq/asynq/job/internal/logic/immed"
	"quora/app/mq/asynq/job/internal/svc"
	"quora/app/mq/asynq/job/types"
)

type ProcessorJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessorJob(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessorJob {
	return &ProcessorJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessorJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.Handle(types.SyncLikeToMysql, immed.NewSendMsgHandler(l.svcCtx))

	return mux
}
