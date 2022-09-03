package logic

import (
	"context"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentReq) (*pb.DelCommentResp, error) {
	// todo: Trans
	err := l.svcCtx.CommentModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DelCommentResp{}, nil
}
