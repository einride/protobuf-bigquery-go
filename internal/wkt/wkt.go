package wkt

const (
	Timestamp = "google.protobuf.Timestamp"
	Duration  = "google.protobuf.Duration"
	Struct    = "google.protobuf.Struct"
	TimeOfDay = "google.type.TimeOfDay"
	Date      = "google.type.Date"
	LatLng    = "google.type.LatLng"
)

func IsWellKnownType(t string) bool {
	switch t {
	case Timestamp,
		Duration,
		Struct,
		TimeOfDay,
		Date,
		LatLng:
		return true
	default:
		return false
	}
}
