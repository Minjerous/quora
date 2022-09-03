package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/comment/cmd/rpc/internal/config"
	"quora/app/comment/model"
	userModel "quora/app/user/model"
)

type ServiceContext struct {
	Config            config.Config
	CommentModel      model.CommentModel
	ChildCommentModel model.ChildCommentModel
	UserInfo          userModel.UserBasicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		CommentModel:      model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ChildCommentModel: model.NewChildCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserInfo:          userModel.NewUserBasicModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
