package logic

import (
	"context"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerListByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerListByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerListByUidLogic {
	return &GetAnswerListByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerListByUidLogic) GetAnswerListByUid(in *pb.GetAnswerListByUidReq) (*pb.GetAnswerListByUidResp, error) {

	return &pb.GetAnswerListByUidResp{}, nil
}
