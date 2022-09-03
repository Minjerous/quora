package answer

import (
	"context"
	"quora/app/qa/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAnswerLogic {
	return &DeleteAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAnswerLogic) DeleteAnswer(req *types.DeleteAnswerReq) (resp *types.DeleteQuestionResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	_, err = l.svcCtx.QA.DelAnswer(l.ctx, &pb.DelAnswerReq{Id: User.Id})

	if err != nil {
		return nil, err
	}

	return &types.DeleteQuestionResp{
		Message: "删除成功",
		Status:  200,
	}, err
}
