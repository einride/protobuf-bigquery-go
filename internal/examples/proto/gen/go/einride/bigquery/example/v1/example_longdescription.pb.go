// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_longdescription.proto

package examplev1

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

type ExampleLongDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pharetra rhoncus odio congue sodales. Suspendisse id nisi congue, maximus lectus lobortis, viverra magna. Donec viverra mauris at justo blandit, non posuere lacus dapibus. Pellentesque nec neque aliquet augue vehicula fringilla et et diam. Praesent eget lacus arcu. Donec consectetur arcu odio, eu pulvinar libero porta sit amet. In rutrum elit at mattis posuere. Aenean ac facilisis lacus. Fusce nisi arcu, rutrum varius felis sit amet, fermentum molestie lectus. Aliquam et erat rutrum, bibendum nisl non, posuere quam.
	// Duis eu tincidunt sapien, hendrerit vulputate dui. In tristique eu urna ut tempus. Donec ullamcorper tincidunt mi a hendrerit. Maecenas viverra ornare nunc, id bibendum nunc iaculis in. Nam faucibus maximus nisi, nec vehicula mi porta nec. Vivamus ultrices risus et velit blandit, vitae mattis erat viverra. Aliquam sed elit eu turpis tempus maximus. Nullam at ipsum non urna iaculis facilisis et nec orci. Pellentesque habitant morbi tristique senectus et lectus.
	FieldWithLongDescription int64 `protobuf:"varint,1,opt,name=field_with_long_description,json=fieldWithLongDescription,proto3" json:"field_with_long_description,omitempty"` // The above description is 1053 bytes
}

func (x *ExampleLongDescription) Reset() {
	*x = ExampleLongDescription{}
	mi := &file_einride_bigquery_example_v1_example_longdescription_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleLongDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleLongDescription) ProtoMessage() {}

func (x *ExampleLongDescription) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_longdescription_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleLongDescription.ProtoReflect.Descriptor instead.
func (*ExampleLongDescription) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_longdescription_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleLongDescription) GetFieldWithLongDescription() int64 {
	if x != nil {
		return x.FieldWithLongDescription
	}
	return 0
}

var File_einride_bigquery_example_v1_example_longdescription_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_longdescription_proto_rawDesc = []byte{
	0x0a, 0x39, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x16, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x4c, 0x6f, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x4c, 0x6f, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0xb2, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x6f,
	0x6e, 0x67, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x42, 0x45, 0xaa,
	0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5c, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a,
	0x3a, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_example_v1_example_longdescription_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_longdescription_proto_rawDescData = file_einride_bigquery_example_v1_example_longdescription_proto_rawDesc
)

func file_einride_bigquery_example_v1_example_longdescription_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_longdescription_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_longdescription_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_example_v1_example_longdescription_proto_rawDescData)
	})
	return file_einride_bigquery_example_v1_example_longdescription_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_longdescription_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_example_v1_example_longdescription_proto_goTypes = []any{
	(*ExampleLongDescription)(nil), // 0: einride.bigquery.example.v1.ExampleLongDescription
}
var file_einride_bigquery_example_v1_example_longdescription_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_longdescription_proto_init() }
func file_einride_bigquery_example_v1_example_longdescription_proto_init() {
	if File_einride_bigquery_example_v1_example_longdescription_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_bigquery_example_v1_example_longdescription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_longdescription_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_longdescription_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_example_v1_example_longdescription_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_longdescription_proto = out.File
	file_einride_bigquery_example_v1_example_longdescription_proto_rawDesc = nil
	file_einride_bigquery_example_v1_example_longdescription_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_longdescription_proto_depIdxs = nil
}
