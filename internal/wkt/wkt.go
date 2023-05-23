package wkt

const (
	Timestamp   = "google.protobuf.Timestamp"
	Duration    = "google.protobuf.Duration"
	Struct      = "google.protobuf.Struct"
	Value       = "google.protobuf.Value"
	TimeOfDay   = "google.type.TimeOfDay"
	Date        = "google.type.Date"
	DateTime    = "google.type.DateTime"
	LatLng      = "google.type.LatLng"
	DoubleValue = "google.protobuf.DoubleValue"
	FloatValue  = "google.protobuf.FloatValue"
	Int32Value  = "google.protobuf.Int32Value"
	Int64Value  = "google.protobuf.Int64Value"
	UInt32Value = "google.protobuf.UInt32Value"
	UInt64Value = "google.protobuf.UInt64Value"
	BoolValue   = "google.protobuf.BoolValue"
	StringValue = "google.protobuf.StringValue"
	BytesValue  = "google.protobuf.BytesValue"
)

func IsWellKnownType(t string) bool {
	switch t {
	case Timestamp,
		Duration,
		Struct,
		TimeOfDay,
		Date,
		DateTime,
		LatLng,
		DoubleValue,
		FloatValue,
		Int32Value,
		Int64Value,
		UInt32Value,
		UInt64Value,
		BoolValue,
		StringValue,
		BytesValue:
		return true
	default:
		return false
	}
}
