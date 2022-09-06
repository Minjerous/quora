package logic

import (
	"context"
	"fmt"
	"quora/app/action/cmd/rpc/internal/svc"
	"quora/app/action/cmd/rpc/pb"
	redisModel "quora/common/redis"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// redis +1  asynq 读取redis -> mysql  notice

func (l *LikeLogic) Like(in *pb.LikeReq) (*pb.LikeResp, error) {

	RS := redisModel.NewRedisSet(l.ctx, fmt.Sprintf("like:uid:%v:%v:%v", in.Puid, in.Object, in.Pid), in.Uid, l.svcCtx.Redis)

	//nsq notice Puid
	// task sync redis
	err := RS.SetAction()

	if err != nil {
		return nil, err
	}

	return &pb.LikeResp{}, nil
}
