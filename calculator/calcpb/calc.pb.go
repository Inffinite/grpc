// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Calculation struct {
	FirstInt             int64    `protobuf:"varint,1,opt,name=first_int,json=firstInt,proto3" json:"first_int,omitempty"`
	SecondInt            int64    `protobuf:"varint,2,opt,name=second_int,json=secondInt,proto3" json:"second_int,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calculation) Reset()         { *m = Calculation{} }
func (m *Calculation) String() string { return proto.CompactTextString(m) }
func (*Calculation) ProtoMessage()    {}
func (*Calculation) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{0}
}

func (m *Calculation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calculation.Unmarshal(m, b)
}
func (m *Calculation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calculation.Marshal(b, m, deterministic)
}
func (m *Calculation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calculation.Merge(m, src)
}
func (m *Calculation) XXX_Size() int {
	return xxx_messageInfo_Calculation.Size(m)
}
func (m *Calculation) XXX_DiscardUnknown() {
	xxx_messageInfo_Calculation.DiscardUnknown(m)
}

var xxx_messageInfo_Calculation proto.InternalMessageInfo

func (m *Calculation) GetFirstInt() int64 {
	if m != nil {
		return m.FirstInt
	}
	return 0
}

func (m *Calculation) GetSecondInt() int64 {
	if m != nil {
		return m.SecondInt
	}
	return 0
}

type CalcRequest struct {
	Calculation          *Calculation `protobuf:"bytes,1,opt,name=calculation,proto3" json:"calculation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{1}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetCalculation() *Calculation {
	if m != nil {
		return m.Calculation
	}
	return nil
}

type CalcResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{2}
}

