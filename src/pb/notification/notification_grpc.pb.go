// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: notification.proto

package notificationpb

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	GetByUserId(ctx context.Context, in *GetNotificationByUserIdRequest, opts ...grpc.CallOption) (*GetNotificationByUserIdResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetByUserId(ctx context.Context, in *GetNotificationByUserIdRequest, opts ...grpc.CallOption) (*GetNotificationByUserIdResponse, error) {
	out := new(GetNotificationByUserIdResponse)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/GetByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	Create(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	GetByUserId(context.Context, *GetNotificationByUserIdRequest) (*GetNotificationByUserIdResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Create(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNotificationServiceServer) GetByUserId(context.Context, *GetNotificationByUserIdRequest) (*GetNotificationByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserId not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Create(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/GetByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetByUserId(ctx, req.(*GetNotificationByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NotificationService_Create_Handler,
		},
		{
			MethodName: "GetByUserId",
			Handler:    _NotificationService_GetByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}