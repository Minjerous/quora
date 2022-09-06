package logic

import (
	"context"
	"fmt"
	redisModel "quora/common/redis"

	"quora/app/action/cmd/rpc/internal/svc"
	"quora/app/action/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	//object: user question
	RS := redisModel.NewRedisSet(l.ctx, fmt.Sprintf("follow:uid:%v:%v:%v", in.Puid, in.Object, in.Pid), in.Uid, l.svcCtx.Redis)
	err := RS.SetAction()

	if err != nil {
		return nil, err
	}
	return &pb.FollowResp{Msg: "success"}, nil
}
