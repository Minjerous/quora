package delay

import (
	"context"
	"github.com/hibiken/asynq"
	"quora/app/mq/asynq/job/internal/svc"
)

type DeferUserInfo struct {
	svcCtx *svc.ServiceContext
}

func NewDeferUserInfoHandler(svcCtx *svc.ServiceContext) *DeferUserInfo {
	return &DeferUserInfo{
		svcCtx: svcCtx,
	}
}

func (l *DeferUserInfo) ProcessTask(ctx context.Context, _ *asynq.Task) (err error) {
	return nil
}
