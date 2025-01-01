// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_datetime.proto

package examplev1

import (
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

type ExampleDateTime struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DateTime      *datetime.DateTime     `protobuf:"bytes,1,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleDateTime) Reset() {
	*x = ExampleDateTime{}
	mi := &file_einride_bigquery_example_v1_example_datetime_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleDateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleDateTime) ProtoMessage() {}

func (x *ExampleDateTime) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_datetime_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleDateTime.ProtoReflect.Descriptor instead.
func (*ExampleDateTime) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_datetime_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleDateTime) GetDateTime() *datetime.DateTime {
	if x != nil {
		return x.DateTime
	}
	return nil
}

var File_einride_bigquery_example_v1_example_datetime_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_datetime_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x0f, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x32, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0xab, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x45, 0xaa, 0x02, 0x1b, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_example_v1_example_datetime_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_datetime_proto_rawDescData = file_einride_bigquery_example_v1_example_datetime_proto_rawDesc
)

func file_einride_bigquery_example_v1_example_datetime_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_datetime_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_datetime_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_example_v1_example_datetime_proto_rawDescData)
	})
	return file_einride_bigquery_example_v1_example_datetime_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_datetime_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_example_v1_example_datetime_proto_goTypes = []any{
	(*ExampleDateTime)(nil),   // 0: einride.bigquery.example.v1.ExampleDateTime
	(*datetime.DateTime)(nil), // 1: google.type.DateTime
}
var file_einride_bigquery_example_v1_example_datetime_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.example.v1.ExampleDateTime.date_time:type_name -> google.type.DateTime
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_datetime_proto_init() }
func file_einride_bigquery_example_v1_example_datetime_proto_init() {
	if File_einride_bigquery_example_v1_example_datetime_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_bigquery_example_v1_example_datetime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_datetime_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_datetime_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_example_v1_example_datetime_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_datetime_proto = out.File
	file_einride_bigquery_example_v1_example_datetime_proto_rawDesc = nil
	file_einride_bigquery_example_v1_example_datetime_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_datetime_proto_depIdxs = nil
}
