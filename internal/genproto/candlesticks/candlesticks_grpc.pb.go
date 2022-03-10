// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: candlesticks.proto

package candlesticks

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

// CandlesticksServiceClient is the client API for CandlesticksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandlesticksServiceClient interface {
	ReadCandlesticks(ctx context.Context, in *ReadCandlesticksRequest, opts ...grpc.CallOption) (*ReadCandlesticksResponse, error)
}

type candlesticksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCandlesticksServiceClient(cc grpc.ClientConnInterface) CandlesticksServiceClient {
	return &candlesticksServiceClient{cc}
}

func (c *candlesticksServiceClient) ReadCandlesticks(ctx context.Context, in *ReadCandlesticksRequest, opts ...grpc.CallOption) (*ReadCandlesticksResponse, error) {
	out := new(ReadCandlesticksResponse)
	err := c.cc.Invoke(ctx, "/candlesticks.CandlesticksService/ReadCandlesticks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandlesticksServiceServer is the server API for CandlesticksService service.
// All implementations should embed UnimplementedCandlesticksServiceServer
// for forward compatibility
type CandlesticksServiceServer interface {
	ReadCandlesticks(context.Context, *ReadCandlesticksRequest) (*ReadCandlesticksResponse, error)
}

// UnimplementedCandlesticksServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCandlesticksServiceServer struct {
}

func (UnimplementedCandlesticksServiceServer) ReadCandlesticks(context.Context, *ReadCandlesticksRequest) (*ReadCandlesticksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCandlesticks not implemented")
}

// UnsafeCandlesticksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandlesticksServiceServer will
// result in compilation errors.
type UnsafeCandlesticksServiceServer interface {
	mustEmbedUnimplementedCandlesticksServiceServer()
}

func RegisterCandlesticksServiceServer(s grpc.ServiceRegistrar, srv CandlesticksServiceServer) {
	s.RegisterService(&CandlesticksService_ServiceDesc, srv)
}

func _CandlesticksService_ReadCandlesticks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCandlesticksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandlesticksServiceServer).ReadCandlesticks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/candlesticks.CandlesticksService/ReadCandlesticks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandlesticksServiceServer).ReadCandlesticks(ctx, req.(*ReadCandlesticksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CandlesticksService_ServiceDesc is the grpc.ServiceDesc for CandlesticksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CandlesticksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "candlesticks.CandlesticksService",
	HandlerType: (*CandlesticksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadCandlesticks",
			Handler:    _CandlesticksService_ReadCandlesticks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "candlesticks.proto",
}
