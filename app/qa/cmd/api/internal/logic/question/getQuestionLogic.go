package question

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/qa/cmd/rpc/qa"
	"sync"

	"quora/app/qa/cmd/api/internal/svc"
	"quora/app/qa/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var wg sync.WaitGroup

type GetQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionLogic {
	return &GetQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionLogic) GetQuestion(req *types.GetQuestionReq) (resp *types.GetQuestionResp, err error) {
	var (
		QuestionResp *qa.GetQuestionByIdResp
		//AnswerListResp *qa.GetAnswerListByQidResp
	)
	//User := tool.GetUserFromCtx(l.ctx)

	wg.Wait()
	if err != nil {
		return nil, err
	}

	var QuestionInfo types.QuestionInfo
	_ = copier.Copy(&QuestionInfo, QuestionResp.Question)

	return &types.GetQuestionResp{
		AnswerLists:  []types.AnswerInfo{},
		QuestionInfo: QuestionInfo,
		Message:      "成功",
		Status:       200,
	}, err

	return
}