func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (m *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(m, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecomposeRequest struct {
	PrimeNumber          int64    `protobuf:"varint,1,opt,name=primeNumber,proto3" json:"primeNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecomposeRequest) Reset()         { *m = PrimeNumberDecomposeRequest{} }
func (m *PrimeNumberDecomposeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposeRequest) ProtoMessage()    {}
func (*PrimeNumberDecomposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{3}
}

func (m *PrimeNumberDecomposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecomposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecomposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecomposeRequest.Merge(m, src)
}
func (m *PrimeNumberDecomposeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Size(m)
}
func (m *PrimeNumberDecomposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecomposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecomposeRequest proto.InternalMessageInfo

func (m *PrimeNumberDecomposeRequest) GetPrimeNumber() int64 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

type PrimeNumberDecomposeResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecomposeResponse) Reset()         { *m = PrimeNumberDecomposeResponse{} }
func (m *PrimeNumberDecomposeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposeResponse) ProtoMessage()    {}
func (*PrimeNumberDecomposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{4}
}

func (m *PrimeNumberDecomposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecomposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecomposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecomposeResponse.Merge(m, src)
}
func (m *PrimeNumberDecomposeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Size(m)
}
func (m *PrimeNumberDecomposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecomposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecomposeResponse proto.InternalMessageInfo

func (m *PrimeNumberDecomposeResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Calculation)(nil), "calc.Calculation")
	proto.RegisterType((*CalcRequest)(nil), "calc.CalcRequest")
	proto.RegisterType((*CalcResponse)(nil), "calc.CalcResponse")
	proto.RegisterType((*PrimeNumberDecomposeRequest)(nil), "calc.PrimeNumberDecomposeRequest")
	proto.RegisterType((*PrimeNumberDecomposeResponse)(nil), "calc.PrimeNumberDecomposeResponse")
}

func init() { proto.RegisterFile("calculator/calcpb/calc.proto", fileDescriptor_cfc74e58bc0fa04b) }

var fileDescriptor_cfc74e58bc0fa04b = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x5a, 0x42, 0x3b, 0xf1, 0xe2, 0x20, 0x22, 0x6d, 0x85, 0xba, 0x07, 0xf1, 0xd4, 0x4a,
	0x0b, 0x5e, 0x85, 0xea, 0x25, 0x17, 0x91, 0x78, 0xf3, 0x22, 0xc9, 0x3a, 0xc2, 0x42, 0xb2, 0xbb,
	0xee, 0x6e, 0xfc, 0x1f, 0xff, 0x54, 0xb2, 0x9b, 0xd0, 0x3d, 0x94, 0x9c, 0x76, 0xe6, 0xcd, 0xdb,
	0xf7, 0xe6, 0x31, 0xb0, 0xe4, 0x65, 0xcd, 0xdb, 0xba, 0x74, 0xca, 0x6c, 0xba, 0x52, 0x57, 0xfe,
	0x59, 0x6b, 0xa3, 0x9c, 0xc2, 0x49, 0x57, 0xb3, 0x1c, 0xb2, 0xe7, 0x9e, 0x25, 0x94, 0xc4, 0x05,
	0xcc, 0xbe, 0x85, 0xb1, 0xee, 0x53, 0x48, 0x77, 0x9d, 0xac, 0x92, 0xfb, 0xb3, 0x62, 0xea, 0x81,
	0x5c, 0x3a, 0xbc, 0x01, 0xb0, 0xc4, 0x95, 0xfc, 0xf2, 0xd3, 0x53, 0x3f, 0x9d, 0x05, 0x24, 0x97,
	0x8e, 0xed, 0x83, 0x54, 0x41, 0x3f, 0x2d, 0x59, 0x87, 0x3b, 0xc8, 0xf8, 0x41, 0xd9, 0x8b, 0x65,
	0xdb, 0x8b, 0xb5, 0xdf, 0x20, 0xb2, 0x2c, 0x62, 0x16, 0xbb, 0x83, 0xf3, 0xa0, 0x61, 0xb5, 0x92,
	0x96, 0xf0, 0x0a, 0x52, 0x43, 0xb6, 0xad, 0x87, 0x65, 0xfa, 0x8e, 0x3d, 0xc1, 0xe2, 0xcd, 0x88,
	0x86, 0x5e, 0xdb, 0xa6, 0x22, 0xf3, 0x42, 0x5c, 0x35, 0x5a, 0x59, 0x1a, 0xbc, 0x57, 0x90, 0xe9,
	0xc3, 0xb8, 0xff, 0x1b, 0x43, 0xec, 0x11, 0x96, 0xc7, 0x05, 0xc6, 0x8d, 0xb7, 0x7f, 0x49, 0x48,
	0xf9, 0x4e, 0xe6, 0x57, 0x70, 0xc2, 0x0d, 0x4c, 0xba, 0x16, 0xa3, 0x60, 0xfd, 0x12, 0x73, 0x8c,
	0xa1, 0x20, 0xcb, 0x4e, 0xb0, 0x84, 0xcb, 0x63, 0xc6, 0x78, 0x1b, 0xd8, 0x23, 0xa9, 0xe6, 0x6c,
	0x8c, 0x32, 0x18, 0x3c, 0x24, 0xfb, 0xe9, 0x47, 0x1a, 0xce, 0x5d, 0xa5, 0xfe, 0xd4, 0xbb, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x4e, 0x6c, 0xe4, 0x0a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// Unary
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	// Server streaming
	PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposeRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecomposeClient, error)
}

type calcServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalcServiceClient(cc *grpc.ClientConn) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposeRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecomposeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calc.CalcService/PrimeNumberDecompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeNumberDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeNumberDecomposeClient interface {
	Recv() (*PrimeNumberDecomposeResponse, error)
	grpc.ClientStream
}

type calcServicePrimeNumberDecomposeClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeNumberDecomposeClient) Recv() (*PrimeNumberDecomposeResponse, error) {
	m := new(PrimeNumberDecomposeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// Unary
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
	// Server streaming
	PrimeNumberDecompose(*PrimeNumberDecomposeRequest, CalcService_PrimeNumberDecomposeServer) error
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Calc(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}
func (*UnimplementedCalcServiceServer) PrimeNumberDecompose(req *PrimeNumberDecomposeRequest, srv CalcService_PrimeNumberDecomposeServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumberDecompose not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeNumberDecompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecomposeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeNumberDecompose(m, &calcServicePrimeNumberDecomposeServer{stream})
}

type CalcService_PrimeNumberDecomposeServer interface {
	Send(*PrimeNumberDecomposeResponse) error
	grpc.ServerStream
}

type calcServicePrimeNumberDecomposeServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeNumberDecomposeServer) Send(m *PrimeNumberDecomposeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecompose",
			Handler:       _CalcService_PrimeNumberDecompose_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}