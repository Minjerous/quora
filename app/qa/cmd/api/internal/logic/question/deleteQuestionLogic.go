package question

import (
	"context"
	"quora/app/qa/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(req *types.DeleteQuestionReq) (resp *types.DeleteQuestionResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	_, err = l.svcCtx.QA.DelQuestion(l.ctx, &pb.DelQuestionReq{Id: req.Qid, Uid: User.Id})

	if err != nil {
		return nil, err
	}
	return &types.DeleteQuestionResp{
		Message: "删除成功",
		Status:  200,
	}, err
}
