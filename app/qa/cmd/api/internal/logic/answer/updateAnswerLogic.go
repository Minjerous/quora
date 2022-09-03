package answer

import (
	"context"
	"quora/app/qa/cmd/rpc/pb"
	"quora/common/tool"
	"time"

	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAnswerLogic {
	return &UpdateAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAnswerLogic) UpdateAnswer(req *types.UpdateAnswerReq) (resp *types.UpdateAnswerResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	_, err = l.svcCtx.QA.UpdateAnswer(l.ctx, &pb.UpdateAnswerReq{
		Uid:      User.Id,
		Pid:      req.Pid,
		Value:    req.Value,
		UpdateAt: time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateAnswerResp{
		Data:    *req,
		Message: "修改成功",
		Status:  200,
	}, err
}
