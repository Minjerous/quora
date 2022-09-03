package logic

import (
	"context"
	"quora/app/qa/model"
	"time"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnswerLogic {
	return &UpdateAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAnswerLogic) UpdateAnswer(in *pb.UpdateAnswerReq) (*pb.UpdateAnswerResp, error) {
	err := l.svcCtx.AnswerModel.Update(l.ctx, &model.Answer{
		Value:    in.Value,
		UpdateAt: time.Unix(in.UpdateAt, 0),
	})

	if err != nil {
		return nil, err
	}
	return &pb.UpdateAnswerResp{}, nil
}
