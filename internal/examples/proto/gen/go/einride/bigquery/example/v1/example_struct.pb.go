// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_struct.proto

package examplev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleStruct struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         *structpb.Value        `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleStruct) Reset() {
	*x = ExampleStruct{}
	mi := &file_einride_bigquery_example_v1_example_struct_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleStruct) ProtoMessage() {}

func (x *ExampleStruct) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_struct_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleStruct.ProtoReflect.Descriptor instead.
func (*ExampleStruct) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_struct_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleStruct) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_einride_bigquery_example_v1_example_struct_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_struct_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a,
	0x0d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xa9, 0x02, 0x0a,
	0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2d, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42,
	0x45, 0xaa, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_einride_bigquery_example_v1_example_struct_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_struct_proto_rawDescData []byte
)

func file_einride_bigquery_example_v1_example_struct_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_struct_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_struct_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_struct_proto_rawDesc), len(file_einride_bigquery_example_v1_example_struct_proto_rawDesc)))
	})
	return file_einride_bigquery_example_v1_example_struct_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_example_v1_example_struct_proto_goTypes = []any{
	(*ExampleStruct)(nil),  // 0: einride.bigquery.example.v1.ExampleStruct
	(*structpb.Value)(nil), // 1: google.protobuf.Value
}
var file_einride_bigquery_example_v1_example_struct_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.example.v1.ExampleStruct.value:type_name -> google.protobuf.Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_struct_proto_init() }
func file_einride_bigquery_example_v1_example_struct_proto_init() {
	if File_einride_bigquery_example_v1_example_struct_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_struct_proto_rawDesc), len(file_einride_bigquery_example_v1_example_struct_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_struct_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_struct_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_example_v1_example_struct_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_struct_proto = out.File
	file_einride_bigquery_example_v1_example_struct_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_struct_proto_depIdxs = nil
}
