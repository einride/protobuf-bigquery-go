// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: einride/bigquery/public/v1/film_location.proto

package publicv1

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

// Protobuf schema for the BigQuery public table:
//
//  bigquery-public-data.san_francisco_film_locations.film_locations
type FilmLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`                                                  // STRING NULLABLE
	ReleaseYear       int64  `protobuf:"varint,2,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`                  // INTEGER NULLABLE
	Locations         string `protobuf:"bytes,3,opt,name=locations,proto3" json:"locations,omitempty"`                                          // STRING NULLABLE
	FunFacts          string `protobuf:"bytes,4,opt,name=fun_facts,json=funFacts,proto3" json:"fun_facts,omitempty"`                            // STRING NULLABLE
	ProductionCompany string `protobuf:"bytes,5,opt,name=production_company,json=productionCompany,proto3" json:"production_company,omitempty"` // STRING NULLABLE
	Distributor       string `protobuf:"bytes,6,opt,name=distributor,proto3" json:"distributor,omitempty"`                                      // STRING NULLABLE
	Director          string `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`                                            // STRING NULLABLE
	Writer            string `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer,omitempty"`                                                // STRING NULLABLE
	Actor_1           string `protobuf:"bytes,9,opt,name=actor_1,json=actor1,proto3" json:"actor_1,omitempty"`                                  // STRING NULLABLE
	Actor_2           string `protobuf:"bytes,10,opt,name=actor_2,json=actor2,proto3" json:"actor_2,omitempty"`                                 // STRING NULLABLE
	Actor_3           string `protobuf:"bytes,11,opt,name=actor_3,json=actor3,proto3" json:"actor_3,omitempty"`                                 // STRING NULLABLE
}

func (x *FilmLocation) Reset() {
	*x = FilmLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_film_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmLocation) ProtoMessage() {}

func (x *FilmLocation) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_film_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmLocation.ProtoReflect.Descriptor instead.
func (*FilmLocation) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_film_location_proto_rawDescGZIP(), []int{0}
}

func (x *FilmLocation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FilmLocation) GetReleaseYear() int64 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *FilmLocation) GetLocations() string {
	if x != nil {
		return x.Locations
	}
	return ""
}

func (x *FilmLocation) GetFunFacts() string {
	if x != nil {
		return x.FunFacts
	}
	return ""
}

func (x *FilmLocation) GetProductionCompany() string {
	if x != nil {
		return x.ProductionCompany
	}
	return ""
}

func (x *FilmLocation) GetDistributor() string {
	if x != nil {
		return x.Distributor
	}
	return ""
}

func (x *FilmLocation) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *FilmLocation) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *FilmLocation) GetActor_1() string {
	if x != nil {
		return x.Actor_1
	}
	return ""
}

func (x *FilmLocation) GetActor_2() string {
	if x != nil {
		return x.Actor_2
	}
	return ""
}

func (x *FilmLocation) GetActor_3() string {
	if x != nil {
		return x.Actor_3
	}
	return ""
}

var File_einride_bigquery_public_v1_film_location_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_film_location_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x22, 0xd2, 0x02, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x46, 0x61, 0x63, 0x74,
	0x73, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x31, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x31, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x33, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x33, 0x42, 0xa3, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x46, 0x69, 0x6c, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x61, 0x67, 0x6f, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65,
	0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45,
	0x42, 0x50, 0xaa, 0x02, 0x1a, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x42, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1b, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x3a, 0x3a, 0x42, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x3a, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_film_location_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_film_location_proto_rawDescData = file_einride_bigquery_public_v1_film_location_proto_rawDesc
)

func file_einride_bigquery_public_v1_film_location_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_film_location_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_film_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_film_location_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_film_location_proto_rawDescData
}

var file_einride_bigquery_public_v1_film_location_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_film_location_proto_goTypes = []interface{}{
	(*FilmLocation)(nil), // 0: einride.bigquery.public.v1.FilmLocation
}
var file_einride_bigquery_public_v1_film_location_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_film_location_proto_init() }
func file_einride_bigquery_public_v1_film_location_proto_init() {
	if File_einride_bigquery_public_v1_film_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_film_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmLocation); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_film_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_film_location_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_film_location_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_film_location_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_film_location_proto = out.File
	file_einride_bigquery_public_v1_film_location_proto_rawDesc = nil
	file_einride_bigquery_public_v1_film_location_proto_goTypes = nil
	file_einride_bigquery_public_v1_film_location_proto_depIdxs = nil
}
