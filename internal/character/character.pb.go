// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: internal/character/character.proto

package character

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_character_character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_internal_character_character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_internal_character_character_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_character_character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_internal_character_character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_internal_character_character_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Output) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Output) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

var File_internal_character_character_proto protoreflect.FileDescriptor

var file_internal_character_character_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22,
	0x1f, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x42, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x32, 0x49, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_character_character_proto_rawDescOnce sync.Once
	file_internal_character_character_proto_rawDescData = file_internal_character_character_proto_rawDesc
)

func file_internal_character_character_proto_rawDescGZIP() []byte {
	file_internal_character_character_proto_rawDescOnce.Do(func() {
		file_internal_character_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_character_character_proto_rawDescData)
	})
	return file_internal_character_character_proto_rawDescData
}

var file_internal_character_character_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_character_character_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: character.Input
	(*Output)(nil), // 1: character.Output
}
var file_internal_character_character_proto_depIdxs = []int32{
	0, // 0: character.CharacterService.GetCharacter:input_type -> character.Input
	1, // 1: character.CharacterService.GetCharacter:output_type -> character.Output
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_character_character_proto_init() }
func file_internal_character_character_proto_init() {
	if File_internal_character_character_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_character_character_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_internal_character_character_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_internal_character_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_character_character_proto_goTypes,
		DependencyIndexes: file_internal_character_character_proto_depIdxs,
		MessageInfos:      file_internal_character_character_proto_msgTypes,
	}.Build()
	File_internal_character_character_proto = out.File
	file_internal_character_character_proto_rawDesc = nil
	file_internal_character_character_proto_goTypes = nil
	file_internal_character_character_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CharacterServiceClient interface {
	GetCharacter(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) GetCharacter(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/character.CharacterService/GetCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
type CharacterServiceServer interface {
	GetCharacter(context.Context, *Input) (*Output, error)
}

// UnimplementedCharacterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (*UnimplementedCharacterServiceServer) GetCharacter(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}

func RegisterCharacterServiceServer(s *grpc.Server, srv CharacterServiceServer) {
	s.RegisterService(&_CharacterService_serviceDesc, srv)
}

func _CharacterService_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/character.CharacterService/GetCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacter(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

var _CharacterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "character.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharacter",
			Handler:    _CharacterService_GetCharacter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/character/character.proto",
}
