package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"quora/app/qa/cmd/api/internal/config"
	"quora/app/qa/cmd/rpc/qa"
)

type ServiceContext struct {
	Config config.Config
	QA     qa.Qa
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		QA:     qa.NewQa(zrpc.MustNewClient(c.QaRpcConf)),
	}
}
