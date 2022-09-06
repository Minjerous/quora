package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"quora/app/user/cmd/api/internal/config"
	"quora/app/user/cmd/rpc/user"
	"quora/common/midleware"
	redisModel "quora/common/redis"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	Redis   *redis.Client
	Auth    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Redis:   redisModel.InitRedis(c.RedisDB.Addr, c.RedisDB.Pass),
		Auth:    midleware.NewAuthMiddleware().Handle,
	}
}
