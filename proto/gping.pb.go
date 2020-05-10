// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: gping.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DstAddr  string `protobuf:"bytes,1,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
	Count    int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Ttl      int32  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Size     int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Interval string `protobuf:"bytes,5,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_gping_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetDstAddr() string {
	if x != nil {
		return x.DstAddr
	}
	return ""
}

func (x *PingRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PingRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *PingRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PingRequest) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rtt  float64 `protobuf:"fixed64,1,opt,name=rtt,proto3" json:"rtt,omitempty"`
	Seq  int32   `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Ttl  int32   `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Size int32   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Addr string  `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	Err  string  `protobuf:"bytes,6,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_gping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_gping_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetRtt() float64 {
	if x != nil {
		return x.Rtt
	}
	return 0
}

func (x *PingReply) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *PingReply) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *PingReply) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PingReply) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *PingReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_gping_proto protoreflect.FileDescriptor

var file_gping_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x22, 0x7b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x74, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x72, 0x74, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x2f, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gping_proto_rawDescOnce sync.Once
	file_gping_proto_rawDescData = file_gping_proto_rawDesc
)

func file_gping_proto_rawDescGZIP() []byte {
	file_gping_proto_rawDescOnce.Do(func() {
		file_gping_proto_rawDescData = protoimpl.X.CompressGZIP(file_gping_proto_rawDescData)
	})
	return file_gping_proto_rawDescData
}

var file_gping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gping_proto_goTypes = []interface{}{
	(*PingRequest)(nil), // 0: PingRequest
	(*PingReply)(nil),   // 1: PingReply
}
var file_gping_proto_depIdxs = []int32{
	0, // 0: Ping.GetPing:input_type -> PingRequest
	1, // 1: Ping.GetPing:output_type -> PingReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gping_proto_init() }
func file_gping_proto_init() {
	if File_gping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gping_proto_goTypes,
		DependencyIndexes: file_gping_proto_depIdxs,
		MessageInfos:      file_gping_proto_msgTypes,
	}.Build()
	File_gping_proto = out.File
	file_gping_proto_rawDesc = nil
	file_gping_proto_goTypes = nil
	file_gping_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingClient interface {
	GetPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (Ping_GetPingClient, error)
}

type pingClient struct {
	cc grpc.ClientConnInterface
}

func NewPingClient(cc grpc.ClientConnInterface) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) GetPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (Ping_GetPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ping_serviceDesc.Streams[0], "/Ping/GetPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingGetPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ping_GetPingClient interface {
	Recv() (*PingReply, error)
	grpc.ClientStream
}

type pingGetPingClient struct {
	grpc.ClientStream
}

func (x *pingGetPingClient) Recv() (*PingReply, error) {
	m := new(PingReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingServer is the server API for Ping service.
type PingServer interface {
	GetPing(*PingRequest, Ping_GetPingServer) error
}

// UnimplementedPingServer can be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (*UnimplementedPingServer) GetPing(*PingRequest, Ping_GetPingServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPing not implemented")
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_GetPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PingServer).GetPing(m, &pingGetPingServer{stream})
}

type Ping_GetPingServer interface {
	Send(*PingReply) error
	grpc.ServerStream
}

type pingGetPingServer struct {
	grpc.ServerStream
}

func (x *pingGetPingServer) Send(m *PingReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Ping",
	HandlerType: (*PingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPing",
			Handler:       _Ping_GetPing_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gping.proto",
}
