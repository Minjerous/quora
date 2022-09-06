package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MqRpcConf zrpc.RpcClientConf

	DB struct {
		DataSource string
	}
	MetaData struct {
		UserName string
		PassWord string
	}
	JwtAuth struct {
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
