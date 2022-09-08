package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	DB struct {
		DataSource string
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	RedisDB struct {
		Addr string
		Pass string
	}

	Email struct {
		ServerEmail string
		PassWord    string
	}

	Cache cache.CacheConf
}
