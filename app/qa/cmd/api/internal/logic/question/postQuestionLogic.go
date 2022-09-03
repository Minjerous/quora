package question

import (
	"context"
	"quora/app/qa/cmd/rpc/pb"
	"quora/common/tool"
	"time"

	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostQuestionLogic {
	return &PostQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostQuestionLogic) PostQuestion(req *types.PostQuestionReq) (resp *types.PostQuestionResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	_, err = l.svcCtx.QA.AddQuestion(l.ctx, &pb.AddQuestionReq{
		Uid:       User.Id,
		Topic:     req.Topic,
		Title:     req.Title,
		Value:     req.Value,
		DeletedAt: time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}
	return &types.PostQuestionResp{
		Data:    *req,
		Message: "发布成功",
		Status:  200,
	}, err

}
