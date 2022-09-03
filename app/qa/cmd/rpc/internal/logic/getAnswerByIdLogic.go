package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerByIdLogic {
	return &GetAnswerByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerByIdLogic) GetAnswerById(in *pb.GetAnswerByIdReq) (*pb.GetAnswerByIdResp, error) {
	Answer, err := l.svcCtx.AnswerModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	var respAnswer pb.Answer
	_ = copier.Copy(&respAnswer, Answer)

	return &pb.GetAnswerByIdResp{Answer: &respAnswer}, nil
}
