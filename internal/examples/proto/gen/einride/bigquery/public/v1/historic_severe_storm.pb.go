// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: einride/bigquery/public/v1/historic_severe_storm.proto

package publicv1

import (
	proto "github.com/golang/protobuf/proto"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// Protobuf schema for the BigQuery public table:
//
//  bigquery-public-data.noaa_historic_severe_storms.storms_*
type HistoricSevereStorm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpisodeId         string             `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`                          // STRING NULLABLE
	EventId           string             `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`                                // STRING NULLABLE
	State             string             `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`                                                   // STRING NULLABLE
	StateFipsCode     string             `protobuf:"bytes,4,opt,name=state_fips_code,json=stateFipsCode,proto3" json:"state_fips_code,omitempty"`            // STRING NULLABLE
	EventType         string             `protobuf:"bytes,5,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`                          // STRING NULLABLE
	CzType            string             `protobuf:"bytes,6,opt,name=cz_type,json=czType,proto3" json:"cz_type,omitempty"`                                   // STRING NULLABLE
	CzFipsCode        string             `protobuf:"bytes,7,opt,name=cz_fips_code,json=czFipsCode,proto3" json:"cz_fips_code,omitempty"`                     // STRING NULLABLE
	CzName            string             `protobuf:"bytes,8,opt,name=cz_name,json=czName,proto3" json:"cz_name,omitempty"`                                   // STRING NULLABLE
	Wfo               string             `protobuf:"bytes,9,opt,name=wfo,proto3" json:"wfo,omitempty"`                                                       // STRING NULLABLE
	EventBeginTime    *datetime.DateTime `protobuf:"bytes,10,opt,name=event_begin_time,json=eventBeginTime,proto3" json:"event_begin_time,omitempty"`        // DATETIME NULLABLE
	EventTimezone     string             `protobuf:"bytes,11,opt,name=event_timezone,json=eventTimezone,proto3" json:"event_timezone,omitempty"`             // STRING NULLABLE
	EventEndTime      *datetime.DateTime `protobuf:"bytes,12,opt,name=event_end_time,json=eventEndTime,proto3" json:"event_end_time,omitempty"`              // DATETIME NULLABLE
	InjuriesDirect    int64              `protobuf:"varint,13,opt,name=injuries_direct,json=injuriesDirect,proto3" json:"injuries_direct,omitempty"`         // INTEGER NULLABLE
	InjuriesIndirect  int64              `protobuf:"varint,14,opt,name=injuries_indirect,json=injuriesIndirect,proto3" json:"injuries_indirect,omitempty"`   // INTEGER NULLABLE
	DeathsDirect      int64              `protobuf:"varint,15,opt,name=deaths_direct,json=deathsDirect,proto3" json:"deaths_direct,omitempty"`               // INTEGER NULLABLE
	DeathsIndirect    int64              `protobuf:"varint,16,opt,name=deaths_indirect,json=deathsIndirect,proto3" json:"deaths_indirect,omitempty"`         // INTEGER NULLABLE
	DamageProperty    int64              `protobuf:"varint,17,opt,name=damage_property,json=damageProperty,proto3" json:"damage_property,omitempty"`         // INTEGER NULLABLE
	DamageCrops       int64              `protobuf:"varint,18,opt,name=damage_crops,json=damageCrops,proto3" json:"damage_crops,omitempty"`                  // INTEGER NULLABLE
	Source            string             `protobuf:"bytes,19,opt,name=source,proto3" json:"source,omitempty"`                                                // STRING NULLABLE
	Magnitude         float64            `protobuf:"fixed64,20,opt,name=magnitude,proto3" json:"magnitude,omitempty"`                                        // FLOAT NULLABLE
	MagnitudeType     string             `protobuf:"bytes,21,opt,name=magnitude_type,json=magnitudeType,proto3" json:"magnitude_type,omitempty"`             // STRING NULLABLE
	FloodCause        string             `protobuf:"bytes,22,opt,name=flood_cause,json=floodCause,proto3" json:"flood_cause,omitempty"`                      // STRING NULLABLE
	TorFScale         string             `protobuf:"bytes,23,opt,name=tor_f_scale,json=torFScale,proto3" json:"tor_f_scale,omitempty"`                       // STRING NULLABLE
	TorLength         string             `protobuf:"bytes,24,opt,name=tor_length,json=torLength,proto3" json:"tor_length,omitempty"`                         // STRING NULLABLE
	TorWidth          string             `protobuf:"bytes,25,opt,name=tor_width,json=torWidth,proto3" json:"tor_width,omitempty"`                            // STRING NULLABLE
	TorOtherWfo       string             `protobuf:"bytes,26,opt,name=tor_other_wfo,json=torOtherWfo,proto3" json:"tor_other_wfo,omitempty"`                 // STRING NULLABLE
	LocationIndex     string             `protobuf:"bytes,27,opt,name=location_index,json=locationIndex,proto3" json:"location_index,omitempty"`             // STRING NULLABLE
	EventRange        float64            `protobuf:"fixed64,28,opt,name=event_range,json=eventRange,proto3" json:"event_range,omitempty"`                    // FLOAT NULLABLE
	EventAzimuth      string             `protobuf:"bytes,29,opt,name=event_azimuth,json=eventAzimuth,proto3" json:"event_azimuth,omitempty"`                // STRING NULLABLE
	ReferenceLocation string             `protobuf:"bytes,30,opt,name=reference_location,json=referenceLocation,proto3" json:"reference_location,omitempty"` // STRING NULLABLE
	EventLatitude     float64            `protobuf:"fixed64,31,opt,name=event_latitude,json=eventLatitude,proto3" json:"event_latitude,omitempty"`           // FLOAT NULLABLE
	EventLongitude    float64            `protobuf:"fixed64,32,opt,name=event_longitude,json=eventLongitude,proto3" json:"event_longitude,omitempty"`        // FLOAT NULLABLE
	EventPoint        *latlng.LatLng     `protobuf:"bytes,33,opt,name=event_point,json=eventPoint,proto3" json:"event_point,omitempty"`                      // GEOGRAPHY NULLABLE
}

