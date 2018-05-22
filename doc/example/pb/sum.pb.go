// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package sumpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SumRequest is a request for Summator service.
type SumRequest struct {
	// A is the number we're adding to. Can't be zero for the sake of example.
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	// B is the number we're adding.
	B                    int64    `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sum_179ee1cd7782bcb6, []int{0}
}
func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (dst *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(dst, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Sum                  int64    `protobuf:"varint,1,opt,name=sum" json:"sum,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sum_179ee1cd7782bcb6, []int{1}
}
func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (dst *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(dst, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SumResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sumpb.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sumpb.SumResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SummatorClient is the client API for Summator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SummatorClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type summatorClient struct {
	cc *grpc.ClientConn
}

func NewSummatorClient(cc *grpc.ClientConn) SummatorClient {
	return &summatorClient{cc}
}

func (c *summatorClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sumpb.Summator/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummatorServer is the server API for Summator service.
type SummatorServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSummatorServer(s *grpc.Server, srv SummatorServer) {
	s.RegisterService(&_Summator_serviceDesc, srv)
}

func _Summator_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummatorServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sumpb.Summator/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummatorServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Summator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sumpb.Summator",
	HandlerType: (*SummatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Summator_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum.proto",
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_sum_179ee1cd7782bcb6) }

var fileDescriptor_sum_179ee1cd7782bcb6 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0x03, 0x41,
	0x0c, 0x86, 0xd9, 0x2e, 0x15, 0x1b, 0x3d, 0x68, 0x10, 0xac, 0x45, 0x50, 0xf6, 0xd4, 0xd3, 0x0e,
	0x2a, 0x3e, 0x86, 0x97, 0xdd, 0x93, 0xc7, 0x0c, 0x0c, 0xa5, 0xd0, 0x6c, 0xc6, 0xc9, 0x44, 0x84,
	0xd2, 0x8b, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x88, 0xec, 0x8e, 0x60, 0x6f, 0xf9, 0xc2, 0x17,
	0xfe, 0xfc, 0xb0, 0x50, 0xe3, 0x36, 0x26, 0xc9, 0x82, 0x73, 0x35, 0x8e, 0x7e, 0x75, 0xbb, 0x11,
	0xd9, 0xec, 0x82, 0xa3, 0xb8, 0x75, 0x34, 0x0c, 0x92, 0x29, 0x6f, 0x65, 0xd0, 0x22, 0x35, 0x6b,
	0x80, 0xde, 0xb8, 0x0b, 0x6f, 0x16, 0x34, 0xe3, 0x39, 0x54, 0xb4, 0xac, 0xee, 0xab, 0x75, 0xdd,
	0x55, 0x34, 0x92, 0x5f, 0xce, 0x0a, 0xf9, 0xe6, 0x19, 0xce, 0x26, 0x53, 0xa3, 0x0c, 0x1a, 0xf0,
	0x02, 0x6a, 0x35, 0xfe, 0x93, 0xc7, 0x11, 0xaf, 0x60, 0x1e, 0x52, 0x92, 0x34, 0x9d, 0x2c, 0xba,
	0x02, 0x8f, 0xaf, 0x70, 0xda, 0x1b, 0x33, 0x65, 0x49, 0xf8, 0x02, 0x75, 0x6f, 0x8c, 0x97, 0xed,
	0xf4, 0x59, 0xfb, 0x1f, 0xbc, 0xc2, 0xe3, 0x55, 0x49, 0x68, 0xee, 0x3e, 0xbf, 0x7f, 0xbe, 0x66,
	0x37, 0x78, 0xed, 0xde, 0x1f, 0x5c, 0xf8, 0x20, 0x8e, 0xbb, 0xe0, 0xd4, 0xd8, 0xed, 0xe9, 0xe0,
	0xf6, 0xfe, 0xe0, 0x4f, 0xa6, 0x0a, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x35, 0x06,
	0x1f, 0xf4, 0x00, 0x00, 0x00,
}
