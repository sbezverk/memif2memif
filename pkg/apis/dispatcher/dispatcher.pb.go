// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dispatcher.proto

package dispatcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Operation int32

const (
	Operation_DEFAULT Operation = 0
	Operation_INTRO   Operation = 1
	Operation_CONNECT Operation = 2
	Operation_LISTEN  Operation = 3
	Operation_SOCKET  Operation = 4
)

var Operation_name = map[int32]string{
	0: "DEFAULT",
	1: "INTRO",
	2: "CONNECT",
	3: "LISTEN",
	4: "SOCKET",
}
var Operation_value = map[string]int32{
	"DEFAULT": 0,
	"INTRO":   1,
	"CONNECT": 2,
	"LISTEN":  3,
	"SOCKET":  4,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{0}
}

type ID struct {
	PodName              string   `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	PodNamespace         string   `protobuf:"bytes,2,opt,name=pod_namespace,json=podNamespace,proto3" json:"pod_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *ID) GetPodNamespace() string {
	if m != nil {
		return m.PodNamespace
	}
	return ""
}

// Intro message is sent from a client to the dispatcher to announce
// its presence on the node.
type Intro struct {
	PodId                *ID      `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Intro) Reset()         { *m = Intro{} }
func (m *Intro) String() string { return proto.CompactTextString(m) }
func (*Intro) ProtoMessage()    {}
func (*Intro) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{1}
}
func (m *Intro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Intro.Unmarshal(m, b)
}
func (m *Intro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Intro.Marshal(b, m, deterministic)
}
func (dst *Intro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intro.Merge(dst, src)
}
func (m *Intro) XXX_Size() int {
	return xxx_messageInfo_Intro.Size(m)
}
func (m *Intro) XXX_DiscardUnknown() {
	xxx_messageInfo_Intro.DiscardUnknown(m)
}

var xxx_messageInfo_Intro proto.InternalMessageInfo

func (m *Intro) GetPodId() *ID {
	if m != nil {
		return m.PodId
	}
	return nil
}

// Connect message is sent by the client which wants to build memif connection
// to another pod which is in Listen mode.
type Connect struct {
	SrcId                *ID      `protobuf:"bytes,1,opt,name=src_id,json=srcId,proto3" json:"src_id,omitempty"`
	DstId                *ID      `protobuf:"bytes,2,opt,name=dst_id,json=dstId,proto3" json:"dst_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{2}
}
func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (dst *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(dst, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetSrcId() *ID {
	if m != nil {
		return m.SrcId
	}
	return nil
}

func (m *Connect) GetDstId() *ID {
	if m != nil {
		return m.DstId
	}
	return nil
}

// Listen message is sent by the the client which wants to listen for incoming
// memif connection. Usually a server type application.
type Listen struct {
	ListenerId           *ID      `protobuf:"bytes,1,opt,name=listener_id,json=listenerId,proto3" json:"listener_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listen) Reset()         { *m = Listen{} }
func (m *Listen) String() string { return proto.CompactTextString(m) }
func (*Listen) ProtoMessage()    {}
func (*Listen) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{3}
}
func (m *Listen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listen.Unmarshal(m, b)
}
func (m *Listen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listen.Marshal(b, m, deterministic)
}
func (dst *Listen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listen.Merge(dst, src)
}
func (m *Listen) XXX_Size() int {
	return xxx_messageInfo_Listen.Size(m)
}
func (m *Listen) XXX_DiscardUnknown() {
	xxx_messageInfo_Listen.DiscardUnknown(m)
}

var xxx_messageInfo_Listen proto.InternalMessageInfo

func (m *Listen) GetListenerId() *ID {
	if m != nil {
		return m.ListenerId
	}
	return nil
}

// Socket message is sent by the dispatcher to both client and server with the sockets
// they have to use to build memif connection
type Socket struct {
	PodId                *ID      `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Fd                   int64    `protobuf:"varint,2,opt,name=fd,proto3" json:"fd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Socket) Reset()         { *m = Socket{} }
func (m *Socket) String() string { return proto.CompactTextString(m) }
func (*Socket) ProtoMessage()    {}
func (*Socket) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{4}
}
func (m *Socket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Socket.Unmarshal(m, b)
}
func (m *Socket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Socket.Marshal(b, m, deterministic)
}
func (dst *Socket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Socket.Merge(dst, src)
}
func (m *Socket) XXX_Size() int {
	return xxx_messageInfo_Socket.Size(m)
}
func (m *Socket) XXX_DiscardUnknown() {
	xxx_messageInfo_Socket.DiscardUnknown(m)
}

var xxx_messageInfo_Socket proto.InternalMessageInfo

func (m *Socket) GetPodId() *ID {
	if m != nil {
		return m.PodId
	}
	return nil
}

func (m *Socket) GetFd() int64 {
	if m != nil {
		return m.Fd
	}
	return 0
}

// Reply message is sent by the dispatcher as a confirmation echoing requesting pod
// id and operation, extended error may contain additional information is success is false
type Reply struct {
	PodId                *ID       `protobuf:"bytes,1,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Op                   Operation `protobuf:"varint,2,opt,name=op,proto3,enum=dispatcher.Operation" json:"op,omitempty"`
	Success              bool      `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	ExtendedError        string    `protobuf:"bytes,4,opt,name=extended_error,json=extendedError,proto3" json:"extended_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dispatcher_336dc1b27576c40d, []int{5}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetPodId() *ID {
	if m != nil {
		return m.PodId
	}
	return nil
}

func (m *Reply) GetOp() Operation {
	if m != nil {
		return m.Op
	}
	return Operation_DEFAULT
}

func (m *Reply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Reply) GetExtendedError() string {
	if m != nil {
		return m.ExtendedError
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "dispatcher.ID")
	proto.RegisterType((*Intro)(nil), "dispatcher.Intro")
	proto.RegisterType((*Connect)(nil), "dispatcher.Connect")
	proto.RegisterType((*Listen)(nil), "dispatcher.Listen")
	proto.RegisterType((*Socket)(nil), "dispatcher.Socket")
	proto.RegisterType((*Reply)(nil), "dispatcher.Reply")
	proto.RegisterEnum("dispatcher.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	SendIntro(ctx context.Context, in *Intro, opts ...grpc.CallOption) (*Reply, error)
	SendConnect(ctx context.Context, in *Connect, opts ...grpc.CallOption) (*Reply, error)
	SendListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Reply, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) SendIntro(ctx context.Context, in *Intro, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dispatcher.Client/SendIntro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) SendConnect(ctx context.Context, in *Connect, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dispatcher.Client/SendConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) SendListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dispatcher.Client/SendListen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	SendIntro(context.Context, *Intro) (*Reply, error)
	SendConnect(context.Context, *Connect) (*Reply, error)
	SendListen(context.Context, *Listen) (*Reply, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_SendIntro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Intro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).SendIntro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.Client/SendIntro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).SendIntro(ctx, req.(*Intro))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_SendConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).SendConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.Client/SendConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).SendConnect(ctx, req.(*Connect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_SendListen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Listen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).SendListen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.Client/SendListen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).SendListen(ctx, req.(*Listen))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcher.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendIntro",
			Handler:    _Client_SendIntro_Handler,
		},
		{
			MethodName: "SendConnect",
			Handler:    _Client_SendConnect_Handler,
		},
		{
			MethodName: "SendListen",
			Handler:    _Client_SendListen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher.proto",
}

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	SendSocket(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Reply, error)
}

