package logic

import (
	"context"
	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"
	"quora/app/qa/newModel"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddQuestionLogic {
	return &AddQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddQuestionLogic) AddQuestion(in *pb.AddQuestionReq) (*pb.AddQuestionResp, error) {
	_, err := l.svcCtx.NewQuestionModel.Insert(l.ctx, &newModel.Question{
		Uid:       in.Uid,
		Topic:     in.Topic,
		Title:     in.Title,
		Value:     in.Value,
		CreatedAt: time.Unix(in.CreatedAt, 0),
		UpdateAt:  time.Unix(in.CreatedAt, 0),
	})

	if err != nil {
		return nil, err
	}

	return &pb.AddQuestionResp{}, nil
}
