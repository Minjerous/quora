package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"quora/app/qa/cmd/api/internal/config"
	"quora/app/qa/cmd/rpc/qa"
	"quora/common/midleware"
)

type ServiceContext struct {
	Config config.Config
	Auth   func() func(ctx *gin.Context)
	QA     qa.Qa
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   midleware.NewAuthMiddleware().Handle,
		QA:     qa.NewQa(zrpc.MustNewClient(c.QaRpc)),
	}
}
