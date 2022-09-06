package immed

import (
	"context"
	"github.com/hibiken/asynq"
	"quora/app/mq/asynq/job/internal/svc"
)

type SendMsgHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSendMsgHandler(svcCtx *svc.ServiceContext) *SendMsgHandler {
	return &SendMsgHandler{
		svcCtx: svcCtx,
	}
}

type SyncRedisToMysql struct {
	svcCtx *svc.ServiceContext
}

func NewSyncRedisToMysqlHandler(svcCtx *svc.ServiceContext) *SendMsgHandler {
	return &SendMsgHandler{
		svcCtx: svcCtx,
	}
}
func (l *SendMsgHandler) ProcessTask(ctx context.Context, _ *asynq.Task) (err error) {

	return nil
}

func (l SendMsgHandler) name() {

}
