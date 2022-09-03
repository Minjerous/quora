package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"
)

type DelAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAnswerLogic {
	return &DelAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelAnswerLogic) DelAnswer(in *pb.DelAnswerReq) (*pb.DelAnswerResp, error) {

	if err := l.svcCtx.AnswerModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		err := l.svcCtx.AnswerModel.Delete(l.ctx, in.Id, session)
		if err != nil {
			return errors.Wrapf(errors.New("Del DB ERR "), "err:%v,userId:%+v", err, in.Uid)
		}
		err = l.svcCtx.QuestionModel.SubQuestionNumber(l.ctx, in.Pid, "comments", session)

		if err != nil {
			return errors.Wrapf(errors.New("Update DB ERR "), "SubQuestionNumber db question Update err:%v,userId:%+v,Qid:%+v", err, in.Uid, in.Pid)
		}

		if err != nil {
			return errors.Wrapf(errors.New("Get DB_Index ERR "), "err:%v,userId:%+v,Qid:%+v", err, in.Uid, in.Pid)
		}
		return nil

	}); err != nil {
		return nil, err
	} else {
		return nil, err
	}
}
