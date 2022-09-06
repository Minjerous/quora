package question

import (
	"context"
	"github.com/jinzhu/copier"
	"quora/app/qa/cmd/rpc/pb"
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
		QuestionResp   *qa.GetQuestionByIdResp
		AnswerListResp *qa.GetAnswerListByPidResp
	)
	wg.Add(3)

	go func() {
		defer wg.Done()
		AnswerListResp, err = l.svcCtx.QA.GetAnswerListByQid(l.ctx, &pb.GetAnswerListByPidReq{
			Qid:      req.Id,
			PageSize: req.PageSize,
			OderBy:   req.OderBy,
			Page:     req.Page,
		})
	}()

	go func() {
		defer wg.Done()
		QuestionResp, err = l.svcCtx.QA.GetQuestionById(l.ctx, &pb.GetQuestionByIdReq{
			Id:       req.Id,
			PageSize: req.PageSize,
		})
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	var QuestionInfo types.QuestionInfo
	_ = copier.Copy(&QuestionInfo, QuestionResp.Question)

	var typesAnswerInfoList []types.AnswerInfo

	if len(AnswerListResp.Answer) > 0 {

		for _, Answer := range AnswerListResp.Answer {

			var typeHomestayOrder types.AnswerInfo
			_ = copier.Copy(&typeHomestayOrder, Answer)
			typesAnswerInfoList = append(typesAnswerInfoList, typeHomestayOrder)
		}
	}

	return &types.GetQuestionResp{
		AnswerLists:  typesAnswerInfoList,
		QuestionInfo: QuestionInfo,
		Message:      "成功",
		Status:       200,
	}, err

	return
}
