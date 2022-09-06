package like

import (
	"context"
	"quora/app/action/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeReq) (resp *types.LikeResp, err error) {
	User := tool.GetUserFromCtx(l.ctx)

	LikeRpcResp, err := l.svcCtx.ActionRpc.Like(l.ctx, &pb.LikeReq{
		Uid:    User.Id,
		Pid:    req.Pid,
		Puid:   req.PUid,
		Object: req.Object,
	})

	if err != nil {
		return nil, err
	}

	return &types.LikeResp{
		Message: LikeRpcResp.Msg,
		Status:  200,
	}, err
}
