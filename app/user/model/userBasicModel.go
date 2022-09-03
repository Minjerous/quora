package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserBasicModel = (*customUserBasicModel)(nil)

type (
	// UserBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBasicModel.
	UserBasicModel interface {
		userBasicModel
	}

	customUserBasicModel struct {
		*defaultUserBasicModel
	}
)

// NewUserBasicModel returns a model for the database table.
func NewUserBasicModel(conn sqlx.SqlConn, c cache.CacheConf) UserBasicModel {
	return &customUserBasicModel{
		defaultUserBasicModel: newUserBasicModel(conn, c),
	}
}
