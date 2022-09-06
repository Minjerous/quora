package logic

import (
	"context"
	"fmt"
	redisModel "quora/common/redis"

	"quora/app/action/cmd/rpc/internal/svc"
	"quora/app/action/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAgreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAgreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAgreeLogic {
	return &AddAgreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAgreeLogic) AddAgree(in *pb.AgreeReq) (*pb.AgreeResp, error) {
	RS := redisModel.NewRedisSet(l.ctx, fmt.Sprintf("agree:uid:%v:%v:%v", in.Puid, in.Object, in.Pid), in.Uid, l.svcCtx.Redis)

	val, err := RS.Conn.SAdd(RS.Context, RS.Object, RS.Id).Result()
	RS.Conn.Incr(RS.Context, fmt.Sprintf("uid:%v:agreeadd", in.Puid))
	if val == 0 {
		return &pb.AgreeResp{Msg: "不能重复赞同"}, nil
	}
	if err != nil {
		return nil, err
	}

	return &pb.AgreeResp{Msg: "success"}, nil
}
