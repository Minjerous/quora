package user

import (
	"context"
	"quora/app/user/cmd/api/internal/svc"
	"quora/app/user/cmd/api/internal/types"
	"quora/app/user/cmd/rpc/pb"
	"quora/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	if len(req.Password) == 0 || len(req.Email) == 0 {
		return &types.LoginResp{
			Status: 40001,
			Msg:    "密码或密码不能为空",
		}, nil
	}

	if !tool.VerifyEmailFormat(req.Email) {
		return &types.LoginResp{
			Status: 40001,
			Msg:    "邮箱格式错误",
		}, nil
	}
	rpcResp, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResp{
			Status: 500,
			Msg:    "内部错误",
		}, nil
	}
	return &types.LoginResp{
		Status:       200,
		Msg:          "芜湖~ 登录 成功",
		AccessToken:  resp.AccessToken,
		AccessExpire: rpcResp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}, nil
}
