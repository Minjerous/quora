package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"
	"quora/app/user/cmd/rpc/user"
	"quora/common/tool"
)

var ErrGenerateTokenError = errors.New("生成token失败")
var ErrUsernamePwdError = errors.New("账号或密码不正确")
var ErrFormatError = errors.New("格式错误")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	if flag := tool.VerifyEmailFormat(in.Email); !flag || len(in.Password) > 16 {

		return &pb.LoginResp{
			OK: false,
		}, ErrFormatError
	}
	User, err := l.svcCtx.UserBasicModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {

		return &pb.LoginResp{
			OK: false,
		}, ErrUsernamePwdError

	} else {
		if !tool.Match(User.Password, in.Password, User.Salt) {
			return &pb.LoginResp{
				OK: false,
			}, errors.New("账号或密码错误")
		}
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: User.Id,
		Name:   User.Name,
	})
	if err != nil {
		return &pb.LoginResp{
			OK: false,
		}, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", User.Id)
	}

	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
		OK:           true,
	}, nil
}
