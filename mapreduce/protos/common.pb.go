// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: protos/common.proto

package master

import (
	proto "github.com/golang/protobuf/proto"
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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ShutdownReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NTasks int32 `protobuf:"varint,1,opt,name=nTasks,proto3" json:"nTasks,omitempty"`
}

func (x *ShutdownReply) Reset() {
	*x = ShutdownReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownReply) ProtoMessage() {}

func (x *ShutdownReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownReply.ProtoReflect.Descriptor instead.
func (*ShutdownReply) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{1}
}

func (x *ShutdownReply) GetNTasks() int32 {
	if x != nil {
		return x.NTasks
	}
	return 0
}

type DoTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName       string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName,omitempty"`
	Filename      string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Phase         string `protobuf:"bytes,3,opt,name=phase,proto3" json:"phase,omitempty"`
	TaskNo        int32  `protobuf:"varint,4,opt,name=taskNo,proto3" json:"taskNo,omitempty"`
	NumOtherPhase int32  `protobuf:"varint,5,opt,name=numOtherPhase,proto3" json:"numOtherPhase,omitempty"`
}

func (x *DoTaskRequest) Reset() {
	*x = DoTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTaskRequest) ProtoMessage() {}

func (x *DoTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTaskRequest.ProtoReflect.Descriptor instead.
func (*DoTaskRequest) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{2}
}

func (x *DoTaskRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *DoTaskRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DoTaskRequest) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *DoTaskRequest) GetTaskNo() int32 {
	if x != nil {
		return x.TaskNo
	}
	return 0
}

func (x *DoTaskRequest) GetNumOtherPhase() int32 {
	if x != nil {
		return x.NumOtherPhase
	}
	return 0
}

var File_protos_common_proto protoreflect.FileDescriptor

var file_protos_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x2b, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d,
	0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x68, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x68, 0x61, 0x73, 0x65, 0x42,
	0x27, 0x5a, 0x25, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_proto_rawDescOnce sync.Once
	file_protos_common_proto_rawDescData = file_protos_common_proto_rawDesc
)

func file_protos_common_proto_rawDescGZIP() []byte {
	file_protos_common_proto_rawDescOnce.Do(func() {
		file_protos_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_proto_rawDescData)
	})
	return file_protos_common_proto_rawDescData
}

var file_protos_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_common_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil), // 0: common.RegisterRequest
	(*ShutdownReply)(nil),   // 1: common.ShutdownReply
	(*DoTaskRequest)(nil),   // 2: common.DoTaskRequest
}
var file_protos_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_common_proto_init() }
func file_protos_common_proto_init() {
	if File_protos_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_protos_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownReply); i {
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
		file_protos_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTaskRequest); i {
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
			RawDescriptor: file_protos_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_proto_goTypes,
		DependencyIndexes: file_protos_common_proto_depIdxs,
		MessageInfos:      file_protos_common_proto_msgTypes,
	}.Build()
	File_protos_common_proto = out.File
	file_protos_common_proto_rawDesc = nil
	file_protos_common_proto_goTypes = nil
	file_protos_common_proto_depIdxs = nil
}
