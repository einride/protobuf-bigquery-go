// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: einride/bigquery/example/v1/example_list.proto

package examplev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
	state          protoimpl.MessageState   `protogen:"open.v1"`
	Int64List      []int64                  `protobuf:"varint,1,rep,packed,name=int64_list,json=int64List,proto3" json:"int64_list,omitempty"`
	StringList     []string                 `protobuf:"bytes,2,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
	EnumList       []ExampleList_Enum       `protobuf:"varint,3,rep,packed,name=enum_list,json=enumList,proto3,enum=einride.bigquery.example.v1.ExampleList_Enum" json:"enum_list,omitempty"`
	NestedList     []*ExampleList_Nested    `protobuf:"bytes,4,rep,name=nested_list,json=nestedList,proto3" json:"nested_list,omitempty"`
	FloatValueList []*wrapperspb.FloatValue `protobuf:"bytes,5,rep,name=float_value_list,json=floatValueList,proto3" json:"float_value_list,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ExampleList) Reset() {
	*x = ExampleList{}
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleList) ProtoMessage() {}

func (x *ExampleList) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[0]
	if x != nil {
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

func (x *ExampleList) GetFloatValueList() []*wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValueList
	}
	return nil
}

type ExampleList_Nested struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StringList    []string               `protobuf:"bytes,1,rep,name=string_list,json=stringList,proto3" json:"string_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExampleList_Nested) Reset() {
	*x = ExampleList_Nested{}
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExampleList_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleList_Nested) ProtoMessage() {}

func (x *ExampleList_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_example_v1_example_list_proto_msgTypes[1]
	if x != nil {
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

const file_einride_bigquery_example_v1_example_list_proto_rawDesc = "" +
	"\n" +
	".einride/bigquery/example/v1/example_list.proto\x12\x1beinride.bigquery.example.v1\x1a\x1egoogle/protobuf/wrappers.proto\"\x9d\x03\n" +
	"\vExampleList\x12\x1d\n" +
	"\n" +
	"int64_list\x18\x01 \x03(\x03R\tint64List\x12\x1f\n" +
	"\vstring_list\x18\x02 \x03(\tR\n" +
	"stringList\x12J\n" +
	"\tenum_list\x18\x03 \x03(\x0e2-.einride.bigquery.example.v1.ExampleList.EnumR\benumList\x12P\n" +
	"\vnested_list\x18\x04 \x03(\v2/.einride.bigquery.example.v1.ExampleList.NestedR\n" +
	"nestedList\x12E\n" +
	"\x10float_value_list\x18\x05 \x03(\v2\x1b.google.protobuf.FloatValueR\x0efloatValueList\x1a)\n" +
	"\x06Nested\x12\x1f\n" +
	"\vstring_list\x18\x01 \x03(\tR\n" +
	"stringList\">\n" +
	"\x04Enum\x12\x14\n" +
	"\x10ENUM_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vENUM_VALUE1\x10\x01\x12\x0f\n" +
	"\vENUM_VALUE2\x10\x02B\xa7\x02\n" +
	"\x1fcom.einride.bigquery.example.v1B\x10ExampleListProtoP\x01Zcgo.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/example/v1;examplev1\xa2\x02\x03EBE\xaa\x02\x1bEinride.Bigquery.Example.V1\xca\x02\x1bEinride\\Bigquery\\Example\\V1\xe2\x02'Einride\\Bigquery\\Example\\V1\\GPBMetadata\xea\x02\x1eEinride::Bigquery::Example::V1b\x06proto3"

var (
	file_einride_bigquery_example_v1_example_list_proto_rawDescOnce sync.Once
	file_einride_bigquery_example_v1_example_list_proto_rawDescData []byte
)

func file_einride_bigquery_example_v1_example_list_proto_rawDescGZIP() []byte {
	file_einride_bigquery_example_v1_example_list_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_example_v1_example_list_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_list_proto_rawDesc), len(file_einride_bigquery_example_v1_example_list_proto_rawDesc)))
	})
	return file_einride_bigquery_example_v1_example_list_proto_rawDescData
}

var file_einride_bigquery_example_v1_example_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_bigquery_example_v1_example_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_bigquery_example_v1_example_list_proto_goTypes = []any{
	(ExampleList_Enum)(0),         // 0: einride.bigquery.example.v1.ExampleList.Enum
	(*ExampleList)(nil),           // 1: einride.bigquery.example.v1.ExampleList
	(*ExampleList_Nested)(nil),    // 2: einride.bigquery.example.v1.ExampleList.Nested
	(*wrapperspb.FloatValue)(nil), // 3: google.protobuf.FloatValue
}
var file_einride_bigquery_example_v1_example_list_proto_depIdxs = []int32{
	0, // 0: einride.bigquery.example.v1.ExampleList.enum_list:type_name -> einride.bigquery.example.v1.ExampleList.Enum
	2, // 1: einride.bigquery.example.v1.ExampleList.nested_list:type_name -> einride.bigquery.example.v1.ExampleList.Nested
	3, // 2: einride.bigquery.example.v1.ExampleList.float_value_list:type_name -> google.protobuf.FloatValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_einride_bigquery_example_v1_example_list_proto_init() }
func file_einride_bigquery_example_v1_example_list_proto_init() {
	if File_einride_bigquery_example_v1_example_list_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_example_v1_example_list_proto_rawDesc), len(file_einride_bigquery_example_v1_example_list_proto_rawDesc)),
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
	file_einride_bigquery_example_v1_example_list_proto_goTypes = nil
	file_einride_bigquery_example_v1_example_list_proto_depIdxs = nil
}