type dispatcherClient struct {
	cc *grpc.ClientConn
}

func NewDispatcherClient(cc *grpc.ClientConn) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) SendSocket(ctx context.Context, in *Socket, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/dispatcher.Dispatcher/SendSocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	SendSocket(context.Context, *Socket) (*Reply, error)
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_SendSocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Socket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).SendSocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.Dispatcher/SendSocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).SendSocket(ctx, req.(*Socket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcher.Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSocket",
			Handler:    _Dispatcher_SendSocket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher.proto",
}

func init() { proto.RegisterFile("dispatcher.proto", fileDescriptor_dispatcher_336dc1b27576c40d) }

var fileDescriptor_dispatcher_336dc1b27576c40d = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x14, 0x85, 0xcd, 0x74, 0x9b, 0x6e, 0x6f, 0xdd, 0x12, 0x47, 0x84, 0xe8, 0xd3, 0x12, 0x29, 0x2c,
	0x3e, 0x54, 0xec, 0xe2, 0x83, 0x4f, 0xb2, 0x24, 0x11, 0x07, 0x4b, 0x0a, 0x69, 0xc4, 0xc7, 0x25,
	0xce, 0xdc, 0xc5, 0x60, 0x77, 0x66, 0x98, 0x19, 0x41, 0xff, 0x88, 0x7f, 0xc1, 0xbf, 0x29, 0x93,
	0x34, 0x6b, 0x0b, 0x71, 0xd9, 0xb7, 0xc9, 0x77, 0xcf, 0x39, 0xc9, 0xcd, 0x61, 0x20, 0x12, 0x8d,
	0xd5, 0xb5, 0xe3, 0xdf, 0xd0, 0x2c, 0xb5, 0x51, 0x4e, 0x51, 0xf8, 0x47, 0x92, 0x0c, 0x08, 0xcb,
	0xe8, 0x73, 0x38, 0xd5, 0x4a, 0x5c, 0xcb, 0xfa, 0x16, 0xe3, 0xe0, 0x3c, 0xb8, 0x98, 0x96, 0x13,
	0xad, 0x44, 0x51, 0xdf, 0x22, 0x7d, 0x09, 0x67, 0xfd, 0xc8, 0xea, 0x9a, 0x63, 0x4c, 0xda, 0xf9,
	0xe3, 0xfd, 0xbc, 0x65, 0xc9, 0x12, 0xc6, 0x4c, 0x3a, 0xa3, 0xe8, 0x02, 0x42, 0xaf, 0x6e, 0x44,
	0x1b, 0x33, 0x5b, 0xcd, 0x97, 0x07, 0x6f, 0x67, 0x59, 0x39, 0xd6, 0x4a, 0x30, 0x91, 0x7c, 0x81,
	0x49, 0xaa, 0xa4, 0x44, 0xee, 0xbc, 0xc3, 0x1a, 0x7e, 0x8f, 0xc3, 0x1a, 0xce, 0x84, 0x97, 0x09,
	0xeb, 0xbc, 0x8c, 0x0c, 0xcb, 0x84, 0x75, 0x4c, 0x24, 0xef, 0x20, 0x5c, 0x37, 0xd6, 0xa1, 0xa4,
	0xaf, 0x61, 0xb6, 0x6b, 0x4f, 0x68, 0xfe, 0x1f, 0x0e, 0xbd, 0x84, 0x89, 0xe4, 0x3d, 0x84, 0x5b,
	0xc5, 0xbf, 0xa3, 0x7b, 0xe0, 0x12, 0x74, 0x0e, 0xe4, 0xa6, 0xfb, 0x9c, 0x51, 0x49, 0x6e, 0x44,
	0xf2, 0x3b, 0x80, 0x71, 0x89, 0x7a, 0xf7, 0xeb, 0xa1, 0x01, 0x0b, 0x20, 0x4a, 0xb7, 0x01, 0xf3,
	0xd5, 0xb3, 0x43, 0xc9, 0x46, 0xa3, 0xa9, 0x5d, 0xa3, 0x64, 0x49, 0x94, 0xa6, 0x31, 0x4c, 0xec,
	0x0f, 0xce, 0xd1, 0xda, 0x78, 0x74, 0x1e, 0x5c, 0x9c, 0x96, 0xfd, 0x23, 0x5d, 0xc0, 0x1c, 0x7f,
	0x3a, 0x94, 0x02, 0xc5, 0x35, 0x1a, 0xa3, 0x4c, 0x7c, 0xd2, 0x96, 0x73, 0xd6, 0xd3, 0xdc, 0xc3,
	0x57, 0x1f, 0x61, 0x7a, 0x97, 0x48, 0x67, 0x30, 0xc9, 0xf2, 0x0f, 0x57, 0x9f, 0xd7, 0x55, 0xf4,
	0x88, 0x4e, 0x61, 0xcc, 0x8a, 0xaa, 0xdc, 0x44, 0x81, 0xe7, 0xe9, 0xa6, 0x28, 0xf2, 0xb4, 0x8a,
	0x08, 0x05, 0x08, 0xd7, 0x6c, 0x5b, 0xe5, 0x45, 0x34, 0xf2, 0xe7, 0xed, 0x26, 0xfd, 0x94, 0x57,
	0xd1, 0xc9, 0xea, 0x4f, 0x00, 0x61, 0xba, 0x6b, 0x50, 0x3a, 0xfa, 0x06, 0xa6, 0x5b, 0x94, 0xa2,
	0xab, 0xfd, 0xc9, 0xd1, 0x82, 0x1e, 0xbd, 0x38, 0x42, 0xdd, 0x6f, 0x79, 0x0b, 0x33, 0x6f, 0xe9,
	0x9b, 0x7f, 0x7a, 0xa8, 0xd8, 0xc3, 0x21, 0xdb, 0x25, 0x80, 0xb7, 0xed, 0x7b, 0xa5, 0x87, 0x82,
	0x8e, 0x0d, 0x98, 0x56, 0x57, 0x00, 0xd9, 0x1d, 0xeb, 0x23, 0xf6, 0xfd, 0x1e, 0x45, 0x74, 0x6c,
	0x20, 0xe2, 0x6b, 0xd8, 0xde, 0x96, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x89, 0x6f,
	0xa3, 0x41, 0x03, 0x00, 0x00,
}