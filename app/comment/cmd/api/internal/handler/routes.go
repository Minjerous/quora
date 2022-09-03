// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "quora/app/comment/cmd/api/internal/handler/comment"
	"quora/app/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/comment/post",
					Handler: comment.PostCommentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/comment/child/post",
					Handler: comment.PostChildCommentHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/comment/delete",
					Handler: comment.DeleteCommentHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/comment/get",
				Handler: comment.GetCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/child/get",
				Handler: comment.GetChildCommentHandler(serverCtx),
			},
		},
	)
}
