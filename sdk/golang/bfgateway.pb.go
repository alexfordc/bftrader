// Code generated by protoc-gen-go.
// source: bfgateway.proto
// DO NOT EDIT!

/*
Package bftrader_bfgateway is a generated protocol buffer package.

It is generated from these files:
	bfgateway.proto

It has these top-level messages:
*/
package bftrader_bfgateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bftrader "."
import google_protobuf "google/protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for BfGatewayService service

type BfGatewayServiceClient interface {
	// 请求建立到proxy的连接
	Connect(ctx context.Context, in *bftrader.BfConnectReq, opts ...grpc.CallOption) (BfGatewayService_ConnectClient, error)
	// 活跃检测
	Ping(ctx context.Context, in *bftrader.BfPingData, opts ...grpc.CallOption) (*bftrader.BfPingData, error)
	// 请求断开到proxy的连接
	Disconnect(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 获取交易合约信息
	GetContract(ctx context.Context, in *bftrader.BfGetContractReq, opts ...grpc.CallOption) (*bftrader.BfContractData, error)
	// 发单
	SendOrder(ctx context.Context, in *bftrader.BfSendOrderReq, opts ...grpc.CallOption) (*bftrader.BfSendOrderResp, error)
	// 撤单
	CancelOrder(ctx context.Context, in *bftrader.BfCancelOrderReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 查询账户信息
	QueryAccount(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
	// 查询持仓信息
	QueryPosition(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error)
}

type bfGatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewBfGatewayServiceClient(cc *grpc.ClientConn) BfGatewayServiceClient {
	return &bfGatewayServiceClient{cc}
}

func (c *bfGatewayServiceClient) Connect(ctx context.Context, in *bftrader.BfConnectReq, opts ...grpc.CallOption) (BfGatewayService_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfGatewayService_serviceDesc.Streams[0], c.cc, "/bftrader.bfgateway.BfGatewayService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfGatewayServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfGatewayService_ConnectClient interface {
	Recv() (*google_protobuf.Any, error)
	grpc.ClientStream
}

type bfGatewayServiceConnectClient struct {
	grpc.ClientStream
}

func (x *bfGatewayServiceConnectClient) Recv() (*google_protobuf.Any, error) {
	m := new(google_protobuf.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfGatewayServiceClient) Ping(ctx context.Context, in *bftrader.BfPingData, opts ...grpc.CallOption) (*bftrader.BfPingData, error) {
	out := new(bftrader.BfPingData)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) Disconnect(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/Disconnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) GetContract(ctx context.Context, in *bftrader.BfGetContractReq, opts ...grpc.CallOption) (*bftrader.BfContractData, error) {
	out := new(bftrader.BfContractData)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/GetContract", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) SendOrder(ctx context.Context, in *bftrader.BfSendOrderReq, opts ...grpc.CallOption) (*bftrader.BfSendOrderResp, error) {
	out := new(bftrader.BfSendOrderResp)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/SendOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) CancelOrder(ctx context.Context, in *bftrader.BfCancelOrderReq, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) QueryAccount(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/QueryAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfGatewayServiceClient) QueryPosition(ctx context.Context, in *bftrader.BfVoid, opts ...grpc.CallOption) (*bftrader.BfVoid, error) {
	out := new(bftrader.BfVoid)
	err := grpc.Invoke(ctx, "/bftrader.bfgateway.BfGatewayService/QueryPosition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BfGatewayService service

type BfGatewayServiceServer interface {
	// 请求建立到proxy的连接
	Connect(*bftrader.BfConnectReq, BfGatewayService_ConnectServer) error
	// 活跃检测
	Ping(context.Context, *bftrader.BfPingData) (*bftrader.BfPingData, error)
	// 请求断开到proxy的连接
	Disconnect(context.Context, *bftrader.BfVoid) (*bftrader.BfVoid, error)
	// 获取交易合约信息
	GetContract(context.Context, *bftrader.BfGetContractReq) (*bftrader.BfContractData, error)
	// 发单
	SendOrder(context.Context, *bftrader.BfSendOrderReq) (*bftrader.BfSendOrderResp, error)
	// 撤单
	CancelOrder(context.Context, *bftrader.BfCancelOrderReq) (*bftrader.BfVoid, error)
	// 查询账户信息
	QueryAccount(context.Context, *bftrader.BfVoid) (*bftrader.BfVoid, error)
	// 查询持仓信息
	QueryPosition(context.Context, *bftrader.BfVoid) (*bftrader.BfVoid, error)
}

func RegisterBfGatewayServiceServer(s *grpc.Server, srv BfGatewayServiceServer) {
	s.RegisterService(&_BfGatewayService_serviceDesc, srv)
}

func _BfGatewayService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(bftrader.BfConnectReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfGatewayServiceServer).Connect(m, &bfGatewayServiceConnectServer{stream})
}

type BfGatewayService_ConnectServer interface {
	Send(*google_protobuf.Any) error
	grpc.ServerStream
}

type bfGatewayServiceConnectServer struct {
	grpc.ServerStream
}

func (x *bfGatewayServiceConnectServer) Send(m *google_protobuf.Any) error {
	return x.ServerStream.SendMsg(m)
}

func _BfGatewayService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfPingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).Disconnect(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfGetContractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).GetContract(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfSendOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).SendOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfCancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).CancelOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_QueryAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).QueryAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BfGatewayService_QueryPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(bftrader.BfVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BfGatewayServiceServer).QueryPosition(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BfGatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bftrader.bfgateway.BfGatewayService",
	HandlerType: (*BfGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BfGatewayService_Ping_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _BfGatewayService_Disconnect_Handler,
		},
		{
			MethodName: "GetContract",
			Handler:    _BfGatewayService_GetContract_Handler,
		},
		{
			MethodName: "SendOrder",
			Handler:    _BfGatewayService_SendOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _BfGatewayService_CancelOrder_Handler,
		},
		{
			MethodName: "QueryAccount",
			Handler:    _BfGatewayService_QueryAccount_Handler,
		},
		{
			MethodName: "QueryPosition",
			Handler:    _BfGatewayService_QueryPosition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _BfGatewayService_Connect_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xdd, 0x4e, 0xc2, 0x30,
	0x14, 0xc7, 0x31, 0x1a, 0x8d, 0x07, 0x3f, 0x48, 0x43, 0x8c, 0xf4, 0x92, 0x07, 0x28, 0x06, 0xf1,
	0xca, 0x78, 0xc1, 0xc0, 0x70, 0x29, 0x4a, 0xe2, 0x7d, 0xd7, 0x9d, 0x2d, 0x4d, 0x48, 0x8b, 0x5d,
	0xa7, 0xd9, 0x5b, 0xfb, 0x08, 0x76, 0xdd, 0x5c, 0x8a, 0x42, 0xc2, 0x65, 0xff, 0x1f, 0xbf, 0x73,
	0x4e, 0x0a, 0xd7, 0x71, 0x9a, 0x71, 0x8b, 0x5f, 0xbc, 0x64, 0x1b, 0xa3, 0xad, 0x26, 0x24, 0x4e,
	0xad, 0xe1, 0x09, 0x1a, 0xd6, 0x3a, 0xf4, 0xaa, 0xd5, 0x7c, 0x86, 0x0e, 0x32, 0xad, 0xb3, 0x35,
	0x8e, 0xfc, 0x2b, 0x2e, 0xd2, 0x11, 0x57, 0x4d, 0x7d, 0xfc, 0x7d, 0x0c, 0xbd, 0x28, 0x5d, 0xd4,
	0xc5, 0x15, 0x9a, 0x4f, 0x29, 0x90, 0x3c, 0xc2, 0xd9, 0x4c, 0x2b, 0x85, 0xc2, 0x92, 0x1b, 0xd6,
	0xb2, 0xa2, 0xb4, 0x11, 0xdf, 0xf0, 0x83, 0xf6, 0x59, 0xcd, 0x64, 0xbf, 0x4c, 0x36, 0x55, 0xe5,
	0xb0, 0x73, 0x77, 0x44, 0x26, 0x70, 0xb2, 0x94, 0x2a, 0x23, 0xfd, 0xb0, 0x59, 0x29, 0x73, 0x6e,
	0x39, 0xdd, 0xa9, 0x0e, 0x3b, 0x64, 0x0c, 0x30, 0x97, 0xb9, 0x68, 0xa6, 0xf6, 0xc2, 0xd4, 0xbb,
	0x96, 0x09, 0xfd, 0xa7, 0xb8, 0xce, 0x33, 0x74, 0x17, 0x68, 0xdd, 0x52, 0xce, 0x70, 0x25, 0x1a,
	0x46, 0x02, 0xa3, 0x5a, 0xf7, 0xf6, 0xcf, 0x19, 0xde, 0x68, 0x46, 0x47, 0x70, 0xbe, 0x42, 0x95,
	0xbc, 0x18, 0xe7, 0x92, 0xad, 0x60, 0x2b, 0x57, 0x88, 0xc1, 0x1e, 0x27, 0xdf, 0x38, 0xc6, 0x13,
	0x74, 0x67, 0x5c, 0x09, 0x5c, 0xd7, 0x94, 0xad, 0x55, 0x02, 0xa3, 0xe2, 0xec, 0xba, 0x64, 0x02,
	0x17, 0xaf, 0x05, 0x9a, 0x72, 0x2a, 0x84, 0x2e, 0xd4, 0xa1, 0xf7, 0x3f, 0xc0, 0xa5, 0x6f, 0x2d,
	0x75, 0x2e, 0xad, 0xd4, 0xea, 0xb0, 0x5a, 0x7c, 0xea, 0xbf, 0xec, 0xfe, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0x80, 0xaf, 0xd9, 0x4b, 0x02, 0x00, 0x00,
}
