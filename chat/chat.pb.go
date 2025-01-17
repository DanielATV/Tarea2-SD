// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: chat.proto

package chat

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log    string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	Partes int64  `protobuf:"varint,2,opt,name=partes,proto3" json:"partes,omitempty"`
	Nombre string `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *LogInfo) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *LogInfo) GetPartes() int64 {
	if x != nil {
		return x.Partes
	}
	return 0
}

func (x *LogInfo) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk  []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Indice int64  `protobuf:"varint,2,opt,name=indice,proto3" json:"indice,omitempty"`
	Nombre string `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Total  int64  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *Chunk) GetIndice() int64 {
	if x != nil {
		return x.Indice
	}
	return 0
}

func (x *Chunk) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Chunk) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x2d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4b, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x22, 0x63,
	0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x32, 0xac, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0b, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2b, 0x0a, 0x09,
	0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44, 0x69, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: chat.Message
	(*LogInfo)(nil), // 1: chat.LogInfo
	(*Chunk)(nil),   // 2: chat.Chunk
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.SayHello:input_type -> chat.Message
	2, // 1: chat.ChatService.SendChunk:input_type -> chat.Chunk
	0, // 2: chat.ChatService.LibrosDis:input_type -> chat.Message
	0, // 3: chat.ChatService.SendPropuesta:input_type -> chat.Message
	0, // 4: chat.ChatService.RequestLog:input_type -> chat.Message
	0, // 5: chat.ChatService.RequestChunk:input_type -> chat.Message
	0, // 6: chat.ChatService.CheckStatus:input_type -> chat.Message
	1, // 7: chat.ChatService.WriteLog:input_type -> chat.LogInfo
	2, // 8: chat.ChatService.DistributeChunk:input_type -> chat.Chunk
	0, // 9: chat.ChatService.SayHello:output_type -> chat.Message
	0, // 10: chat.ChatService.SendChunk:output_type -> chat.Message
	0, // 11: chat.ChatService.LibrosDis:output_type -> chat.Message
	0, // 12: chat.ChatService.SendPropuesta:output_type -> chat.Message
	0, // 13: chat.ChatService.RequestLog:output_type -> chat.Message
	2, // 14: chat.ChatService.RequestChunk:output_type -> chat.Chunk
	0, // 15: chat.ChatService.CheckStatus:output_type -> chat.Message
	0, // 16: chat.ChatService.WriteLog:output_type -> chat.Message
	0, // 17: chat.ChatService.DistributeChunk:output_type -> chat.Message
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	SendChunk(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendChunkClient, error)
	LibrosDis(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	SendPropuesta(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	RequestLog(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	RequestChunk(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Chunk, error)
	CheckStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	WriteLog(ctx context.Context, in *LogInfo, opts ...grpc.CallOption) (*Message, error)
	DistributeChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendChunk(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendChunkClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/SendChunk", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendChunkClient{stream}
	return x, nil
}

type ChatService_SendChunkClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type chatServiceSendChunkClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendChunkClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendChunkClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) LibrosDis(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/LibrosDis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendPropuesta(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SendPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RequestLog(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RequestLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RequestChunk(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RequestChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CheckStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CheckStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) WriteLog(ctx context.Context, in *LogInfo, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/WriteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DistributeChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/DistributeChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	SendChunk(ChatService_SendChunkServer) error
	LibrosDis(context.Context, *Message) (*Message, error)
	SendPropuesta(context.Context, *Message) (*Message, error)
	RequestLog(context.Context, *Message) (*Message, error)
	RequestChunk(context.Context, *Message) (*Chunk, error)
	CheckStatus(context.Context, *Message) (*Message, error)
	WriteLog(context.Context, *LogInfo) (*Message, error)
	DistributeChunk(context.Context, *Chunk) (*Message, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedChatServiceServer) SendChunk(ChatService_SendChunkServer) error {
	return status.Errorf(codes.Unimplemented, "method SendChunk not implemented")
}
func (*UnimplementedChatServiceServer) LibrosDis(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LibrosDis not implemented")
}
func (*UnimplementedChatServiceServer) SendPropuesta(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPropuesta not implemented")
}
func (*UnimplementedChatServiceServer) RequestLog(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLog not implemented")
}
func (*UnimplementedChatServiceServer) RequestChunk(context.Context, *Message) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestChunk not implemented")
}
func (*UnimplementedChatServiceServer) CheckStatus(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (*UnimplementedChatServiceServer) WriteLog(context.Context, *LogInfo) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (*UnimplementedChatServiceServer) DistributeChunk(context.Context, *Chunk) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeChunk not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendChunk(&chatServiceSendChunkServer{stream})
}

type ChatService_SendChunkServer interface {
	SendAndClose(*Message) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type chatServiceSendChunkServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendChunkServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendChunkServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_LibrosDis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LibrosDis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/LibrosDis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LibrosDis(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SendPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendPropuesta(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RequestLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RequestLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RequestLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RequestLog(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RequestChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RequestChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RequestChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RequestChunk(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CheckStatus(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_WriteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).WriteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/WriteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).WriteLog(ctx, req.(*LogInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DistributeChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DistributeChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/DistributeChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DistributeChunk(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
		{
			MethodName: "LibrosDis",
			Handler:    _ChatService_LibrosDis_Handler,
		},
		{
			MethodName: "SendPropuesta",
			Handler:    _ChatService_SendPropuesta_Handler,
		},
		{
			MethodName: "RequestLog",
			Handler:    _ChatService_RequestLog_Handler,
		},
		{
			MethodName: "RequestChunk",
			Handler:    _ChatService_RequestChunk_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _ChatService_CheckStatus_Handler,
		},
		{
			MethodName: "WriteLog",
			Handler:    _ChatService_WriteLog_Handler,
		},
		{
			MethodName: "DistributeChunk",
			Handler:    _ChatService_DistributeChunk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendChunk",
			Handler:       _ChatService_SendChunk_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
