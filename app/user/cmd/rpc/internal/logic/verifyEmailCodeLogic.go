package logic

import (
	"context"
	"fmt"
	redisModel "quora/common/redis"

	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailCodeLogic {
	return &VerifyEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyEmailCodeLogic) VerifyEmailCode(in *pb.VerifyEmailCodeReq) (*pb.VerifyEmailCodeResp, error) {
	rs := redisModel.NewRedisStore(fmt.Sprintf("Email:%s", in.Email), l.svcCtx.Redis, 120, l.ctx)
	redisCode, err := rs.GetRedisValue()
	if err != nil {
		return nil, err
	}
	if redisCode == in.Code {
		return &pb.VerifyEmailCodeResp{OK: true}, nil
	} else {
		return &pb.VerifyEmailCodeResp{OK: false}, nil

	}
	return &pb.VerifyEmailCodeResp{OK: false}, nil
}
