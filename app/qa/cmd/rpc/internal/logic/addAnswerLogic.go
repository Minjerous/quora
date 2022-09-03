package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/qa/model"
	"strconv"
	"time"

	"quora/app/qa/cmd/rpc/internal/svc"
	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAnswerLogic {
	return &AddAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAnswerLogic) AddAnswer(in *pb.AddAnswerReq) (*pb.AddAnswerResp, error) {

	if err := l.svcCtx.AnswerModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		Rs, err := l.svcCtx.AnswerModel.Insert(l.ctx, &model.Answer{
			Uid:       in.Uid,
			Pid:       in.Pid,
			Type:      in.Type,
			Value:     in.Value,
			CreatedAt: time.Unix(in.CreatedAt, 0),
		}, session)
		if err != nil {
			return errors.Wrapf(errors.New("Insert DB ERR "), "AddAnswer db answer Insert err:%v,userId:%+v", err, in.Uid)
		}
		err = l.svcCtx.QuestionModel.AddQuestionNumber(l.ctx, in.Pid, "comments", session)

		if err != nil {
			return errors.Wrapf(errors.New("Update DB ERR "), "UpdateQuestion db question Update err:%v,userId:%+v,Qid:%+v", err, in.Uid, in.Pid)
		}

		Id, err := Rs.LastInsertId()

		if err != nil {
			return errors.Wrapf(errors.New("Get DB_Index ERR "), "err:%v,userId:%+v,Qid:%+v", err, in.Uid, in.Pid)
		}
		return errors.New(fmt.Sprintf("NO %v", Id))
	}); err.Error()[:2] != "NO" {
		return nil, err
	} else {
		Id, err := strconv.ParseInt(err.Error()[3:], 64, 10)

		if err != nil {
			return nil, errors.Wrapf(errors.New("Parse Int ERR "), "err:%v,userId:%+v,Qid:%+v", err, in.Uid, in.Pid)
		}

		return &pb.AddAnswerResp{Id: Id}, nil

	}

}