func (x *HistoricSevereStorm) Reset() {
	*x = HistoricSevereStorm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_bigquery_public_v1_historic_severe_storm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricSevereStorm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricSevereStorm) ProtoMessage() {}

func (x *HistoricSevereStorm) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_historic_severe_storm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricSevereStorm.ProtoReflect.Descriptor instead.
func (*HistoricSevereStorm) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescGZIP(), []int{0}
}

func (x *HistoricSevereStorm) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *HistoricSevereStorm) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *HistoricSevereStorm) GetStateFipsCode() string {
	if x != nil {
		return x.StateFipsCode
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *HistoricSevereStorm) GetCzType() string {
	if x != nil {
		return x.CzType
	}
	return ""
}

func (x *HistoricSevereStorm) GetCzFipsCode() string {
	if x != nil {
		return x.CzFipsCode
	}
	return ""
}

func (x *HistoricSevereStorm) GetCzName() string {
	if x != nil {
		return x.CzName
	}
	return ""
}

func (x *HistoricSevereStorm) GetWfo() string {
	if x != nil {
		return x.Wfo
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventBeginTime() *datetime.DateTime {
	if x != nil {
		return x.EventBeginTime
	}
	return nil
}

func (x *HistoricSevereStorm) GetEventTimezone() string {
	if x != nil {
		return x.EventTimezone
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventEndTime() *datetime.DateTime {
	if x != nil {
		return x.EventEndTime
	}
	return nil
}

func (x *HistoricSevereStorm) GetInjuriesDirect() int64 {
	if x != nil {
		return x.InjuriesDirect
	}
	return 0
}

func (x *HistoricSevereStorm) GetInjuriesIndirect() int64 {
	if x != nil {
		return x.InjuriesIndirect
	}
	return 0
}

func (x *HistoricSevereStorm) GetDeathsDirect() int64 {
	if x != nil {
		return x.DeathsDirect
	}
	return 0
}

func (x *HistoricSevereStorm) GetDeathsIndirect() int64 {
	if x != nil {
		return x.DeathsIndirect
	}
	return 0
}

func (x *HistoricSevereStorm) GetDamageProperty() int64 {
	if x != nil {
		return x.DamageProperty
	}
	return 0
}

func (x *HistoricSevereStorm) GetDamageCrops() int64 {
	if x != nil {
		return x.DamageCrops
	}
	return 0
}

func (x *HistoricSevereStorm) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *HistoricSevereStorm) GetMagnitude() float64 {
	if x != nil {
		return x.Magnitude
	}
	return 0
}

func (x *HistoricSevereStorm) GetMagnitudeType() string {
	if x != nil {
		return x.MagnitudeType
	}
	return ""
}

func (x *HistoricSevereStorm) GetFloodCause() string {
	if x != nil {
		return x.FloodCause
	}
	return ""
}

func (x *HistoricSevereStorm) GetTorFScale() string {
	if x != nil {
		return x.TorFScale
	}
	return ""
}

func (x *HistoricSevereStorm) GetTorLength() string {
	if x != nil {
		return x.TorLength
	}
	return ""
}

func (x *HistoricSevereStorm) GetTorWidth() string {
	if x != nil {
		return x.TorWidth
	}
	return ""
}

func (x *HistoricSevereStorm) GetTorOtherWfo() string {
	if x != nil {
		return x.TorOtherWfo
	}
	return ""
}

func (x *HistoricSevereStorm) GetLocationIndex() string {
	if x != nil {
		return x.LocationIndex
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventRange() float64 {
	if x != nil {
		return x.EventRange
	}
	return 0
}

func (x *HistoricSevereStorm) GetEventAzimuth() string {
	if x != nil {
		return x.EventAzimuth
	}
	return ""
}

func (x *HistoricSevereStorm) GetReferenceLocation() string {
	if x != nil {
		return x.ReferenceLocation
	}
	return ""
}

func (x *HistoricSevereStorm) GetEventLatitude() float64 {
	if x != nil {
		return x.EventLatitude
	}
	return 0
}

func (x *HistoricSevereStorm) GetEventLongitude() float64 {
	if x != nil {
		return x.EventLongitude
	}
	return 0
}

func (x *HistoricSevereStorm) GetEventPoint() *latlng.LatLng {
	if x != nil {
		return x.EventPoint
	}
	return nil
}

var File_einride_bigquery_public_v1_historic_severe_storm_proto protoreflect.FileDescriptor

var file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61,
	0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x09, 0x0a, 0x13, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x53, 0x65, 0x76, 0x65, 0x72, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x70, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x70, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x7a, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x7a, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x7a, 0x5f, 0x66, 0x69, 0x70, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x7a, 0x46, 0x69, 0x70, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x7a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x77, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x66, 0x6f, 0x12,
	0x3f, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x6a, 0x75, 0x72, 0x69, 0x65, 0x73,
	0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69,
	0x6e, 0x6a, 0x75, 0x72, 0x69, 0x65, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a,
	0x11, 0x69, 0x6e, 0x6a, 0x75, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x69, 0x6e, 0x6a, 0x75, 0x72, 0x69,
	0x65, 0x73, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x61, 0x74, 0x68, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x74, 0x68, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x65, 0x61, 0x74, 0x68, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x65, 0x61, 0x74, 0x68, 0x73,
	0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x70,
	0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x43,
	0x72, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61,
	0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x67, 0x6e, 0x69, 0x74, 0x75, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x6f, 0x64, 0x43, 0x61, 0x75,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x5f, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x72, 0x46, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x72, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x72, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x72, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x22,
	0x0a, 0x0d, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x77, 0x66, 0x6f, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x57,
	0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x7a, 0x69, 0x6d, 0x75, 0x74, 0x68, 0x18, 0x1d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x7a, 0x69, 0x6d, 0x75, 0x74, 0x68, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x34,
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescData = file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDesc
)

func file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescData)
	})
	return file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDescData
}

var file_einride_bigquery_public_v1_historic_severe_storm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_einride_bigquery_public_v1_historic_severe_storm_proto_goTypes = []interface{}{
	(*HistoricSevereStorm)(nil), // 0: einride.bigquery.public.v1.HistoricSevereStorm
	(*datetime.DateTime)(nil),   // 1: google.type.DateTime
	(*latlng.LatLng)(nil),       // 2: google.type.LatLng
}
var file_einride_bigquery_public_v1_historic_severe_storm_proto_depIdxs = []int32{
	1, // 0: einride.bigquery.public.v1.HistoricSevereStorm.event_begin_time:type_name -> google.type.DateTime
	1, // 1: einride.bigquery.public.v1.HistoricSevereStorm.event_end_time:type_name -> google.type.DateTime
	2, // 2: einride.bigquery.public.v1.HistoricSevereStorm.event_point:type_name -> google.type.LatLng
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_historic_severe_storm_proto_init() }
func file_einride_bigquery_public_v1_historic_severe_storm_proto_init() {
	if File_einride_bigquery_public_v1_historic_severe_storm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_bigquery_public_v1_historic_severe_storm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricSevereStorm); i {
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
			RawDescriptor: file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_historic_severe_storm_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_historic_severe_storm_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_historic_severe_storm_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_historic_severe_storm_proto = out.File
	file_einride_bigquery_public_v1_historic_severe_storm_proto_rawDesc = nil
	file_einride_bigquery_public_v1_historic_severe_storm_proto_goTypes = nil
	file_einride_bigquery_public_v1_historic_severe_storm_proto_depIdxs = nil
}