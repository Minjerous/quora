package logic

import (
	"context"
	"github.com/pkg/errors"
	"quora/app/user/cmd/rpc/user"
	"quora/app/user/model"

	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	User, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if err == model.ErrNotFound {
			return &pb.GetUserInfoResp{
				User: nil,
				OK:   false,
			}, errors.New("无效Uid")
		}
		return nil, err
	}
	var respUser user.UserInfo
	_ = copier.Copy(&respUser, User)

	return &pb.GetUserInfoResp{
		User: &respUser,
		OK:   false,
	}, errors.New("无效Uid")
}
