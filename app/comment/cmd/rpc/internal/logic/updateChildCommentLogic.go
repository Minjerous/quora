package logic

import (
	"context"
	"quora/app/comment/model"
	"time"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChildCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChildCommentLogic {
	return &UpdateChildCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateChildCommentLogic) UpdateChildComment(in *pb.UpdateChildCommentReq) (*pb.UpdateChildCommentResp, error) {
	// todo: Trans And ERR
	err := l.svcCtx.ChildCommentModel.Update(l.ctx, &model.ChildComment{
		Content:  in.Content,
		CreateAt: time.Unix(in.CreateAt, 0),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdateChildCommentResp{}, nil
}
