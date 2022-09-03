package logic

import (
	"context"
	"quora/app/qa/model"
	"time"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

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
	RS, err := l.svcCtx.QuestionModel.Insert(l.ctx, &model.Question{
		Uid:       in.Uid,
		Topic:     in.Topic,
		Title:     in.Title,
		Value:     in.Title,
		CreatedAt: time.Unix(in.CreatedAt, 0),
	})
	Id, err := RS.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &pb.AddQuestionResp{Id: Id}, nil
}
