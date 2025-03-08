// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: commentsManager/commentsManager.proto

package cmv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CommentsManager_GetCommentById_FullMethodName         = "/github.chas3air.protos.commentsManager.CommentsManager/GetCommentById"
	CommentsManager_GetCommentsByArticleId_FullMethodName = "/github.chas3air.protos.commentsManager.CommentsManager/GetCommentsByArticleId"
	CommentsManager_Insert_FullMethodName                 = "/github.chas3air.protos.commentsManager.CommentsManager/Insert"
	CommentsManager_Delete_FullMethodName                 = "/github.chas3air.protos.commentsManager.CommentsManager/Delete"
)

// CommentsManagerClient is the client API for CommentsManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsManagerClient interface {
	GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error)
	GetCommentsByArticleId(ctx context.Context, in *GetCommentsByArticleIdRequest, opts ...grpc.CallOption) (*GetCommentsByArticleIdResponse, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type commentsManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsManagerClient(cc grpc.ClientConnInterface) CommentsManagerClient {
	return &commentsManagerClient{cc}
}

func (c *commentsManagerClient) GetCommentById(ctx context.Context, in *GetCommentByIdRequest, opts ...grpc.CallOption) (*GetCommentByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentByIdResponse)
	err := c.cc.Invoke(ctx, CommentsManager_GetCommentById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsManagerClient) GetCommentsByArticleId(ctx context.Context, in *GetCommentsByArticleIdRequest, opts ...grpc.CallOption) (*GetCommentsByArticleIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsByArticleIdResponse)
	err := c.cc.Invoke(ctx, CommentsManager_GetCommentsByArticleId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsManagerClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, CommentsManager_Insert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsManagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, CommentsManager_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsManagerServer is the server API for CommentsManager service.
// All implementations must embed UnimplementedCommentsManagerServer
// for forward compatibility.
type CommentsManagerServer interface {
	GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error)
	GetCommentsByArticleId(context.Context, *GetCommentsByArticleIdRequest) (*GetCommentsByArticleIdResponse, error)
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedCommentsManagerServer()
}

// UnimplementedCommentsManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentsManagerServer struct{}

func (UnimplementedCommentsManagerServer) GetCommentById(context.Context, *GetCommentByIdRequest) (*GetCommentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedCommentsManagerServer) GetCommentsByArticleId(context.Context, *GetCommentsByArticleIdRequest) (*GetCommentsByArticleIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByArticleId not implemented")
}
func (UnimplementedCommentsManagerServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedCommentsManagerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCommentsManagerServer) mustEmbedUnimplementedCommentsManagerServer() {}
func (UnimplementedCommentsManagerServer) testEmbeddedByValue()                         {}

// UnsafeCommentsManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsManagerServer will
// result in compilation errors.
type UnsafeCommentsManagerServer interface {
	mustEmbedUnimplementedCommentsManagerServer()
}

func RegisterCommentsManagerServer(s grpc.ServiceRegistrar, srv CommentsManagerServer) {
	// If the following call pancis, it indicates UnimplementedCommentsManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommentsManager_ServiceDesc, srv)
}

func _CommentsManager_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsManagerServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsManager_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsManagerServer).GetCommentById(ctx, req.(*GetCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsManager_GetCommentsByArticleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByArticleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsManagerServer).GetCommentsByArticleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsManager_GetCommentsByArticleId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsManagerServer).GetCommentsByArticleId(ctx, req.(*GetCommentsByArticleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsManager_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsManagerServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsManager_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsManagerServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsManager_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsManagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentsManager_ServiceDesc is the grpc.ServiceDesc for CommentsManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentsManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.chas3air.protos.commentsManager.CommentsManager",
	HandlerType: (*CommentsManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommentById",
			Handler:    _CommentsManager_GetCommentById_Handler,
		},
		{
			MethodName: "GetCommentsByArticleId",
			Handler:    _CommentsManager_GetCommentsByArticleId_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _CommentsManager_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CommentsManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commentsManager/commentsManager.proto",
}
