package logic

import (
	"context"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchChildCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchChildCommentLogic {
	return &SearchChildCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchChildCommentLogic) SearchChildComment(in *pb.SearchChildCommentReq) (*pb.SearchChildCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchChildCommentResp{}, nil
}
