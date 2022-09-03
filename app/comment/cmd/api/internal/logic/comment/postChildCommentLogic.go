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

type PostChildCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostChildCommentLogic {
	return &PostChildCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostChildCommentLogic) PostChildComment(req *types.PostChildCommentReq) (resp *types.PostChlidCommentResp, err error) {

	User := tool.GetUserFromCtx(l.ctx)
	//TODO Trans
	ChildCommentRpcReps, err := l.svcCtx.CommentRpc.AddChildComment(l.ctx, &pb.AddChildCommentReq{
		Pid:      req.Pid,
		Content:  req.Content,
		Uid:      User.Id,
		CreateAt: time.Now().Unix(),
		ReplyUid: req.Repy_uid,
		ReplyCid: req.Repy_cid,
	})

	if err != nil {
		return nil, err
	}
	Id := ChildCommentRpcReps.GetId()

	return &types.PostChlidCommentResp{
		Data: *req,
		Id:   Id,
	}, nil
}
