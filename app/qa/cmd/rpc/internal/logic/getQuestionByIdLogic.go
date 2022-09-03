package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"sync"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var wg sync.WaitGroup

type GetQuestionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuestionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionByIdLogic {
	return &GetQuestionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuestionByIdLogic) GetQuestionById(in *pb.GetQuestionByIdReq) (*pb.GetQuestionByIdResp, error) {

	Question, err := l.svcCtx.QuestionModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	var respQuestion pb.Question
	_ = copier.Copy(&respQuestion, Question)

	return &pb.GetQuestionByIdResp{Question: &respQuestion}, nil
}
