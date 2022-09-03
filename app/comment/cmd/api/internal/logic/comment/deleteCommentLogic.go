package comment

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"
	"quora/app/comment/cmd/rpc/pb"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommmentReq) (resp *types.DeleteCommmentResp, err error) {

	switch req.Type {
	case "comment":
		_, err = l.svcCtx.CommentRpc.DelComment(l.ctx, &pb.DelCommentReq{Id: req.Id})
	case "child_comment":
		_, err = l.svcCtx.CommentRpc.DelChildComment(l.ctx, &pb.DelChildCommentReq{Id: req.Id})
	default:
		return &types.DeleteCommmentResp{
			Message: "对象类型错误",
			Status:  200,
		}, nil
	}

	return &types.DeleteCommmentResp{
		Message: "删除成功",
		Status:  200,
	}, nil
}
