package logic

import (
	"context"
	"quora/app/qa/model"
	"time"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(in *pb.UpdateQuestionReq) (*pb.UpdateQuestionResp, error) {
	err := l.svcCtx.QuestionModel.Update(l.ctx, &model.Question{
		Topic:    in.Topic,
		Title:    in.Title,
		Value:    in.Value,
		UpdateAt: time.Unix(in.UpdateAt, 0),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateQuestionResp{}, nil
}
