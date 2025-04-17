// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/san_fransisco_transit_stop_time.proto

package publicv1

import (
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
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
//	bigquery-public-data.san_francisco_transit_muni.stop_times
type SanFransiscoTransitStopTime struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	StopId         int64                  `protobuf:"varint,1,opt,name=stop_id,json=stopId,proto3" json:"stop_id,omitempty"`                           // INTEGER NULLABLE
	TripId         int64                  `protobuf:"varint,2,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`                           // INTEGER NULLABLE
	StopSequence   int64                  `protobuf:"varint,3,opt,name=stop_sequence,json=stopSequence,proto3" json:"stop_sequence,omitempty"`         // INTEGER NULLABLE
	ArrivalTime    *timeofday.TimeOfDay   `protobuf:"bytes,4,opt,name=arrival_time,json=arrivalTime,proto3" json:"arrival_time,omitempty"`             // TIME NULLABLE
	ArrivesNextDay bool                   `protobuf:"varint,5,opt,name=arrives_next_day,json=arrivesNextDay,proto3" json:"arrives_next_day,omitempty"` // BOOLEAN NULLABLE
	DepartureTime  *timeofday.TimeOfDay   `protobuf:"bytes,6,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty"`       // TIME NULLABLE
	DepartsNextDay bool                   `protobuf:"varint,7,opt,name=departs_next_day,json=departsNextDay,proto3" json:"departs_next_day,omitempty"` // BOOLEAN NULLABLE
	DropoffType    string                 `protobuf:"bytes,8,opt,name=dropoff_type,json=dropoffType,proto3" json:"dropoff_type,omitempty"`             // STRING NULLABLE
	ExactTimepoint bool                   `protobuf:"varint,9,opt,name=exact_timepoint,json=exactTimepoint,proto3" json:"exact_timepoint,omitempty"`   // BOOLEAN NULLABLE
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SanFransiscoTransitStopTime) Reset() {
	*x = SanFransiscoTransitStopTime{}
	mi := &file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SanFransiscoTransitStopTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanFransiscoTransitStopTime) ProtoMessage() {}

func (x *SanFransiscoTransitStopTime) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanFransiscoTransitStopTime.ProtoReflect.Descriptor instead.
func (*SanFransiscoTransitStopTime) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescGZIP(), []int{0}
}

func (x *SanFransiscoTransitStopTime) GetStopId() int64 {
	if x != nil {
		return x.StopId
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetTripId() int64 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetStopSequence() int64 {
	if x != nil {
		return x.StopSequence
	}
	return 0
}

func (x *SanFransiscoTransitStopTime) GetArrivalTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.ArrivalTime
	}
	return nil
}

func (x *SanFransiscoTransitStopTime) GetArrivesNextDay() bool {
	if x != nil {
		return x.ArrivesNextDay
	}
	return false
}

func (x *SanFransiscoTransitStopTime) GetDepartureTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

func (x *SanFransiscoTransitStopTime) GetDepartsNextDay() bool {
	if x != nil {
		return x.DepartsNextDay
	}
	return false
}

func (x *SanFransiscoTransitStopTime) GetDropoffType() string {
	if x != nil {
		return x.DropoffType
	}
	return ""
}

func (x *SanFransiscoTransitStopTime) GetExactTimepoint() bool {
	if x != nil {
		return x.ExactTimepoint
	}
	return false
}

var File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto protoreflect.FileDescriptor

const file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc = "" +
	"\n" +
	"@einride/bigquery/public/v1/san_fransisco_transit_stop_time.proto\x12\x1aeinride.bigquery.public.v1\x1a\x1bgoogle/type/timeofday.proto\"\x8e\x03\n" +
	"\x1bSanFransiscoTransitStopTime\x12\x17\n" +
	"\astop_id\x18\x01 \x01(\x03R\x06stopId\x12\x17\n" +
	"\atrip_id\x18\x02 \x01(\x03R\x06tripId\x12#\n" +
	"\rstop_sequence\x18\x03 \x01(\x03R\fstopSequence\x129\n" +
	"\farrival_time\x18\x04 \x01(\v2\x16.google.type.TimeOfDayR\varrivalTime\x12(\n" +
	"\x10arrives_next_day\x18\x05 \x01(\bR\x0earrivesNextDay\x12=\n" +
	"\x0edeparture_time\x18\x06 \x01(\v2\x16.google.type.TimeOfDayR\rdepartureTime\x12(\n" +
	"\x10departs_next_day\x18\a \x01(\bR\x0edepartsNextDay\x12!\n" +
	"\fdropoff_type\x18\b \x01(\tR\vdropoffType\x12'\n" +
	"\x0fexact_timepoint\x18\t \x01(\bR\x0eexactTimepointB\xb2\x02\n" +
	"\x1ecom.einride.bigquery.public.v1B SanFransiscoTransitStopTimeProtoP\x01Zago.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/public/v1;publicv1\xa2\x02\x03EBP\xaa\x02\x1aEinride.Bigquery.Public.V1\xca\x02\x1bEinride\\Bigquery\\Public_\\V1\xe2\x02'Einride\\Bigquery\\Public_\\V1\\GPBMetadata\xea\x02\x1dEinride::Bigquery::Public::V1b\x06proto3"

var (
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData []byte
)

func file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc), len(file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc)))
	})
	return file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDescData
}

var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes = []any{
	(*SanFransiscoTransitStopTime)(nil), // 0: einride.bigquery.public.v1.SanFransiscoTransitStopTime
	(*timeofday.TimeOfDay)(nil),         // 1: google.type.TimeOfDay
}
var file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.SanFransiscoTransitStopTime.arrival_time:type_name -> google.type.TimeOfDay
	1, // 1: einride.bigquery.public.v1.SanFransiscoTransitStopTime.departure_time:type_name -> google.type.TimeOfDay
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_init() }
func file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_init() {
	if File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc), len(file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto = out.File
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_goTypes = nil
	file_einride_bigquery_public_v1_san_fransisco_transit_stop_time_proto_depIdxs = nil
}
