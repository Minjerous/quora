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

type UpdateQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(req *types.UpdateQuestionReq) (resp *types.UpdateQuestionResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	_, err = l.svcCtx.QA.UpdateQuestion(l.ctx, &pb.UpdateQuestionReq{
		Id:       req.Qid,
		Uid:      User.Id,
		Topic:    req.Topic,
		Title:    req.Title,
		Value:    req.Value,
		UpdateAt: time.Now().Unix(),
	})

	if err != nil {
		return &types.UpdateQuestionResp{
			Message: "修改失败",
			Status:  200,
		}, err
	}

	return &types.UpdateQuestionResp{
		Data:    *req,
		Message: "跟新成功",
		Status:  200,
	}, err
}
