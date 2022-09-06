// Code generated by goctl. DO NOT EDIT!
// Source: action.proto

package action

import (
	"context"

	"quora/app/action/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AgreeReq       = pb.AgreeReq
	AgreeResp      = pb.AgreeResp
	CollectionReq  = pb.CollectionReq
	CollectionResp = pb.CollectionResp
	FollowReq      = pb.FollowReq
	FollowResp     = pb.FollowResp
	LikeReq        = pb.LikeReq
	LikeResp       = pb.LikeResp

	Action interface {
		AddAgree(ctx context.Context, in *AgreeReq, opts ...grpc.CallOption) (*AgreeResp, error)
		DisAgree(ctx context.Context, in *AgreeReq, opts ...grpc.CallOption) (*AgreeResp, error)
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
		Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeResp, error)
		Collection(ctx context.Context, in *CollectionReq, opts ...grpc.CallOption) (*CollectionResp, error)
	}

	defaultAction struct {
		cli zrpc.Client
	}
)

func NewAction(cli zrpc.Client) Action {
	return &defaultAction{
		cli: cli,
	}
}

func (m *defaultAction) AddAgree(ctx context.Context, in *AgreeReq, opts ...grpc.CallOption) (*AgreeResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.AddAgree(ctx, in, opts...)
}

func (m *defaultAction) DisAgree(ctx context.Context, in *AgreeReq, opts ...grpc.CallOption) (*AgreeResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.DisAgree(ctx, in, opts...)
}

func (m *defaultAction) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultAction) Like(ctx context.Context, in *LikeReq, opts ...grpc.CallOption) (*LikeResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Like(ctx, in, opts...)
}

func (m *defaultAction) Collection(ctx context.Context, in *CollectionReq, opts ...grpc.CallOption) (*CollectionResp, error) {
	client := pb.NewActionClient(m.cli.Conn())
	return client.Collection(ctx, in, opts...)
}