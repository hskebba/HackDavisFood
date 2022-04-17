// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data_api_v1alpha

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

// DataApiServiceClient is the client API for DataApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataApiServiceClient interface {
	UpsertOrder(ctx context.Context, in *UpsertOrderRequest, opts ...grpc.CallOption) (*UpsertOrderResponse, error)
	InsertProduct(ctx context.Context, in *InsertProductRequest, opts ...grpc.CallOption) (*InsertProductResponse, error)
	UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error)
}

type dataApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataApiServiceClient(cc grpc.ClientConnInterface) DataApiServiceClient {
	return &dataApiServiceClient{cc}
}

func (c *dataApiServiceClient) UpsertOrder(ctx context.Context, in *UpsertOrderRequest, opts ...grpc.CallOption) (*UpsertOrderResponse, error) {
	out := new(UpsertOrderResponse)
	err := c.cc.Invoke(ctx, "/food.data_api.v1alpha.DataApiService/UpsertOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataApiServiceClient) InsertProduct(ctx context.Context, in *InsertProductRequest, opts ...grpc.CallOption) (*InsertProductResponse, error) {
	out := new(InsertProductResponse)
	err := c.cc.Invoke(ctx, "/food.data_api.v1alpha.DataApiService/InsertProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataApiServiceClient) UpdateInventory(ctx context.Context, in *UpdateInventoryRequest, opts ...grpc.CallOption) (*UpdateInventoryResponse, error) {
	out := new(UpdateInventoryResponse)
	err := c.cc.Invoke(ctx, "/food.data_api.v1alpha.DataApiService/UpdateInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataApiServiceServer is the server API for DataApiService service.
// All implementations must embed UnimplementedDataApiServiceServer
// for forward compatibility
type DataApiServiceServer interface {
	UpsertOrder(context.Context, *UpsertOrderRequest) (*UpsertOrderResponse, error)
	InsertProduct(context.Context, *InsertProductRequest) (*InsertProductResponse, error)
	UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error)
	mustEmbedUnimplementedDataApiServiceServer()
}

// UnimplementedDataApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataApiServiceServer struct {
}

func (UnimplementedDataApiServiceServer) UpsertOrder(context.Context, *UpsertOrderRequest) (*UpsertOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOrder not implemented")
}
func (UnimplementedDataApiServiceServer) InsertProduct(context.Context, *InsertProductRequest) (*InsertProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProduct not implemented")
}
func (UnimplementedDataApiServiceServer) UpdateInventory(context.Context, *UpdateInventoryRequest) (*UpdateInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInventory not implemented")
}
func (UnimplementedDataApiServiceServer) mustEmbedUnimplementedDataApiServiceServer() {}

// UnsafeDataApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataApiServiceServer will
// result in compilation errors.
type UnsafeDataApiServiceServer interface {
	mustEmbedUnimplementedDataApiServiceServer()
}

func RegisterDataApiServiceServer(s grpc.ServiceRegistrar, srv DataApiServiceServer) {
	s.RegisterService(&DataApiService_ServiceDesc, srv)
}

func _DataApiService_UpsertOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataApiServiceServer).UpsertOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.data_api.v1alpha.DataApiService/UpsertOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataApiServiceServer).UpsertOrder(ctx, req.(*UpsertOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataApiService_InsertProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataApiServiceServer).InsertProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.data_api.v1alpha.DataApiService/InsertProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataApiServiceServer).InsertProduct(ctx, req.(*InsertProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataApiService_UpdateInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataApiServiceServer).UpdateInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.data_api.v1alpha.DataApiService/UpdateInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataApiServiceServer).UpdateInventory(ctx, req.(*UpdateInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataApiService_ServiceDesc is the grpc.ServiceDesc for DataApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "food.data_api.v1alpha.DataApiService",
	HandlerType: (*DataApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertOrder",
			Handler:    _DataApiService_UpsertOrder_Handler,
		},
		{
			MethodName: "InsertProduct",
			Handler:    _DataApiService_InsertProduct_Handler,
		},
		{
			MethodName: "UpdateInventory",
			Handler:    _DataApiService_UpdateInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_api/food/data_api/v1alpha/data_api_service.proto",
}