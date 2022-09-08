package logic

import (
	"context"
	"quora/app/comment/model"
	"time"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------comment-----------------------
func (l *AddCommentLogic) AddComment(in *pb.AddCommentReq) (*pb.AddCommentResp, error) {

	RS, err := l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
		ResourceType: in.ResourceType,
		Uid:          in.Uid,
		Content:      in.Content,
		CreateAt:     time.Unix(in.CreateAt, 0),
		Ip:           "",
		Pid:          in.Pid,
	})
	id, err := RS.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.AddCommentResp{Id: id}, nil
}
