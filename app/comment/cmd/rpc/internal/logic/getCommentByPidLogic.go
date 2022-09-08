package logic

import (
	"context"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByPidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByPidLogic {
	return &GetCommentByPidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByPidLogic) GetCommentByPid(in *pb.GetCommentByPidReq) (*pb.GetCommentByPidResp, error) {
	_, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	return &pb.GetCommentByPidResp{}, nil
}
