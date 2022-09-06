package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"quora/app/user/cmd/rpc/internal/config"
	"quora/app/user/model"
	redisModel "quora/common/redis"
)

type ServiceContext struct {
	Config         config.Config
	UserBasicModel model.UserBasicModel
	UserInfoModel  model.UserInfoModel
	Redis          *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserBasicModel: model.NewUserBasicModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserInfoModel:  model.NewUserInfoModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		Redis:          redisModel.InitRedis(c.RedisDB.Addr, c.RedisDB.Pass),
	}
}
