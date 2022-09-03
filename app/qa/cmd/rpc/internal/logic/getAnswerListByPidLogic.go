package logic

import (
	"context"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerListByQidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerListByQidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerListByQidLogic {
	return &GetAnswerListByQidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerListByQidLogic) GetAnswerListByQid(in *pb.GetAnswerListByPidReq) (*pb.GetAnswerListByPidResp, error) {

	return &pb.GetAnswerListByPidResp{}, nil
}
