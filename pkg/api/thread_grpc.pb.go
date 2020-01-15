// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/thread.proto

package api

import (
	context "context"
	models "github.com/moguchev/BD-Forum/pkg/api/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ThreadClient is the client API for Thread service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThreadClient interface {
	// Создание ветки
	//
	// Добавление новой ветки обсуждения на форум.
	ThreadCreate(ctx context.Context, in *ThreadCreateRequest, opts ...grpc.CallOption) (*models.Thread, error)
	// Получение информации о ветке обсуждения
	//
	// Получение информации о ветке обсуждения по его имени.
	ThreadGetOne(ctx context.Context, in *ThreadGetOneRequest, opts ...grpc.CallOption) (*models.Thread, error)
	// Сообщения данной ветви обсуждения
	//
	// Получение списка сообщений в данной ветке форуме.
	//
	// Сообщения выводятся отсортированные по дате создания.
	ThreadGetPosts(ctx context.Context, in *ThreadGetPostsRequest, opts ...grpc.CallOption) (*models.Thread, error)
	// Обновление ветки
	//
	// Обновление ветки обсуждения на форуме.
	ThreadUpdate(ctx context.Context, in *ThreadUpdateRequest, opts ...grpc.CallOption) (*models.Thread, error)
	// Проголосовать за ветвь обсуждения
	//
	// Изменение голоса за ветвь обсуждения.
	//
	// Один пользователь учитывается только один раз и может изменить своё
	// мнение.
	ThreadVote(ctx context.Context, in *ThreadVoteRequest, opts ...grpc.CallOption) (*models.Thread, error)
}

type threadClient struct {
	cc grpc.ClientConnInterface
}

func NewThreadClient(cc grpc.ClientConnInterface) ThreadClient {
	return &threadClient{cc}
}

func (c *threadClient) ThreadCreate(ctx context.Context, in *ThreadCreateRequest, opts ...grpc.CallOption) (*models.Thread, error) {
	out := new(models.Thread)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Thread/ThreadCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) ThreadGetOne(ctx context.Context, in *ThreadGetOneRequest, opts ...grpc.CallOption) (*models.Thread, error) {
	out := new(models.Thread)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Thread/ThreadGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) ThreadGetPosts(ctx context.Context, in *ThreadGetPostsRequest, opts ...grpc.CallOption) (*models.Thread, error) {
	out := new(models.Thread)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Thread/ThreadGetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) ThreadUpdate(ctx context.Context, in *ThreadUpdateRequest, opts ...grpc.CallOption) (*models.Thread, error) {
	out := new(models.Thread)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Thread/ThreadUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) ThreadVote(ctx context.Context, in *ThreadVoteRequest, opts ...grpc.CallOption) (*models.Thread, error) {
	out := new(models.Thread)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Thread/ThreadVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThreadServer is the server API for Thread service.
// All implementations must embed UnimplementedThreadServer
// for forward compatibility
type ThreadServer interface {
	// Создание ветки
	//
	// Добавление новой ветки обсуждения на форум.
	ThreadCreate(context.Context, *ThreadCreateRequest) (*models.Thread, error)
	// Получение информации о ветке обсуждения
	//
	// Получение информации о ветке обсуждения по его имени.
	ThreadGetOne(context.Context, *ThreadGetOneRequest) (*models.Thread, error)
	// Сообщения данной ветви обсуждения
	//
	// Получение списка сообщений в данной ветке форуме.
	//
	// Сообщения выводятся отсортированные по дате создания.
	ThreadGetPosts(context.Context, *ThreadGetPostsRequest) (*models.Thread, error)
	// Обновление ветки
	//
	// Обновление ветки обсуждения на форуме.
	ThreadUpdate(context.Context, *ThreadUpdateRequest) (*models.Thread, error)
	// Проголосовать за ветвь обсуждения
	//
	// Изменение голоса за ветвь обсуждения.
	//
	// Один пользователь учитывается только один раз и может изменить своё
	// мнение.
	ThreadVote(context.Context, *ThreadVoteRequest) (*models.Thread, error)
	mustEmbedUnimplementedThreadServer()
}

// UnimplementedThreadServer must be embedded to have forward compatible implementations.
type UnimplementedThreadServer struct {
}

func (UnimplementedThreadServer) ThreadCreate(context.Context, *ThreadCreateRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadCreate not implemented")
}
func (UnimplementedThreadServer) ThreadGetOne(context.Context, *ThreadGetOneRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadGetOne not implemented")
}
func (UnimplementedThreadServer) ThreadGetPosts(context.Context, *ThreadGetPostsRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadGetPosts not implemented")
}
func (UnimplementedThreadServer) ThreadUpdate(context.Context, *ThreadUpdateRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadUpdate not implemented")
}
func (UnimplementedThreadServer) ThreadVote(context.Context, *ThreadVoteRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadVote not implemented")
}
func (UnimplementedThreadServer) mustEmbedUnimplementedThreadServer() {}

// UnsafeThreadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThreadServer will
// result in compilation errors.
type UnsafeThreadServer interface {
	mustEmbedUnimplementedThreadServer()
}

func RegisterThreadServer(s grpc.ServiceRegistrar, srv ThreadServer) {
	s.RegisterService(&Thread_ServiceDesc, srv)
}

func _Thread_ThreadCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).ThreadCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Thread/ThreadCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).ThreadCreate(ctx, req.(*ThreadCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_ThreadGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadGetOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).ThreadGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Thread/ThreadGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).ThreadGetOne(ctx, req.(*ThreadGetOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_ThreadGetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadGetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).ThreadGetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Thread/ThreadGetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).ThreadGetPosts(ctx, req.(*ThreadGetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_ThreadUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).ThreadUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Thread/ThreadUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).ThreadUpdate(ctx, req.(*ThreadUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_ThreadVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).ThreadVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Thread/ThreadVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).ThreadVote(ctx, req.(*ThreadVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Thread_ServiceDesc is the grpc.ServiceDesc for Thread service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Thread_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.moguchev.BD_Forum.api.Thread",
	HandlerType: (*ThreadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ThreadCreate",
			Handler:    _Thread_ThreadCreate_Handler,
		},
		{
			MethodName: "ThreadGetOne",
			Handler:    _Thread_ThreadGetOne_Handler,
		},
		{
			MethodName: "ThreadGetPosts",
			Handler:    _Thread_ThreadGetPosts_Handler,
		},
		{
			MethodName: "ThreadUpdate",
			Handler:    _Thread_ThreadUpdate_Handler,
		},
		{
			MethodName: "ThreadVote",
			Handler:    _Thread_ThreadVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/thread.proto",
}
