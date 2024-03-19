package protobq

import (
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Marshal writes the given proto.Message in BigQuery format using default options.
func Marshal(msg proto.Message) (map[string]bigquery.Value, error) {
	return MarshalOptions{}.Marshal(msg)
}

// MarshalOptions is a configurable BigQuery format marshaler.
type MarshalOptions struct {
	// Schema contains the schema options.
	Schema SchemaOptions
}

// Marshal marshals the given proto.Message in the BigQuery format using options in
// MarshalOptions.
func (o MarshalOptions) Marshal(msg proto.Message) (map[string]bigquery.Value, error) {
	return o.marshalMessage(msg.ProtoReflect())
}

// marshalMessage marshals the given protoreflect.Message.
func (o MarshalOptions) marshalMessage(msg protoreflect.Message) (map[string]bigquery.Value, error) {
	result := make(map[string]bigquery.Value, msg.Descriptor().Fields().Len())
	var returnErr error
	msg.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		switch {
		case field.IsMap():
			m, err := o.marshalMapValue(field, value)
			if err != nil {
				returnErr = err
				return false
			}
			result[string(field.Name())] = m
		case field.IsList():
			l := make([]bigquery.Value, 0, value.List().Len())
			for i := 0; i < value.List().Len(); i++ {
				f, err := o.marshalValue(field, value.List().Get(i))
				if err != nil {
					returnErr = err
					return false
				}
				l = append(l, f)
			}
			result[string(field.Name())] = l
		default:
			column, errMarshal := o.marshalValue(field, value)
			if errMarshal != nil {
				returnErr = errMarshal
				return false
			}
			if m, ok := column.(map[string]bigquery.Value); ok && len(m) == 0 {
				// don't set anything for empty records
				return true
			}
			result[string(field.Name())] = column
		}
		return true
	})
	if returnErr != nil {
		return nil, returnErr
	}
	if o.Schema.UseOneofFields {
		for i := 0; i < msg.Descriptor().Oneofs().Len(); i++ {
			oneofDescriptor := msg.Descriptor().Oneofs().Get(i)
			oneofField := msg.WhichOneof(oneofDescriptor)
			if oneofField != nil {
				result[string(oneofDescriptor.Name())] = string(oneofField.Name())
			}
		}
	}
	return result, nil
}

func (o MarshalOptions) marshalMapValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) (bigquery.Value, error) {
	switch field.MapKey().Kind() {
	case protoreflect.StringKind:
		return o.marshalStringMapValue(field, value)
	case
		protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		return o.marshalIntMapValue(field, value)
	case
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		return o.marshalUintMapValue(field, value)
	case protoreflect.BoolKind:
		return o.marshalBoolMapValue(field, value)
	case
		protoreflect.EnumKind,
		protoreflect.BytesKind,
		protoreflect.FloatKind, protoreflect.DoubleKind,
		protoreflect.GroupKind, protoreflect.MessageKind:
		fallthrough
	default:
		return nil, fmt.Errorf("unsupported map key kind: %s", field.MapKey().Kind())
	}
}

// marshalStringMapValue marshals the given protoreflect.Value as a map with string keys.
func (o MarshalOptions) marshalStringMapValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) ([]bigquery.Value, error) {
	result := make([]bigquery.Value, 0, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result = append(result, map[string]bigquery.Value{
			"key":   key.String(),
			"value": v,
		})
		return true
	})
	if returnErr != nil {
		return nil, returnErr
	}
	return result, nil
}

// marshalIntMapValue marshals the given protoreflect.Value as a map with int keys.
func (o MarshalOptions) marshalIntMapValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) ([]bigquery.Value, error) {
	result := make([]bigquery.Value, 0, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result = append(result, map[string]bigquery.Value{
			"key":   key.Int(),
			"value": v,
		})
		return true
	})
	if returnErr != nil {
		return nil, returnErr
	}
	return result, nil
}

// marshalUintMapValue marshals the given protoreflect.Value as a map with int keys.
func (o MarshalOptions) marshalUintMapValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) ([]bigquery.Value, error) {
	result := make([]bigquery.Value, 0, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result = append(result, map[string]bigquery.Value{
			"key":   key.Uint(),
			"value": v,
		})
		return true
	})
	if returnErr != nil {
		return nil, returnErr
	}
	return result, nil
}

// marshalBoolMapValue marshals the given protoreflect.Value as a map with bool keys.
func (o MarshalOptions) marshalBoolMapValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) ([]bigquery.Value, error) {
	result := make([]bigquery.Value, 0, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result = append(result, map[string]bigquery.Value{
			"key":   key.Bool(),
			"value": v,
		})
		return true
	})
	if returnErr != nil {
		return nil, returnErr
	}
	return result, nil
}

