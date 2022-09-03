package logic

import (
	"context"
	"github.com/pkg/errors"
	"quora/app/user/cmd/rpc/user"
	"quora/app/user/model"
	"quora/common/tool"
	"strconv"
	"time"

	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	Salt := strconv.FormatInt(time.Now().Unix(), 16)
	RS, err := l.svcCtx.UserBasicModel.Insert(l.ctx, &model.UserBasic{
		Name:      in.Name,
		Password:  tool.HashWithSalted(in.PassWord, Salt),
		Email:     in.Email,
		CreatedAt: time.Now(),
		Salt:      Salt,
	})

	id, err := RS.LastInsertId()

	if err != nil {
		if err.Error()[:len("Error 1062: Duplicate entry")] == "Error 1062: Duplicate entry" {
			return &pb.RegisterResp{
				OK: false,
			}, nil
		}
		return nil, err
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)

	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: id,
		Name:   in.Name,
	})
	if err != nil {
		return &pb.RegisterResp{
			OK: false,
		}, errors.Wrapf(ErrGenerateTokenError, "GenerateToken Emial : %d", in.Email)
	}
	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
		OK:           true,
	}, nil
}
