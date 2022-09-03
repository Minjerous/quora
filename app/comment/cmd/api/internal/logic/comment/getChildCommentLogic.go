package comment

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/comment/cmd/rpc/pb"

	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChildCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChildCommentLogic {
	return &GetChildCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChildCommentLogic) GetChildComment(req *types.GetChildCommentReq) (resp *types.GetChildCommentResp, err error) {
	ChillCommentRpcResp, err := l.svcCtx.CommentRpc.GetChildCommentByPId(l.ctx, &pb.GetChildCommentByPidReq{Pid: req.Pid})

	if err != nil {
		return nil, err
	}
	var typesChildComments []types.ChildComments

	if len(ChillCommentRpcResp.ChildComment) > 0 {

		for _, ChildComment := range ChillCommentRpcResp.ChildComment {

			var typesChildComment types.ChildComments
			_ = copier.Copy(&typesChildComment, ChildComment)

			typesChildComments = append(typesChildComments, typesChildComment)
		}
	}
	return &types.GetChildCommentResp{Data: typesChildComments}, nil
}
