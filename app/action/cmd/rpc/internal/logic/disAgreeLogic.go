package logic

import (
	"context"
	"fmt"
	"quora/app/action/cmd/rpc/internal/svc"
	"quora/app/action/cmd/rpc/pb"
	redisModel "quora/common/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisAgreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisAgreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisAgreeLogic {
	return &DisAgreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisAgreeLogic) DisAgree(in *pb.AgreeReq) (*pb.AgreeResp, error) {

	RS := redisModel.NewRedisSet(l.ctx, fmt.Sprintf("agree:uid:%v:%v:%v", in.Puid, in.Object, in.Pid), in.Uid, l.svcCtx.Redis)
	val, err := RS.Conn.SRem(RS.Context, RS.Object, RS.Id).Result()
	_, err = RS.Conn.Decr(RS.Context, fmt.Sprintf("uid:%v:agreeadd", in.Puid)).Result()
	if val == 0 {
		return &pb.AgreeResp{Msg: "不能重复取消赞同"}, nil
	}
	if err != nil {
		return nil, err
	}
	return &pb.AgreeResp{Msg: "success"}, nil
}
