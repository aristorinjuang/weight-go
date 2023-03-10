// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: weight.proto

package grpc

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

// WeightServiceClient is the client API for WeightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeightServiceClient interface {
	ListWeights(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Weights, error)
	CreateWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReadWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*Weight, error)
	UpdateWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type weightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeightServiceClient(cc grpc.ClientConnInterface) WeightServiceClient {
	return &weightServiceClient{cc}
}

func (c *weightServiceClient) ListWeights(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Weights, error) {
	out := new(Weights)
	err := c.cc.Invoke(ctx, "/WeightService/ListWeights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightServiceClient) CreateWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WeightService/CreateWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightServiceClient) ReadWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*Weight, error) {
	out := new(Weight)
	err := c.cc.Invoke(ctx, "/WeightService/ReadWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightServiceClient) UpdateWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WeightService/UpdateWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightServiceClient) DeleteWeight(ctx context.Context, in *WeightParams, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WeightService/DeleteWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeightServiceServer is the server API for WeightService service.
// All implementations should embed UnimplementedWeightServiceServer
// for forward compatibility
type WeightServiceServer interface {
	ListWeights(context.Context, *emptypb.Empty) (*Weights, error)
	CreateWeight(context.Context, *WeightParams) (*emptypb.Empty, error)
	ReadWeight(context.Context, *WeightParams) (*Weight, error)
	UpdateWeight(context.Context, *WeightParams) (*emptypb.Empty, error)
	DeleteWeight(context.Context, *WeightParams) (*emptypb.Empty, error)
}

// UnimplementedWeightServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWeightServiceServer struct {
}

func (UnimplementedWeightServiceServer) ListWeights(context.Context, *emptypb.Empty) (*Weights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWeights not implemented")
}
func (UnimplementedWeightServiceServer) CreateWeight(context.Context, *WeightParams) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWeight not implemented")
}
func (UnimplementedWeightServiceServer) ReadWeight(context.Context, *WeightParams) (*Weight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadWeight not implemented")
}
func (UnimplementedWeightServiceServer) UpdateWeight(context.Context, *WeightParams) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWeight not implemented")
}
func (UnimplementedWeightServiceServer) DeleteWeight(context.Context, *WeightParams) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWeight not implemented")
}

// UnsafeWeightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeightServiceServer will
// result in compilation errors.
type UnsafeWeightServiceServer interface {
	mustEmbedUnimplementedWeightServiceServer()
}

func RegisterWeightServiceServer(s grpc.ServiceRegistrar, srv WeightServiceServer) {
	s.RegisterService(&WeightService_ServiceDesc, srv)
}

func _WeightService_ListWeights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightServiceServer).ListWeights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeightService/ListWeights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightServiceServer).ListWeights(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightService_CreateWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeightParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightServiceServer).CreateWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeightService/CreateWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightServiceServer).CreateWeight(ctx, req.(*WeightParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightService_ReadWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeightParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightServiceServer).ReadWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeightService/ReadWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightServiceServer).ReadWeight(ctx, req.(*WeightParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightService_UpdateWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeightParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightServiceServer).UpdateWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeightService/UpdateWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightServiceServer).UpdateWeight(ctx, req.(*WeightParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightService_DeleteWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeightParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightServiceServer).DeleteWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeightService/DeleteWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightServiceServer).DeleteWeight(ctx, req.(*WeightParams))
	}
	return interceptor(ctx, in, info, handler)
}

// WeightService_ServiceDesc is the grpc.ServiceDesc for WeightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeightService",
	HandlerType: (*WeightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWeights",
			Handler:    _WeightService_ListWeights_Handler,
		},
		{
			MethodName: "CreateWeight",
			Handler:    _WeightService_CreateWeight_Handler,
		},
		{
			MethodName: "ReadWeight",
			Handler:    _WeightService_ReadWeight_Handler,
		},
		{
			MethodName: "UpdateWeight",
			Handler:    _WeightService_UpdateWeight_Handler,
		},
		{
			MethodName: "DeleteWeight",
			Handler:    _WeightService_DeleteWeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weight.proto",
}
