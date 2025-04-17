// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/london_bicycle_station.proto

package publicv1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Protobuf schema for the BigQuery public table:
//
//	bigquery-public-data.london_bicycles.cycle_stations
type LondonBicycleStation struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                   // INTEGER NULLABLE
	Installed  bool                   `protobuf:"varint,2,opt,name=installed,proto3" json:"installed,omitempty"`                     // BOOLEAN NULLABLE
	Latitude   float64                `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`                      // FLOAT NULLABLE
	Locked     string                 `protobuf:"bytes,4,opt,name=locked,proto3" json:"locked,omitempty"`                            // STRING NULLABLE
	Longitude  float64                `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`                    // FLOAT NULLABLE
	Name       string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`                                // STRING NULLABLE
	BikesCount int64                  `protobuf:"varint,7,opt,name=bikes_count,json=bikesCount,proto3" json:"bikes_count,omitempty"` // INTEGER NULLABLE
	DocksCount int64                  `protobuf:"varint,8,opt,name=docks_count,json=docksCount,proto3" json:"docks_count,omitempty"` // INTEGER NULLABLE
	// int64 nbEmptyDocks = 9; // INTEGER NULLABLE
	Temporary     bool       `protobuf:"varint,10,opt,name=temporary,proto3" json:"temporary,omitempty"`                          // BOOLEAN NULLABLE
	TerminalName  string     `protobuf:"bytes,11,opt,name=terminal_name,json=terminalName,proto3" json:"terminal_name,omitempty"` // STRING NULLABLE
	InstallDate   *date.Date `protobuf:"bytes,12,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`    // DATE NULLABLE
	RemovalDate   *date.Date `protobuf:"bytes,13,opt,name=removal_date,json=removalDate,proto3" json:"removal_date,omitempty"`    // DATE NULLABLE
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LondonBicycleStation) Reset() {
	*x = LondonBicycleStation{}
	mi := &file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LondonBicycleStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LondonBicycleStation) ProtoMessage() {}

func (x *LondonBicycleStation) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LondonBicycleStation.ProtoReflect.Descriptor instead.
func (*LondonBicycleStation) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescGZIP(), []int{0}
}

func (x *LondonBicycleStation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LondonBicycleStation) GetInstalled() bool {
	if x != nil {
		return x.Installed
	}
	return false
}

func (x *LondonBicycleStation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LondonBicycleStation) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *LondonBicycleStation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *LondonBicycleStation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LondonBicycleStation) GetBikesCount() int64 {
	if x != nil {
		return x.BikesCount
	}
	return 0
}

func (x *LondonBicycleStation) GetDocksCount() int64 {
	if x != nil {
		return x.DocksCount
	}
	return 0
}

func (x *LondonBicycleStation) GetTemporary() bool {
	if x != nil {
		return x.Temporary
	}
	return false
}

func (x *LondonBicycleStation) GetTerminalName() string {
	if x != nil {
		return x.TerminalName
	}
	return ""
}

func (x *LondonBicycleStation) GetInstallDate() *date.Date {
	if x != nil {
		return x.InstallDate
	}
	return nil
}

func (x *LondonBicycleStation) GetRemovalDate() *date.Date {
	if x != nil {
		return x.RemovalDate
	}
	return nil
}

var File_einride_bigquery_public_v1_london_bicycle_station_proto protoreflect.FileDescriptor

const file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc = "" +
	"\n" +
	"7einride/bigquery/public/v1/london_bicycle_station.proto\x12\x1aeinride.bigquery.public.v1\x1a\x16google/type/date.proto\"\x9b\x03\n" +
	"\x14LondonBicycleStation\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1c\n" +
	"\tinstalled\x18\x02 \x01(\bR\tinstalled\x12\x1a\n" +
	"\blatitude\x18\x03 \x01(\x01R\blatitude\x12\x16\n" +
	"\x06locked\x18\x04 \x01(\tR\x06locked\x12\x1c\n" +
	"\tlongitude\x18\x05 \x01(\x01R\tlongitude\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04name\x12\x1f\n" +
	"\vbikes_count\x18\a \x01(\x03R\n" +
	"bikesCount\x12\x1f\n" +
	"\vdocks_count\x18\b \x01(\x03R\n" +
	"docksCount\x12\x1c\n" +
	"\ttemporary\x18\n" +
	" \x01(\bR\ttemporary\x12#\n" +
	"\rterminal_name\x18\v \x01(\tR\fterminalName\x124\n" +
	"\finstall_date\x18\f \x01(\v2\x11.google.type.DateR\vinstallDate\x124\n" +
	"\fremoval_date\x18\r \x01(\v2\x11.google.type.DateR\vremovalDateB\xab\x02\n" +
	"\x1ecom.einride.bigquery.public.v1B\x19LondonBicycleStationProtoP\x01Zago.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/public/v1;publicv1\xa2\x02\x03EBP\xaa\x02\x1aEinride.Bigquery.Public.V1\xca\x02\x1bEinride\\Bigquery\\Public_\\V1\xe2\x02'Einride\\Bigquery\\Public_\\V1\\GPBMetadata\xea\x02\x1dEinride::Bigquery::Public::V1b\x06proto3"

var (
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData []byte
)

func file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc), len(file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc)))
	})
	return file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDescData
}

var file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes = []any{
	(*LondonBicycleStation)(nil), // 0: einride.bigquery.public.v1.LondonBicycleStation
	(*date.Date)(nil),            // 1: google.type.Date
}
var file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.LondonBicycleStation.install_date:type_name -> google.type.Date
	1, // 1: einride.bigquery.public.v1.LondonBicycleStation.removal_date:type_name -> google.type.Date
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_london_bicycle_station_proto_init() }
func file_einride_bigquery_public_v1_london_bicycle_station_proto_init() {
	if File_einride_bigquery_public_v1_london_bicycle_station_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc), len(file_einride_bigquery_public_v1_london_bicycle_station_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_london_bicycle_station_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_london_bicycle_station_proto = out.File
	file_einride_bigquery_public_v1_london_bicycle_station_proto_goTypes = nil
	file_einride_bigquery_public_v1_london_bicycle_station_proto_depIdxs = nil
}
