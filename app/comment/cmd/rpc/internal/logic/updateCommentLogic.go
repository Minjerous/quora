package logic

import (
	"context"
	"quora/app/comment/model"
	"time"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	// todo:  Trans And Err
	err := l.svcCtx.CommentModel.Update(l.ctx, &model.Comment{
		Content:  in.Content,
		CreateAt: time.Unix(in.CreateAt, 0),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateCommentResp{}, nil
}
