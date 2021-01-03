// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: einride/bigquery/example/v1/example_list.proto

package examplev1

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

type ExampleList_Enum int32

const (
	ExampleList_ENUM_UNSPECIFIED ExampleList_Enum = 0
	ExampleList_ENUM_VALUE1      ExampleList_Enum = 1
	ExampleList_ENUM_VALUE2      ExampleList_Enum = 2
)

// Enum value maps for ExampleList_Enum.
var (
	ExampleList_Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_VALUE1",
		2: "ENUM_VALUE2",
	}
	ExampleList_Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_VALUE1":      1,
		"ENUM_VALUE2":      2,
	}
)

func (x ExampleList_Enum) Enum() *ExampleList_Enum {
	p := new(ExampleList_Enum)
	*p = x
	return p
}

func (x ExampleList_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleList_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_bigquery_example_v1_example_list_proto_enumTypes[0].Descriptor()
}

func (ExampleList_Enum) Type() protoreflect.EnumType {
	return &file_einride_bigquery_example_v1_example_list_proto_enumTypes[0]
}

func (x ExampleList_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleList_Enum.Descriptor instead.
func (ExampleList_Enum) EnumDescriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_list_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int64List  []int64               `protobuf:"varint,1,rep,packed,name=int64_list,json=int64List,proto3" json:"int64_list,omitempty"`
	StringList []string              `protobuf:"bytes,2,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
	EnumList   []ExampleList_Enum    `protobuf:"varint,3,rep,packed,name=enum_list,json=enumList,proto3,enum=einride.bigquery.example.v1.ExampleList_Enum" json:"enum_list,omitempty"`
	NestedList []*ExampleList_Nested `protobuf:"bytes,4,rep,name=nested_list,json=nestedList,proto3" json:"nested_list,omitempty"`
}

func (x *ExampleList) Reset() {
	*x = ExampleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleList) ProtoMessage() {}

func (x *ExampleList) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleList.ProtoReflect.Descriptor instead.
func (*ExampleList) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_list_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleList) GetInt64List() []int64 {
	if x != nil {
		return x.Int64List
	}
	return nil
}

func (x *ExampleList) GetStringList() []string {
	if x != nil {
		return x.StringList
	}
	return nil
}

func (x *ExampleList) GetEnumList() []ExampleList_Enum {
	if x != nil {
		return x.EnumList
	}
	return nil
}

func (x *ExampleList) GetNestedList() []*ExampleList_Nested {
	if x != nil {
		return x.NestedList
	}
	return nil
}

type ExampleList_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringList []string `protobuf:"bytes,1,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
}

func (x *ExampleList_Nested) Reset() {
	*x = ExampleList_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleList_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleList_Nested) ProtoMessage() {}

func (x *ExampleList_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleList_Nested.ProtoReflect.Descriptor instead.
func (*ExampleList_Nested) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_example_v1_example_list_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExampleList_Nested) GetStringList() []string {
	if x != nil {
		return x.StringList
	}
	return nil
}

var File_einride_bigquery_example_v1_example_list_proto protoreflect.FileDescriptor

var file_einride_bigquery_example_v1_example_list_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xd6, 0x02,
	0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x2d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x08, 0x65, 0x6e, 0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0b, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x0a, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x29, 0x0a, 0x06, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x32, 0x10, 0x02, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_example_v1_example_list_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_list_proto_rawDescData = file_einride_bigquery_example_v1_example_list_proto_rawDesc
)

func file_einride_bigquery_example_v1_example_list_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_list_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_example_v1_example_list_proto_rawDescData)
	})
	return file_einride_bigquery_example_v1_example_list_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_bigquery_example_v1_example_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_bigquery_example_v1_example_list_proto_goTypes = []interface{}{
	(ExampleList_Enum)(0),      // 0: einride.bigquery.example.v1.ExampleList.Enum
	(*ExampleList)(nil),        // 1: einride.bigquery.example.v1.ExampleList
	(*ExampleList_Nested)(nil), // 2: einride.bigquery.example.v1.ExampleList.Nested
}
var file_einride_bigquery_example_v1_example_list_proto_depIdxs = []int32{
	0, // 0: einride.bigquery.example.v1.ExampleList.enum_list:type_name -> einride.bigquery.example.v1.ExampleList.Enum
	2, // 1: einride.bigquery.example.v1.ExampleList.nested_list:type_name -> einride.bigquery.example.v1.ExampleList.Nested
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_list_proto_init() }
func file_einride_bigquery_example_v1_example_list_proto_init() {
	if File_einride_bigquery_example_v1_example_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_example_v1_example_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleList); i {
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
		file_einride_bigquery_example_v1_example_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleList_Nested); i {
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
			RawDescriptor: file_einride_bigquery_example_v1_example_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_example_v1_example_list_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_example_v1_example_list_proto_depIdxs,
		EnumInfos:         file_einride_bigquery_example_v1_example_list_proto_enumTypes,
		MessageInfos:      file_einride_bigquery_example_v1_example_list_proto_msgTypes,
	}.Build()
	File_einride_bigquery_example_v1_example_list_proto = out.File
	file_einride_bigquery_example_v1_example_list_proto_rawDesc = nil
	file_einride_bigquery_example_v1_example_list_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_list_proto_depIdxs = nil
}
