// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ProvisionServiceClient is the client API for ProvisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionServiceClient interface {
	Provision(ctx context.Context, in *ProvisionRequest, opts ...grpc.CallOption) (*ProvisionResponse, error)
}

type provisionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionServiceClient(cc grpc.ClientConnInterface) ProvisionServiceClient {
	return &provisionServiceClient{cc}
}

func (c *provisionServiceClient) Provision(ctx context.Context, in *ProvisionRequest, opts ...grpc.CallOption) (*ProvisionResponse, error) {
	out := new(ProvisionResponse)
	err := c.cc.Invoke(ctx, "/sidero.link.ProvisionService/Provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionServiceServer is the server API for ProvisionService service.
// All implementations must embed UnimplementedProvisionServiceServer
// for forward compatibility
type ProvisionServiceServer interface {
	Provision(context.Context, *ProvisionRequest) (*ProvisionResponse, error)
	mustEmbedUnimplementedProvisionServiceServer()
}

// UnimplementedProvisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProvisionServiceServer struct {
}

func (UnimplementedProvisionServiceServer) Provision(context.Context, *ProvisionRequest) (*ProvisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provision not implemented")
}
func (UnimplementedProvisionServiceServer) mustEmbedUnimplementedProvisionServiceServer() {}

// UnsafeProvisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionServiceServer will
// result in compilation errors.
type UnsafeProvisionServiceServer interface {
	mustEmbedUnimplementedProvisionServiceServer()
}

func RegisterProvisionServiceServer(s grpc.ServiceRegistrar, srv ProvisionServiceServer) {
	s.RegisterService(&ProvisionService_ServiceDesc, srv)
}

func _ProvisionService_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServiceServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidero.link.ProvisionService/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServiceServer).Provision(ctx, req.(*ProvisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvisionService_ServiceDesc is the grpc.ServiceDesc for ProvisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sidero.link.ProvisionService",
	HandlerType: (*ProvisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Provision",
			Handler:    _ProvisionService_Provision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siderolink/provision.proto",
}
