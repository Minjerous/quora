package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"quora/app/action/cmd/api/internal/config"
	"quora/app/action/cmd/rpc/action"
)

type ServiceContext struct {
	Config    config.Config
	ActionRpc action.Action
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		ActionRpc: action.NewAction(zrpc.MustNewClient(c.ActionRpc)),
	}
}
