package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionModel = (*customQuestionModel)(nil)

type (
	// QuestionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionModel.
	QuestionModel interface {
		questionModel
	}

	customQuestionModel struct {
		*defaultQuestionModel
	}
)

// NewQuestionModel returns a model for the database table.
func NewQuestionModel(conn sqlx.SqlConn, c cache.CacheConf) QuestionModel {
	return &customQuestionModel{
		defaultQuestionModel: newQuestionModel(conn, c),
	}
}
