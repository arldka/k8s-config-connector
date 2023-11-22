// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/billing/v1/cloud_catalog.proto

package billingpb

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

// CloudCatalogClient is the client API for CloudCatalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudCatalogClient interface {
	// Lists all public cloud services.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// Lists all publicly available SKUs for a given cloud service.
	ListSkus(ctx context.Context, in *ListSkusRequest, opts ...grpc.CallOption) (*ListSkusResponse, error)
}

type cloudCatalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudCatalogClient(cc grpc.ClientConnInterface) CloudCatalogClient {
	return &cloudCatalogClient{cc}
}

func (c *cloudCatalogClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.billing.v1.CloudCatalog/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCatalogClient) ListSkus(ctx context.Context, in *ListSkusRequest, opts ...grpc.CallOption) (*ListSkusResponse, error) {
	out := new(ListSkusResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.billing.v1.CloudCatalog/ListSkus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudCatalogServer is the server API for CloudCatalog service.
// All implementations must embed UnimplementedCloudCatalogServer
// for forward compatibility
type CloudCatalogServer interface {
	// Lists all public cloud services.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// Lists all publicly available SKUs for a given cloud service.
	ListSkus(context.Context, *ListSkusRequest) (*ListSkusResponse, error)
	mustEmbedUnimplementedCloudCatalogServer()
}

// UnimplementedCloudCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCloudCatalogServer struct {
}

func (UnimplementedCloudCatalogServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedCloudCatalogServer) ListSkus(context.Context, *ListSkusRequest) (*ListSkusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSkus not implemented")
}
func (UnimplementedCloudCatalogServer) mustEmbedUnimplementedCloudCatalogServer() {}

// UnsafeCloudCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudCatalogServer will
// result in compilation errors.
type UnsafeCloudCatalogServer interface {
	mustEmbedUnimplementedCloudCatalogServer()
}

func RegisterCloudCatalogServer(s grpc.ServiceRegistrar, srv CloudCatalogServer) {
	s.RegisterService(&CloudCatalog_ServiceDesc, srv)
}

func _CloudCatalog_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCatalogServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.billing.v1.CloudCatalog/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCatalogServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCatalog_ListSkus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSkusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCatalogServer).ListSkus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.billing.v1.CloudCatalog/ListSkus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCatalogServer).ListSkus(ctx, req.(*ListSkusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudCatalog_ServiceDesc is the grpc.ServiceDesc for CloudCatalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudCatalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.billing.v1.CloudCatalog",
	HandlerType: (*CloudCatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _CloudCatalog_ListServices_Handler,
		},
		{
			MethodName: "ListSkus",
			Handler:    _CloudCatalog_ListSkus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/billing/v1/cloud_catalog.proto",
}
