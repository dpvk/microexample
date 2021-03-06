// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: svc_dummy.proto

package api

import (
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

type DummyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client string `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Orch   string `protobuf:"bytes,2,opt,name=orch,proto3" json:"orch,omitempty"`
}

func (x *DummyRequest) Reset() {
	*x = DummyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_dummy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DummyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyRequest) ProtoMessage() {}

func (x *DummyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dummy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyRequest.ProtoReflect.Descriptor instead.
func (*DummyRequest) Descriptor() ([]byte, []int) {
	return file_svc_dummy_proto_rawDescGZIP(), []int{0}
}

func (x *DummyRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *DummyRequest) GetOrch() string {
	if x != nil {
		return x.Orch
	}
	return ""
}

type DumyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy string `protobuf:"bytes,5,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *DumyResponse) Reset() {
	*x = DumyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svc_dummy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumyResponse) ProtoMessage() {}

func (x *DumyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dummy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumyResponse.ProtoReflect.Descriptor instead.
func (*DumyResponse) Descriptor() ([]byte, []int) {
	return file_svc_dummy_proto_rawDescGZIP(), []int{1}
}

func (x *DumyResponse) GetDummy() string {
	if x != nil {
		return x.Dummy
	}
	return ""
}

var File_svc_dummy_proto protoreflect.FileDescriptor

var file_svc_dummy_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x76, 0x63, 0x5f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x74, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x0c,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6f, 0x72, 0x63, 0x68, 0x22, 0x24, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x32, 0x46,
	0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x3d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x17, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x75, 0x6d, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_dummy_proto_rawDescOnce sync.Once
	file_svc_dummy_proto_rawDescData = file_svc_dummy_proto_rawDesc
)

func file_svc_dummy_proto_rawDescGZIP() []byte {
	file_svc_dummy_proto_rawDescOnce.Do(func() {
		file_svc_dummy_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_dummy_proto_rawDescData)
	})
	return file_svc_dummy_proto_rawDescData
}

var file_svc_dummy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_svc_dummy_proto_goTypes = []interface{}{
	(*DummyRequest)(nil), // 0: kubertest.DummyRequest
	(*DumyResponse)(nil), // 1: kubertest.DumyResponse
}
var file_svc_dummy_proto_depIdxs = []int32{
	0, // 0: kubertest.Dummy.Process:input_type -> kubertest.DummyRequest
	1, // 1: kubertest.Dummy.Process:output_type -> kubertest.DumyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_dummy_proto_init() }
func file_svc_dummy_proto_init() {
	if File_svc_dummy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svc_dummy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DummyRequest); i {
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
		file_svc_dummy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DumyResponse); i {
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
			RawDescriptor: file_svc_dummy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_dummy_proto_goTypes,
		DependencyIndexes: file_svc_dummy_proto_depIdxs,
		MessageInfos:      file_svc_dummy_proto_msgTypes,
	}.Build()
	File_svc_dummy_proto = out.File
	file_svc_dummy_proto_rawDesc = nil
	file_svc_dummy_proto_goTypes = nil
	file_svc_dummy_proto_depIdxs = nil
}
