package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"quora/app/comment/cmd/api/internal/config"
	"quora/app/comment/cmd/rpc/comment"
	"quora/common/midleware"
)

type ServiceContext struct {
	Config     config.Config
	Auth       rest.Middleware
	CommentRpc comment.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Auth:       midleware.NewAuthMiddleware().Handle,
		CommentRpc: comment.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}