// marshalValue marshals the given protoreflect.Value.
func (o MarshalOptions) marshalValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) (bigquery.Value, error) {
	switch field.Kind() {
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return value.Float(), nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		return value.Int(), nil
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return value.Uint(), nil
	case protoreflect.BoolKind:
		return value.Bool(), nil
	case protoreflect.StringKind:
		return value.String(), nil
	case protoreflect.BytesKind:
		return value.Bytes(), nil
	case protoreflect.EnumKind:
		return o.marshalEnumValue(field, value)
	case protoreflect.GroupKind, protoreflect.MessageKind:
		if wkt.IsWellKnownType(string(field.Message().FullName())) {
			return o.marshalWellKnownTypeValue(field, value)
		}
		return o.marshalMessage(value.Message())
	default:
		return nil, fmt.Errorf("unsupported field type: %v", field.Name())
	}
}

func (o MarshalOptions) marshalEnumValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) (bigquery.Value, error) {
	enumNumber := value.Enum()
	if o.Schema.UseEnumNumbers {
		return int64(enumNumber), nil
	}
	if enumValue := field.Enum().Values().ByNumber(enumNumber); enumValue != nil {
		return string(enumValue.Name()), nil
	}
	return nil, fmt.Errorf("unknown enum number: %v", value.Enum())
}

func (o MarshalOptions) marshalWellKnownTypeValue(
	field protoreflect.FieldDescriptor,
	value protoreflect.Value,
) (bigquery.Value, error) {
	switch field.Message().FullName() {
	case wkt.Timestamp:
		return value.Message().Interface().(*timestamppb.Timestamp).AsTime(), nil
	case wkt.Duration:
		return value.Message().Interface().(*durationpb.Duration).AsDuration().Seconds(), nil
	case wkt.DoubleValue:
		return value.Message().Interface().(*wrapperspb.DoubleValue).GetValue(), nil
	case wkt.FloatValue:
		return float64(value.Message().Interface().(*wrapperspb.FloatValue).GetValue()), nil
	case wkt.Int32Value:
		return int64(value.Message().Interface().(*wrapperspb.Int32Value).GetValue()), nil
	case wkt.Int64Value:
		return value.Message().Interface().(*wrapperspb.Int64Value).GetValue(), nil
	case wkt.UInt32Value:
		return uint64(value.Message().Interface().(*wrapperspb.UInt32Value).GetValue()), nil
	case wkt.UInt64Value:
		return value.Message().Interface().(*wrapperspb.UInt64Value).GetValue(), nil
	case wkt.BoolValue:
		return value.Message().Interface().(*wrapperspb.BoolValue).GetValue(), nil
	case wkt.StringValue:
		return value.Message().Interface().(*wrapperspb.StringValue).GetValue(), nil
	case wkt.BytesValue:
		return value.Message().Interface().(*wrapperspb.BytesValue).GetValue(), nil
	case wkt.Struct:
		data, err := value.Message().Interface().(*structpb.Struct).MarshalJSON()
		if err != nil {
			return nil, err
		}
		return string(data), nil
	case wkt.Date:
		d, ok := value.Message().Interface().(*date.Date)
		if !ok {
			return nil, fmt.Errorf("unexpected value for %s: %v", wkt.Date, value)
		}
		return civil.Date{
			Year:  int(d.GetYear()),
			Month: time.Month(d.GetMonth()),
			Day:   int(d.GetDay()),
		}, nil
	case wkt.DateTime:
		d, ok := value.Message().Interface().(*datetime.DateTime)
		if !ok {
			return nil, fmt.Errorf("unexpected value for %s: %v", wkt.DateTime, value)
		}
		if o.Schema.UseDateTimeWithoutOffset {
			return civil.DateTime{
				Date: civil.Date{
					Year:  int(d.GetYear()),
					Month: time.Month(d.GetMonth()),
					Day:   int(d.GetDay()),
				},
				Time: civil.Time{
					Hour:       int(d.GetHours()),
					Minute:     int(d.GetMinutes()),
					Second:     int(d.GetSeconds()),
					Nanosecond: int(d.GetNanos()),
				},
			}, nil
		}
		return nil, fmt.Errorf("TODO: implement support for google.type.DateTime with offset")
	case wkt.LatLng:
		latLng, ok := value.Message().Interface().(*latlng.LatLng)
		if !ok {
			return nil, fmt.Errorf("unexpected value for %s: %v", wkt.LatLng, value)
		}
		return fmt.Sprintf("POINT(%f %f)", latLng.GetLongitude(), latLng.GetLatitude()), nil
	case wkt.TimeOfDay:
		timeOfDay, ok := value.Message().Interface().(*timeofday.TimeOfDay)
		if !ok {
			return nil, fmt.Errorf("unexpected value for %s: %v", wkt.TimeOfDay, value)
		}
		return civil.Time{
			Hour:       int(timeOfDay.GetHours()),
			Minute:     int(timeOfDay.GetMinutes()),
			Second:     int(timeOfDay.GetSeconds()),
			Nanosecond: int(timeOfDay.GetNanos()),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported well-known-type %s", field.Message().FullName())
	}
}
