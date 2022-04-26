// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: webdav.proto

package webdavrpc

import (
	context "context"
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

type MountWebdavRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url        string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Dir        string `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
	User       string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Password   string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ConfigName string `protobuf:"bytes,5,opt,name=configName,proto3" json:"configName,omitempty"`
	Uid        string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid        string `protobuf:"bytes,7,opt,name=gid,proto3" json:"gid,omitempty"`
	Target     string `protobuf:"bytes,8,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *MountWebdavRequest) Reset() {
	*x = MountWebdavRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webdav_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountWebdavRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountWebdavRequest) ProtoMessage() {}

func (x *MountWebdavRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webdav_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountWebdavRequest.ProtoReflect.Descriptor instead.
func (*MountWebdavRequest) Descriptor() ([]byte, []int) {
	return file_webdav_proto_rawDescGZIP(), []int{0}
}

func (x *MountWebdavRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MountWebdavRequest) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *MountWebdavRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MountWebdavRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MountWebdavRequest) GetConfigName() string {
	if x != nil {
		return x.ConfigName
	}
	return ""
}

func (x *MountWebdavRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *MountWebdavRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *MountWebdavRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type MountWebdavResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *MountWebdavResponse) Reset() {
	*x = MountWebdavResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webdav_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountWebdavResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountWebdavResponse) ProtoMessage() {}

