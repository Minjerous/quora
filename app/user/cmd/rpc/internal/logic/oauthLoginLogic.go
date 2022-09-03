package logic

import (
	"context"

	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthLoginLogic) OauthLogin(in *pb.LoginReq) (*pb.OAuthLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OAuthLoginResp{}, nil
}
