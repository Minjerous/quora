package logic

import (
	"context"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAnswerLogic {
	return &SearchAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchAnswerLogic) SearchAnswer(in *pb.SearchAnswerReq) (*pb.SearchAnswerResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchAnswerResp{}, nil
}
