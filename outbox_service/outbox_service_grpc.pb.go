// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: outbox_service/outbox_service.proto

package outbox_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OutboxServiceClient is the client API for OutboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboxServiceClient interface {
	PollAgentEvent(ctx context.Context, in *PollAgentEventRequest, opts ...grpc.CallOption) (*PollAgentEventResponse, error)
	CommitAgentEvent(ctx context.Context, in *CommitAgentEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type outboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboxServiceClient(cc grpc.ClientConnInterface) OutboxServiceClient {
	return &outboxServiceClient{cc}
}

func (c *outboxServiceClient) PollAgentEvent(ctx context.Context, in *PollAgentEventRequest, opts ...grpc.CallOption) (*PollAgentEventResponse, error) {
	out := new(PollAgentEventResponse)
	err := c.cc.Invoke(ctx, "/outbox_service.OutboxService/PollAgentEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboxServiceClient) CommitAgentEvent(ctx context.Context, in *CommitAgentEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/outbox_service.OutboxService/CommitAgentEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutboxServiceServer is the server API for OutboxService service.
// All implementations must embed UnimplementedOutboxServiceServer
// for forward compatibility
type OutboxServiceServer interface {
	PollAgentEvent(context.Context, *PollAgentEventRequest) (*PollAgentEventResponse, error)
	CommitAgentEvent(context.Context, *CommitAgentEventRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOutboxServiceServer()
}

// UnimplementedOutboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOutboxServiceServer struct {
}

func (UnimplementedOutboxServiceServer) PollAgentEvent(context.Context, *PollAgentEventRequest) (*PollAgentEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PollAgentEvent not implemented")
}
func (UnimplementedOutboxServiceServer) CommitAgentEvent(context.Context, *CommitAgentEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitAgentEvent not implemented")
}
func (UnimplementedOutboxServiceServer) mustEmbedUnimplementedOutboxServiceServer() {}

// UnsafeOutboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboxServiceServer will
// result in compilation errors.
type UnsafeOutboxServiceServer interface {
	mustEmbedUnimplementedOutboxServiceServer()
}

func RegisterOutboxServiceServer(s grpc.ServiceRegistrar, srv OutboxServiceServer) {
	s.RegisterService(&OutboxService_ServiceDesc, srv)
}

func _OutboxService_PollAgentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollAgentEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).PollAgentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/outbox_service.OutboxService/PollAgentEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).PollAgentEvent(ctx, req.(*PollAgentEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboxService_CommitAgentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitAgentEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboxServiceServer).CommitAgentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/outbox_service.OutboxService/CommitAgentEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboxServiceServer).CommitAgentEvent(ctx, req.(*CommitAgentEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OutboxService_ServiceDesc is the grpc.ServiceDesc for OutboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "outbox_service.OutboxService",
	HandlerType: (*OutboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PollAgentEvent",
			Handler:    _OutboxService_PollAgentEvent_Handler,
		},
		{
			MethodName: "CommitAgentEvent",
			Handler:    _OutboxService_CommitAgentEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "outbox_service/outbox_service.proto",
}
