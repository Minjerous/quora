package comment

import (
	"context"
	"quora/app/comment/cmd/rpc/pb"
	"quora/common/tool"
	"time"

	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCommentLogic {
	return &PostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostCommentLogic) PostComment(req *types.PostCommentReq) (resp *types.PostCommentResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)
	//TODO Trans
	CommentRpcResp, err := l.svcCtx.CommentRpc.AddComment(l.ctx, &pb.AddCommentReq{
		ResourceType: req.ResourceType,
		Uid:          User.Id,
		Content:      req.Content,
		Pid:          req.Pid,
		CreateAt:     time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}
	return &types.PostCommentResp{
		Data: *req,
		Id:   CommentRpcResp.GetId(),
	}, err
}
