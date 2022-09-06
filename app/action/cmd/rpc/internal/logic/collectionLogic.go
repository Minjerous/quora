package logic

import (
	"context"
	"fmt"
	redisModel "quora/common/redis"

	"quora/app/action/cmd/rpc/internal/svc"
	"quora/app/action/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionLogic {
	return &CollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectionLogic) Collection(in *pb.CollectionReq) (*pb.CollectionResp, error) {
	RS := redisModel.NewRedisSet(l.ctx, fmt.Sprintf("collection:type:%v:id:%v", in.Object, in.Pid), in.Uid, l.svcCtx.Redis)
	val, err := RS.Conn.SAdd(RS.Context, RS.Object, RS.Id).Result()

	if val == 0 {
		return &pb.CollectionResp{Msg: "不能重复收藏"}, nil
	}
	if err != nil {
		return nil, err
	}

	return &pb.CollectionResp{}, nil
}
