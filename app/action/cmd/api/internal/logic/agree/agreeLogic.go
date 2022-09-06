package agree

import (
	"context"
	"github.com/pkg/errors"
	"quora/app/action/cmd/rpc/pb"
	"quora/common/tool"

	"quora/app/action/cmd/api/internal/svc"
	"quora/app/action/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgreeLogic {
	return &AgreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgreeLogic) Agree(req *types.AgreeReq) (resp *types.AgreeResq, err error) {
	User := tool.GetUserFromCtx(l.ctx)
	var AgreeRpcResp *pb.AgreeResp
	switch req.Opt {
	case 0:
		AgreeRpcResp, err = l.svcCtx.ActionRpc.DisAgree(l.ctx, &pb.AgreeReq{
			Uid:    User.Id,
			Pid:    req.Pid,
			Object: req.Object,
			Puid:   req.Puid,
		})

	case 1:
		AgreeRpcResp, err = l.svcCtx.ActionRpc.AddAgree(l.ctx, &pb.AgreeReq{
			Uid:    User.Id,
			Pid:    req.Pid,
			Object: req.Object,
			Puid:   req.Puid,
		})
	default:
		err = errors.New("opt is err")
	}

	if err != nil {
		return nil, err
	}

	return &types.AgreeResq{
		Message: AgreeRpcResp.Msg,
		Status:  200,
	}, err
}
