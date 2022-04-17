// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package orders_api_v1alpha

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

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	// NewOrder process a new Order from a supplier
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error)
	// GetOrders returns all orders
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	// GetOrder returns a specific Order by Id
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	// ProcessOrder updates the status of an order when the Driver begins processing the order at the Food Bank
	ProcessOrder(ctx context.Context, in *ProcessOrderRequest, opts ...grpc.CallOption) (*ProcessOrderResponse, error)
	// CompleteOrder updates the status of an order to Complete
	CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error)
	// ClaimOrder used by a Driver to take ownership of an order
	ClaimOrder(ctx context.Context, in *ClaimOrderRequest, opts ...grpc.CallOption) (*ClaimOrderResponse, error)
	// CancelOrder unassign a order from a driver
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error) {
	out := new(NewOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) ProcessOrder(ctx context.Context, in *ProcessOrderRequest, opts ...grpc.CallOption) (*ProcessOrderResponse, error) {
	out := new(ProcessOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/ProcessOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error) {
	out := new(CompleteOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/CompleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) ClaimOrder(ctx context.Context, in *ClaimOrderRequest, opts ...grpc.CallOption) (*ClaimOrderResponse, error) {
	out := new(ClaimOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/ClaimOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/food.orders.api.v1alpha.OrdersService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	// NewOrder process a new Order from a supplier
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	// GetOrders returns all orders
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	// GetOrder returns a specific Order by Id
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	// ProcessOrder updates the status of an order when the Driver begins processing the order at the Food Bank
	ProcessOrder(context.Context, *ProcessOrderRequest) (*ProcessOrderResponse, error)
	// CompleteOrder updates the status of an order to Complete
	CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error)
	// ClaimOrder used by a Driver to take ownership of an order
	ClaimOrder(context.Context, *ClaimOrderRequest) (*ClaimOrderResponse, error)
	// CancelOrder unassign a order from a driver
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrdersServiceServer) ProcessOrder(context.Context, *ProcessOrderRequest) (*ProcessOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessOrder not implemented")
}
func (UnimplementedOrdersServiceServer) CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedOrdersServiceServer) ClaimOrder(context.Context, *ClaimOrderRequest) (*ClaimOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimOrder not implemented")
}
func (UnimplementedOrdersServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_ProcessOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).ProcessOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/ProcessOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).ProcessOrder(ctx, req.(*ProcessOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/CompleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CompleteOrder(ctx, req.(*CompleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_ClaimOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).ClaimOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/ClaimOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).ClaimOrder(ctx, req.(*ClaimOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/food.orders.api.v1alpha.OrdersService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "food.orders.api.v1alpha.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _OrdersService_NewOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrdersService_GetOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrdersService_GetOrder_Handler,
		},
		{
			MethodName: "ProcessOrder",
			Handler:    _OrdersService_ProcessOrder_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrdersService_CompleteOrder_Handler,
		},
		{
			MethodName: "ClaimOrder",
			Handler:    _OrdersService_ClaimOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrdersService_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/food/orders/api/v1alpha/orders_service.proto",
}