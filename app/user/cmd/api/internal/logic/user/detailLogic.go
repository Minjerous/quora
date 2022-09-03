package user

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/user/cmd/api/internal/svc"
	"quora/app/user/cmd/api/internal/types"
	"quora/app/user/cmd/rpc/pb"
	"quora/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	User := tool.GetUserFromCtx(l.ctx)

	UserRpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: User.Id})

	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, UserRpcResp.User)
	return &types.UserInfoResp{
		Status:   200,
		Msg:      "成功",
		UserInfo: userInfo,
	}, nil
}
