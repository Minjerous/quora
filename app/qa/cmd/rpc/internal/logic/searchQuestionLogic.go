package logic

import (
	"context"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchQuestionLogic {
	return &SearchQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchQuestionLogic) SearchQuestion(in *pb.SearchQuestionReq) (*pb.SearchQuestionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchQuestionResp{}, nil
}
