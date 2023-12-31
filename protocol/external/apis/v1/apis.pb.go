// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: external/apis/v1/apis.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListModelsRequest) Reset() {
	*x = ListModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_apis_v1_apis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsRequest) ProtoMessage() {}

func (x *ListModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_apis_v1_apis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsRequest.ProtoReflect.Descriptor instead.
func (*ListModelsRequest) Descriptor() ([]byte, []int) {
	return file_external_apis_v1_apis_proto_rawDescGZIP(), []int{0}
}

type ListModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *ListModelsResponse) Reset() {
	*x = ListModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_apis_v1_apis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsResponse) ProtoMessage() {}

func (x *ListModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_apis_v1_apis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsResponse.ProtoReflect.Descriptor instead.
func (*ListModelsResponse) Descriptor() ([]byte, []int) {
	return file_external_apis_v1_apis_proto_rawDescGZIP(), []int{1}
}

func (x *ListModelsResponse) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_apis_v1_apis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_external_apis_v1_apis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_external_apis_v1_apis_proto_rawDescGZIP(), []int{2}
}

func (x *Model) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_external_apis_v1_apis_proto protoreflect.FileDescriptor

var file_external_apis_v1_apis_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22,
	0x2b, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x57, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x42, 0x86, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6f, 0x63, 0x68, 0x6f, 0x77, 0x2d, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x65, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x92,
	0x41, 0x4c, 0x12, 0x13, 0x0a, 0x0c, 0x41, 0x50, 0x49, 0x73, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_apis_v1_apis_proto_rawDescOnce sync.Once
	file_external_apis_v1_apis_proto_rawDescData = file_external_apis_v1_apis_proto_rawDesc
)

func file_external_apis_v1_apis_proto_rawDescGZIP() []byte {
	file_external_apis_v1_apis_proto_rawDescOnce.Do(func() {
		file_external_apis_v1_apis_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_apis_v1_apis_proto_rawDescData)
	})
	return file_external_apis_v1_apis_proto_rawDescData
}

var file_external_apis_v1_apis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_apis_v1_apis_proto_goTypes = []interface{}{
	(*ListModelsRequest)(nil),  // 0: ListModelsRequest
	(*ListModelsResponse)(nil), // 1: ListModelsResponse
	(*Model)(nil),              // 2: Model
}
var file_external_apis_v1_apis_proto_depIdxs = []int32{
	2, // 0: ListModelsResponse.models:type_name -> Model
	0, // 1: ApiService.ListModels:input_type -> ListModelsRequest
	1, // 2: ApiService.ListModels:output_type -> ListModelsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_apis_v1_apis_proto_init() }
func file_external_apis_v1_apis_proto_init() {
	if File_external_apis_v1_apis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_apis_v1_apis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelsRequest); i {
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
		file_external_apis_v1_apis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelsResponse); i {
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
		file_external_apis_v1_apis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
			RawDescriptor: file_external_apis_v1_apis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_apis_v1_apis_proto_goTypes,
		DependencyIndexes: file_external_apis_v1_apis_proto_depIdxs,
		MessageInfos:      file_external_apis_v1_apis_proto_msgTypes,
	}.Build()
	File_external_apis_v1_apis_proto = out.File
	file_external_apis_v1_apis_proto_rawDesc = nil
	file_external_apis_v1_apis_proto_goTypes = nil
	file_external_apis_v1_apis_proto_depIdxs = nil
}
