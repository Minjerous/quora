package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/qa/cmd/rpc/internal/config"
	"quora/app/qa/model"
)

type ServiceContext struct {
	Config        config.Config
	QuestionModel model.QuestionModel
	AnswerModel   model.AnswerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		QuestionModel: model.NewQuestionModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		AnswerModel:   model.NewAnswerModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
