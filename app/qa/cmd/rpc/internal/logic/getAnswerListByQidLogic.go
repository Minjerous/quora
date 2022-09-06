package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerListByQidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerListByQidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerListByQidLogic {
	return &GetAnswerListByQidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerListByQidLogic) GetAnswerListByQid(in *pb.GetAnswerListByPidReq) (*pb.GetAnswerListByPidResp, error) {
	AnswerLists, err := l.svcCtx.AnswerModel.FindByPid(l.ctx, in.Page, in.PageSize, in.OderBy, in.Qid)

	if err != nil {
		return nil, err
	}
	var resp []*pb.Answer
	if len(AnswerLists) > 0 {
		for _, Answer := range AnswerLists {
			var answer pb.Answer
			_ = copier.Copy(&Answer, answer)
			answer.CreatedAt = Answer.CreatedAt.Unix()
			answer.UpdateAt = Answer.UpdateAt.Unix()
			resp = append(resp, &answer)
		}
	}
	return &pb.GetAnswerListByPidResp{Answer: resp}, nil
}
