// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/edn/v1/edn.proto

package v1

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

// EdnClient is the client API for Edn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdnClient interface {
	Receiver(ctx context.Context, in *ReceiverRequest, opts ...grpc.CallOption) (*ReceiverReply, error)
}

type ednClient struct {
	cc grpc.ClientConnInterface
}

func NewEdnClient(cc grpc.ClientConnInterface) EdnClient {
	return &ednClient{cc}
}

func (c *ednClient) Receiver(ctx context.Context, in *ReceiverRequest, opts ...grpc.CallOption) (*ReceiverReply, error) {
	out := new(ReceiverReply)
	err := c.cc.Invoke(ctx, "/api.edn.v1.Edn/Receiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdnServer is the server API for Edn service.
// All implementations must embed UnimplementedEdnServer
// for forward compatibility
type EdnServer interface {
	Receiver(context.Context, *ReceiverRequest) (*ReceiverReply, error)
	mustEmbedUnimplementedEdnServer()
}

// UnimplementedEdnServer must be embedded to have forward compatible implementations.
type UnimplementedEdnServer struct {
}

func (UnimplementedEdnServer) Receiver(context.Context, *ReceiverRequest) (*ReceiverReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receiver not implemented")
}
func (UnimplementedEdnServer) mustEmbedUnimplementedEdnServer() {}

// UnsafeEdnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdnServer will
// result in compilation errors.
type UnsafeEdnServer interface {
	mustEmbedUnimplementedEdnServer()
}

func RegisterEdnServer(s grpc.ServiceRegistrar, srv EdnServer) {
	s.RegisterService(&Edn_ServiceDesc, srv)
}

func _Edn_Receiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdnServer).Receiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.edn.v1.Edn/Receiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdnServer).Receiver(ctx, req.(*ReceiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Edn_ServiceDesc is the grpc.ServiceDesc for Edn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Edn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.edn.v1.Edn",
	HandlerType: (*EdnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receiver",
			Handler:    _Edn_Receiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/edn/v1/edn.proto",
}
