package logic

import (
	"context"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChildCommentByPIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChildCommentByPIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChildCommentByPIdLogic {
	return &GetChildCommentByPIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChildCommentByPIdLogic) GetChildCommentByPId(in *pb.GetChildCommentByPidReq) (*pb.GetChildCommentByPidResp, error) {
	// todo: Revise model function
	_, err := l.svcCtx.ChildCommentModel.FindOne(l.ctx, in.Pid)
	if err != nil {
		return nil, err
	}
	return &pb.GetChildCommentByPidResp{}, nil
}
