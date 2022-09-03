package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AnswerModel = (*customAnswerModel)(nil)

type (
	// AnswerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAnswerModel.
	AnswerModel interface {
		answerModel
	}

	customAnswerModel struct {
		*defaultAnswerModel
	}
)

// NewAnswerModel returns a model for the database table.
func NewAnswerModel(conn sqlx.SqlConn, c cache.CacheConf) AnswerModel {
	return &customAnswerModel{
		defaultAnswerModel: newAnswerModel(conn, c),
	}
}
