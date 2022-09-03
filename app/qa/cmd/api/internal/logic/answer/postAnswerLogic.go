package answer

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"
	"quora/app/qa/cmd/rpc/pb"
	"quora/common/tool"
	"time"
)

type PostAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAnswerLogic {
	return &PostAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostAnswerLogic) PostAnswer(req *types.PostAnswerReq) (resp *types.PostAnswerResp, err error) {

	user := tool.GetUserFromCtx(l.ctx)

	if !tool.VerifyType(req.Type) {
		return &types.PostAnswerResp{
			Message: "回答类型错误",
			Status:  400,
		}, err
	}

	_, err = l.svcCtx.QA.AddAnswer(l.ctx, &pb.AddAnswerReq{
		Uid:       user.Id,
		Pid:       req.Pid,
		Type:      req.Type,
		Value:     req.Value,
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &types.PostAnswerResp{
		Data:    *req,
		Message: "回答成功",
		Status:  200,
	}, err

}