func (x *MountWebdavResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webdav_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountWebdavResponse.ProtoReflect.Descriptor instead.
func (*MountWebdavResponse) Descriptor() ([]byte, []int) {
	return file_webdav_proto_rawDescGZIP(), []int{1}
}

func (x *MountWebdavResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type UmountWebdavRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MountTarget string `protobuf:"bytes,1,opt,name=mountTarget,proto3" json:"mountTarget,omitempty"`
	ConfigName  string `protobuf:"bytes,2,opt,name=configName,proto3" json:"configName,omitempty"`
}

func (x *UmountWebdavRequest) Reset() {
	*x = UmountWebdavRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webdav_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountWebdavRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountWebdavRequest) ProtoMessage() {}

func (x *UmountWebdavRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webdav_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmountWebdavRequest.ProtoReflect.Descriptor instead.
func (*UmountWebdavRequest) Descriptor() ([]byte, []int) {
	return file_webdav_proto_rawDescGZIP(), []int{2}
}

func (x *UmountWebdavRequest) GetMountTarget() string {
	if x != nil {
		return x.MountTarget
	}
	return ""
}

func (x *UmountWebdavRequest) GetConfigName() string {
	if x != nil {
		return x.ConfigName
	}
	return ""
}

type UmountWebdavResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *UmountWebdavResponse) Reset() {
	*x = UmountWebdavResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webdav_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountWebdavResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountWebdavResponse) ProtoMessage() {}

func (x *UmountWebdavResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webdav_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmountWebdavResponse.ProtoReflect.Descriptor instead.
func (*UmountWebdavResponse) Descriptor() ([]byte, []int) {
	return file_webdav_proto_rawDescGZIP(), []int{3}
}

func (x *UmountWebdavResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_webdav_proto protoreflect.FileDescriptor

var file_webdav_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x2d, 0x0a,
	0x13, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x57, 0x0a, 0x13,
	0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x57,
	0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xa5, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x57,
	0x65, 0x62, 0x64, 0x61, 0x76, 0x12, 0x1a, 0x2e, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76,
	0x12, 0x1b, 0x2e, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62,
	0x64, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x3b, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webdav_proto_rawDescOnce sync.Once
	file_webdav_proto_rawDescData = file_webdav_proto_rawDesc
)

func file_webdav_proto_rawDescGZIP() []byte {
	file_webdav_proto_rawDescOnce.Do(func() {
		file_webdav_proto_rawDescData = protoimpl.X.CompressGZIP(file_webdav_proto_rawDescData)
	})
	return file_webdav_proto_rawDescData
}

var file_webdav_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_webdav_proto_goTypes = []interface{}{
	(*MountWebdavRequest)(nil),   // 0: webdav.MountWebdavRequest
	(*MountWebdavResponse)(nil),  // 1: webdav.MountWebdavResponse
	(*UmountWebdavRequest)(nil),  // 2: webdav.UmountWebdavRequest
	(*UmountWebdavResponse)(nil), // 3: webdav.UmountWebdavResponse
}
var file_webdav_proto_depIdxs = []int32{
	0, // 0: webdav.MountService.MountWebdav:input_type -> webdav.MountWebdavRequest
	2, // 1: webdav.MountService.UmountWebdav:input_type -> webdav.UmountWebdavRequest
	1, // 2: webdav.MountService.MountWebdav:output_type -> webdav.MountWebdavResponse
	3, // 3: webdav.MountService.UmountWebdav:output_type -> webdav.UmountWebdavResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_webdav_proto_init() }
func file_webdav_proto_init() {
	if File_webdav_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webdav_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountWebdavRequest); i {
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
		file_webdav_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountWebdavResponse); i {
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
		file_webdav_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UmountWebdavRequest); i {
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
		file_webdav_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UmountWebdavResponse); i {
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
			RawDescriptor: file_webdav_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webdav_proto_goTypes,
		DependencyIndexes: file_webdav_proto_depIdxs,
		MessageInfos:      file_webdav_proto_msgTypes,
	}.Build()
	File_webdav_proto = out.File
	file_webdav_proto_rawDesc = nil
	file_webdav_proto_goTypes = nil
	file_webdav_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MountServiceClient is the client API for MountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MountServiceClient interface {
	MountWebdav(ctx context.Context, in *MountWebdavRequest, opts ...grpc.CallOption) (*MountWebdavResponse, error)
	UmountWebdav(ctx context.Context, in *UmountWebdavRequest, opts ...grpc.CallOption) (*UmountWebdavResponse, error)
}

type mountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMountServiceClient(cc grpc.ClientConnInterface) MountServiceClient {
	return &mountServiceClient{cc}
}

func (c *mountServiceClient) MountWebdav(ctx context.Context, in *MountWebdavRequest, opts ...grpc.CallOption) (*MountWebdavResponse, error) {
	out := new(MountWebdavResponse)
	err := c.cc.Invoke(ctx, "/webdav.MountService/MountWebdav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mountServiceClient) UmountWebdav(ctx context.Context, in *UmountWebdavRequest, opts ...grpc.CallOption) (*UmountWebdavResponse, error) {
	out := new(UmountWebdavResponse)
	err := c.cc.Invoke(ctx, "/webdav.MountService/UmountWebdav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MountServiceServer is the server API for MountService service.
type MountServiceServer interface {
	MountWebdav(context.Context, *MountWebdavRequest) (*MountWebdavResponse, error)
	UmountWebdav(context.Context, *UmountWebdavRequest) (*UmountWebdavResponse, error)
}

// UnimplementedMountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMountServiceServer struct {
}

func (*UnimplementedMountServiceServer) MountWebdav(context.Context, *MountWebdavRequest) (*MountWebdavResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountWebdav not implemented")
}
func (*UnimplementedMountServiceServer) UmountWebdav(context.Context, *UmountWebdavRequest) (*UmountWebdavResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UmountWebdav not implemented")
}

func RegisterMountServiceServer(s *grpc.Server, srv MountServiceServer) {
	s.RegisterService(&_MountService_serviceDesc, srv)
}

func _MountService_MountWebdav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountWebdavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountServiceServer).MountWebdav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webdav.MountService/MountWebdav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountServiceServer).MountWebdav(ctx, req.(*MountWebdavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MountService_UmountWebdav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UmountWebdavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MountServiceServer).UmountWebdav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webdav.MountService/UmountWebdav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MountServiceServer).UmountWebdav(ctx, req.(*UmountWebdavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "webdav.MountService",
	HandlerType: (*MountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MountWebdav",
			Handler:    _MountService_MountWebdav_Handler,
		},
		{
			MethodName: "UmountWebdav",
			Handler:    _MountService_UmountWebdav_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webdav.proto",
}