package user

import (
	"context"
	"fmt"
	"quora/app/user/cmd/api/internal/svc"
	"quora/app/user/cmd/api/internal/types"
	"quora/app/user/cmd/rpc/pb"
	"quora/app/user/model"
	"quora/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if !tool.VerifyEmailFormat(req.Email) {
		return &types.RegisterResp{
			Status: 40001,
			Msg:    "邮箱格式错误",
		}, nil
	}
	if len(req.UserName) > 16 || len(req.Password) < 6 || len(req.Password) > 16 {
		return &types.RegisterResp{
			Status: 40001,
			Msg:    "用户名or密码不规范",
		}, nil
	}
	rs := model.NewRedisStore(fmt.Sprintf("Email:%s", req.Email), l.svcCtx.Redis, 120, l.ctx)
	redisCode, err := rs.GetRedisValue()
	if err != nil {
		return &types.RegisterResp{
			Status: 400,
			Msg:    "内部错误",
		}, nil
	}
	if redisCode != req.Code {
		return &types.RegisterResp{
			Status: 200,
			Msg:    "验证码错误",
		}, nil
	}

	UserRpcResp, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		Email:    req.Email,
		PassWord: req.Password,
		Name:     req.UserName,
	})

	if err != nil {
		fmt.Println(err)
		return &types.RegisterResp{
			Status: 400,
			Msg:    "内部错误",
		}, nil
	}
	if UserRpcResp.OK == false {
		return &types.RegisterResp{
			Status: 200,
			Msg:    "昵称已经注册",
		}, nil
	}

	return &types.RegisterResp{
		Status:       200,
		Msg:          "注册成功",
		AccessToken:  UserRpcResp.AccessToken,
		AccessExpire: UserRpcResp.AccessExpire,
		RefreshAfter: UserRpcResp.RefreshAfter,
	}, nil
}
