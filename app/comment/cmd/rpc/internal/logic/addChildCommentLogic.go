package logic

import (
	"context"
	"quora/app/comment/model"
	"time"

	"quora/app/comment/cmd/rpc/internal/svc"
	"quora/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddChildCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChildCommentLogic {
	return &AddChildCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddChildCommentLogic) AddChildComment(in *pb.AddChildCommentReq) (*pb.AddChildCommentResp, error) {
	// todo: Trans
	Rs, err := l.svcCtx.ChildCommentModel.Insert(l.ctx, &model.ChildComment{
		Pid:      in.Pid,
		Content:  in.Content,
		Uid:      in.Uid,
		CreateAt: time.Unix(in.CreateAt, 0),
		ReplyUid: in.ReplyUid,
		Ip:       "",
		ReplyCid: in.ReplyCid,
	})

	id, err := Rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &pb.AddChildCommentResp{Id: id}, nil
}
