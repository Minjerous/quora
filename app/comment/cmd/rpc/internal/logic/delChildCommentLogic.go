package logic

import (
	"context"
	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelChildCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelChildCommentLogic {
	return &DelChildCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelChildCommentLogic) DelChildComment(in *pb.DelChildCommentReq) (*pb.DelChildCommentResp, error) {
	// todo: Trans
	err := l.svcCtx.ChildCommentModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DelChildCommentResp{}, nil
}
