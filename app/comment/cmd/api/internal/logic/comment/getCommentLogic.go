package comment

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/comment/cmd/rpc/pb"

	"quora/app/comment/cmd/api/internal/svc"
	"quora/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentReq) (resp *types.GetCommentResp, err error) {
	CommentRpcResp, err := l.svcCtx.CommentRpc.GetCommentByPid(l.ctx, &pb.GetCommentByPidReq{Id: req.Pid})

	if err != nil {
		return nil, err
	}
	var typesComments []types.Comment

	if len(CommentRpcResp.Comments) > 0 {

		for _, Comment := range CommentRpcResp.Comments {

			var typesComment types.Comment
			_ = copier.Copy(&typesComment, Comment)

			typesComments = append(typesComments, typesComment)
		}
	}

	return &types.GetCommentResp{Data: typesComments}, nil

}
