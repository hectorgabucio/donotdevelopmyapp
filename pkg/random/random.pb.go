// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: pkg/random/random.proto

package random

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type RandomNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *RandomNumber) Reset() {
	*x = RandomNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_random_random_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomNumber) ProtoMessage() {}

func (x *RandomNumber) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_random_random_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomNumber.ProtoReflect.Descriptor instead.
func (*RandomNumber) Descriptor() ([]byte, []int) {
	return file_pkg_random_random_proto_rawDescGZIP(), []int{0}
}

func (x *RandomNumber) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_pkg_random_random_proto protoreflect.FileDescriptor

var file_pkg_random_random_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26,
	0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x4c, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x72,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_random_random_proto_rawDescOnce sync.Once
	file_pkg_random_random_proto_rawDescData = file_pkg_random_random_proto_rawDesc
)

func file_pkg_random_random_proto_rawDescGZIP() []byte {
	file_pkg_random_random_proto_rawDescOnce.Do(func() {
		file_pkg_random_random_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_random_random_proto_rawDescData)
	})
	return file_pkg_random_random_proto_rawDescData
}

var file_pkg_random_random_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_random_random_proto_goTypes = []interface{}{
	(*RandomNumber)(nil), // 0: random.RandomNumber
	(*empty.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_pkg_random_random_proto_depIdxs = []int32{
	1, // 0: random.RandomService.GetRandom:input_type -> google.protobuf.Empty
	0, // 1: random.RandomService.GetRandom:output_type -> random.RandomNumber
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_random_random_proto_init() }
func file_pkg_random_random_proto_init() {
	if File_pkg_random_random_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_random_random_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomNumber); i {
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
			RawDescriptor: file_pkg_random_random_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_random_random_proto_goTypes,
		DependencyIndexes: file_pkg_random_random_proto_depIdxs,
		MessageInfos:      file_pkg_random_random_proto_msgTypes,
	}.Build()
	File_pkg_random_random_proto = out.File
	file_pkg_random_random_proto_rawDesc = nil
	file_pkg_random_random_proto_goTypes = nil
	file_pkg_random_random_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RandomServiceClient is the client API for RandomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomServiceClient interface {
	GetRandom(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RandomNumber, error)
}

type randomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomServiceClient(cc grpc.ClientConnInterface) RandomServiceClient {
	return &randomServiceClient{cc}
}

func (c *randomServiceClient) GetRandom(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RandomNumber, error) {
	out := new(RandomNumber)
	err := c.cc.Invoke(ctx, "/random.RandomService/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomServiceServer is the server API for RandomService service.
type RandomServiceServer interface {
	GetRandom(context.Context, *empty.Empty) (*RandomNumber, error)
}

// UnimplementedRandomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRandomServiceServer struct {
}

func (*UnimplementedRandomServiceServer) GetRandom(context.Context, *empty.Empty) (*RandomNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}

func RegisterRandomServiceServer(s *grpc.Server, srv RandomServiceServer) {
	s.RegisterService(&_RandomService_serviceDesc, srv)
}

func _RandomService_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServiceServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/random.RandomService/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServiceServer).GetRandom(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _RandomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "random.RandomService",
	HandlerType: (*RandomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRandom",
			Handler:    _RandomService_GetRandom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/random/random.proto",
}
