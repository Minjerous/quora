package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	MqRpcConf zrpc.RpcClientConf
	RedisDB   struct {
		Addr string
		Pass string
	}
}
