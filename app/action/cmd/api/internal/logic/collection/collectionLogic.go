package collection

import (
	"context"
	"quora/app/action/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectionLogic {
	return &CollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectionLogic) Collection(req *types.CollectionReq) (resp *types.CollectionResp, err error) {
	user := tool.GetUserFromCtx(l.ctx)

	CollectionRpcResp, err := l.svcCtx.ActionRpc.Collection(l.ctx, &pb.CollectionReq{
		Uid:    user.Id,
		Pid:    req.Pid,
		Object: req.Object,
	})

	if err != nil {
		return nil, err
	}
	return &types.CollectionResp{
		Message: CollectionRpcResp.Msg,
		Status:  200,
	}, nil
}
