package user

import (
	"context"
	"errors"
	"quora/app/user/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/user/cmd/api/internal/svc"
	"quora/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendEmailReq) (resp *types.SendEmailResp, err error) {
	if !tool.VerifyEmailFormat(req.Email) {
		return nil, errors.New("邮箱格式错误")
	}
	UserRpcResp, err := l.svcCtx.UserRpc.SendEmail(l.ctx, &pb.SendEmailReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.SendEmailResp{
		UserRpcResp.Code,
	}, nil
}
