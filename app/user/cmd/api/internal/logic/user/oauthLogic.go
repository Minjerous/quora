package user

import (
	"context"
	"quora/app/user/cmd/api/internal/svc"
	"quora/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLogic {
	return &OauthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OauthLogic) Oauth() (resp *types.LoginResp, err error) {

	return
}
