package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/user/cmd/rpc/internal/config"
	"quora/app/user/model"
)

type ServiceContext struct {
	Config         config.Config
	UserBasicModel model.UserBasicModel
	Redis          *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserBasicModel: model.NewUserBasicModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		Redis:          model.InitRedis(c.RedisDB.Addr, c.RedisDB.Pass),
	}
}
