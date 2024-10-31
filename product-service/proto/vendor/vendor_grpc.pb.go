// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: product-service/proto/vendor/vendor.proto

package __

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

const (
	VendorService_CreateVendor_FullMethodName = "/vendor.VendorService/CreateVendor"
)

// VendorServiceClient is the client API for VendorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VendorServiceClient interface {
	CreateVendor(ctx context.Context, in *CreateVendorRequest, opts ...grpc.CallOption) (*CreateVendorResponse, error)
}

type vendorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVendorServiceClient(cc grpc.ClientConnInterface) VendorServiceClient {
	return &vendorServiceClient{cc}
}

func (c *vendorServiceClient) CreateVendor(ctx context.Context, in *CreateVendorRequest, opts ...grpc.CallOption) (*CreateVendorResponse, error) {
	out := new(CreateVendorResponse)
	err := c.cc.Invoke(ctx, VendorService_CreateVendor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VendorServiceServer is the server API for VendorService service.
// All implementations must embed UnimplementedVendorServiceServer
// for forward compatibility
type VendorServiceServer interface {
	CreateVendor(context.Context, *CreateVendorRequest) (*CreateVendorResponse, error)
	mustEmbedUnimplementedVendorServiceServer()
}

// UnimplementedVendorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVendorServiceServer struct {
}

func (UnimplementedVendorServiceServer) CreateVendor(context.Context, *CreateVendorRequest) (*CreateVendorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVendor not implemented")
}
func (UnimplementedVendorServiceServer) mustEmbedUnimplementedVendorServiceServer() {}

// UnsafeVendorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VendorServiceServer will
// result in compilation errors.
type UnsafeVendorServiceServer interface {
	mustEmbedUnimplementedVendorServiceServer()
}

func RegisterVendorServiceServer(s grpc.ServiceRegistrar, srv VendorServiceServer) {
	s.RegisterService(&VendorService_ServiceDesc, srv)
}

func _VendorService_CreateVendor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVendorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VendorServiceServer).CreateVendor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VendorService_CreateVendor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VendorServiceServer).CreateVendor(ctx, req.(*CreateVendorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VendorService_ServiceDesc is the grpc.ServiceDesc for VendorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VendorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vendor.VendorService",
	HandlerType: (*VendorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVendor",
			Handler:    _VendorService_CreateVendor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product-service/proto/vendor/vendor.proto",
}
