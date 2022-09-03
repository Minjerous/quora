// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: comment.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	AddChildComment(ctx context.Context, in *AddChildCommentReq, opts ...grpc.CallOption) (*AddChildCommentResp, error)
	UpdateChildComment(ctx context.Context, in *UpdateChildCommentReq, opts ...grpc.CallOption) (*UpdateChildCommentResp, error)
	DelChildComment(ctx context.Context, in *DelChildCommentReq, opts ...grpc.CallOption) (*DelChildCommentResp, error)
	GetChildCommentByPId(ctx context.Context, in *GetChildCommentByPidReq, opts ...grpc.CallOption) (*GetChildCommentByPidResp, error)
	SearchChildComment(ctx context.Context, in *SearchChildCommentReq, opts ...grpc.CallOption) (*SearchChildCommentResp, error)
	//-----------------------comment-----------------------
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error)
	DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error)
	GetCommentByPid(ctx context.Context, in *GetCommentByPidReq, opts ...grpc.CallOption) (*GetCommentByPidResp, error)
	SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) AddChildComment(ctx context.Context, in *AddChildCommentReq, opts ...grpc.CallOption) (*AddChildCommentResp, error) {
	out := new(AddChildCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/AddChildComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) UpdateChildComment(ctx context.Context, in *UpdateChildCommentReq, opts ...grpc.CallOption) (*UpdateChildCommentResp, error) {
	out := new(UpdateChildCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/UpdateChildComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DelChildComment(ctx context.Context, in *DelChildCommentReq, opts ...grpc.CallOption) (*DelChildCommentResp, error) {
	out := new(DelChildCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/DelChildComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetChildCommentByPId(ctx context.Context, in *GetChildCommentByPidReq, opts ...grpc.CallOption) (*GetChildCommentByPidResp, error) {
	out := new(GetChildCommentByPidResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetChildCommentByPId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) SearchChildComment(ctx context.Context, in *SearchChildCommentReq, opts ...grpc.CallOption) (*SearchChildCommentResp, error) {
	out := new(SearchChildCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/SearchChildComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error) {
	out := new(UpdateCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentResp, error) {
	out := new(DelCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/DelComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentByPid(ctx context.Context, in *GetCommentByPidReq, opts ...grpc.CallOption) (*GetCommentByPidResp, error) {
	out := new(GetCommentByPidResp)
	err := c.cc.Invoke(ctx, "/pb.comment/GetCommentByPid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) SearchComment(ctx context.Context, in *SearchCommentReq, opts ...grpc.CallOption) (*SearchCommentResp, error) {
	out := new(SearchCommentResp)
	err := c.cc.Invoke(ctx, "/pb.comment/SearchComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	AddChildComment(context.Context, *AddChildCommentReq) (*AddChildCommentResp, error)
	UpdateChildComment(context.Context, *UpdateChildCommentReq) (*UpdateChildCommentResp, error)
	DelChildComment(context.Context, *DelChildCommentReq) (*DelChildCommentResp, error)
	GetChildCommentByPId(context.Context, *GetChildCommentByPidReq) (*GetChildCommentByPidResp, error)
	SearchChildComment(context.Context, *SearchChildCommentReq) (*SearchChildCommentResp, error)
	//-----------------------comment-----------------------
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error)
	DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error)
	GetCommentByPid(context.Context, *GetCommentByPidReq) (*GetCommentByPidResp, error)
	SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) AddChildComment(context.Context, *AddChildCommentReq) (*AddChildCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChildComment not implemented")
}
func (UnimplementedCommentServer) UpdateChildComment(context.Context, *UpdateChildCommentReq) (*UpdateChildCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChildComment not implemented")
}
func (UnimplementedCommentServer) DelChildComment(context.Context, *DelChildCommentReq) (*DelChildCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelChildComment not implemented")
}
func (UnimplementedCommentServer) GetChildCommentByPId(context.Context, *GetChildCommentByPidReq) (*GetChildCommentByPidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildCommentByPId not implemented")
}
func (UnimplementedCommentServer) SearchChildComment(context.Context, *SearchChildCommentReq) (*SearchChildCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChildComment not implemented")
}
func (UnimplementedCommentServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentServer) UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServer) DelComment(context.Context, *DelCommentReq) (*DelCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelComment not implemented")
}
func (UnimplementedCommentServer) GetCommentByPid(context.Context, *GetCommentByPidReq) (*GetCommentByPidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentByPid not implemented")
}
func (UnimplementedCommentServer) SearchComment(context.Context, *SearchCommentReq) (*SearchCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComment not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_AddChildComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChildCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddChildComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/AddChildComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddChildComment(ctx, req.(*AddChildCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_UpdateChildComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChildCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).UpdateChildComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/UpdateChildComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).UpdateChildComment(ctx, req.(*UpdateChildCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DelChildComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelChildCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DelChildComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/DelChildComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DelChildComment(ctx, req.(*DelChildCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetChildCommentByPId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChildCommentByPidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetChildCommentByPId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetChildCommentByPId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetChildCommentByPId(ctx, req.(*GetChildCommentByPidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_SearchChildComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChildCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).SearchChildComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/SearchChildComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).SearchChildComment(ctx, req.(*SearchChildCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).UpdateComment(ctx, req.(*UpdateCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_DelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).DelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/DelComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).DelComment(ctx, req.(*DelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentByPid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByPidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentByPid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/GetCommentByPid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentByPid(ctx, req.(*GetCommentByPidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_SearchComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).SearchComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.comment/SearchComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).SearchComment(ctx, req.(*SearchCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddChildComment",
			Handler:    _Comment_AddChildComment_Handler,
		},
		{
			MethodName: "UpdateChildComment",
			Handler:    _Comment_UpdateChildComment_Handler,
		},
		{
			MethodName: "DelChildComment",
			Handler:    _Comment_DelChildComment_Handler,
		},
		{
			MethodName: "GetChildCommentByPId",
			Handler:    _Comment_GetChildCommentByPId_Handler,
		},
		{
			MethodName: "SearchChildComment",
			Handler:    _Comment_SearchChildComment_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Comment_AddComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Comment_UpdateComment_Handler,
		},
		{
			MethodName: "DelComment",
			Handler:    _Comment_DelComment_Handler,
		},
		{
			MethodName: "GetCommentByPid",
			Handler:    _Comment_GetCommentByPid_Handler,
		},
		{
			MethodName: "SearchComment",
			Handler:    _Comment_SearchComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}