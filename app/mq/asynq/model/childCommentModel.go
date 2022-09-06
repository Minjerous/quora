package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChildCommentModel = (*customChildCommentModel)(nil)

type (
	// ChildCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChildCommentModel.
	ChildCommentModel interface {
		childCommentModel
	}

	customChildCommentModel struct {
		*defaultChildCommentModel
	}
)

// NewChildCommentModel returns a model for the database table.
func NewChildCommentModel(conn sqlx.SqlConn, c cache.CacheConf) ChildCommentModel {
	return &customChildCommentModel{
		defaultChildCommentModel: newChildCommentModel(conn, c),
	}
}
