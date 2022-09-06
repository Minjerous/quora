package follow

import (
	"context"
	"quora/app/action/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {

	User := tool.GetUserFromCtx(l.ctx)
	var (
		FollowRpcResp *pb.FollowResp
	)

	FollowRpcResp, err = l.svcCtx.ActionRpc.Follow(l.ctx, &pb.FollowReq{
		Uid:    User.Id,
		Pid:    req.Pid,
		Object: req.Object,
	})

	if err != nil {
		return &types.FollowResp{
			Message: FollowRpcResp.Msg,
			Status:  400,
		}, err
	}

	return &types.FollowResp{
		Message: FollowRpcResp.Msg,
		Status:  200,
	}, nil
}
