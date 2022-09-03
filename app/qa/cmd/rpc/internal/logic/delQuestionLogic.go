package logic

import (
	"context"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelQuestionLogic {
	return &DelQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelQuestionLogic) DelQuestion(in *pb.DelQuestionReq) (*pb.DelQuestionResp, error) {
	err := l.svcCtx.QuestionModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DelQuestionResp{}, nil
}
