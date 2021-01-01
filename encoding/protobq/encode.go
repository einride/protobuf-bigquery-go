package protobq

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
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
			result[string(field.Name())] = column
		}
		return true
	})
	if returnErr != nil {
		return nil, returnErr
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
) (map[string]bigquery.Value, error) {
	result := make(map[string]bigquery.Value, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result[key.String()] = v
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
) (map[int64]bigquery.Value, error) {
	result := make(map[int64]bigquery.Value, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result[key.Int()] = v
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
) (map[uint64]bigquery.Value, error) {
	result := make(map[uint64]bigquery.Value, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result[key.Uint()] = v
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
) (map[bool]bigquery.Value, error) {
	result := make(map[bool]bigquery.Value, value.Map().Len())
	var returnErr error
	value.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		v, err := o.marshalValue(field.MapValue(), value)
		if err != nil {
			returnErr = err
			return false
		}
		result[key.Bool()] = v
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
	case protoreflect.DoubleKind,
		protoreflect.FloatKind:
		return value.Float(), nil
	case protoreflect.Int64Kind,
		protoreflect.Int32Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind:
		return value.Int(), nil
	case protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind,
		protoreflect.Fixed32Kind,
		protoreflect.Uint32Kind:
		return value.Uint(), nil
	case protoreflect.BoolKind:
		return value.Bool(), nil
	case protoreflect.StringKind:
		return value.String(), nil
	case protoreflect.GroupKind:
		return nil, nil // ignore
	case protoreflect.BytesKind:
		return value.Bytes(), nil
	case protoreflect.EnumKind:
		if enumValue := field.Enum().Values().ByNumber(value.Enum()); enumValue != nil {
			return string(enumValue.Name()), nil
		}
		return nil, fmt.Errorf("unknown enum number: %v", value.Enum())
	case protoreflect.MessageKind:
		switch field.Message().FullName() {
		case wkt.Timestamp:
			return value.Message().Interface().(*timestamppb.Timestamp).AsTime(), nil
		case wkt.Duration:
			return value.Message().Interface().(*durationpb.Duration).AsDuration().Seconds(), nil
		case "google.protobuf.DoubleValue":
			return value.Message().Interface().(*wrapperspb.DoubleValue).GetValue(), nil
		case "google.protobuf.FloatValue":
			return value.Message().Interface().(*wrapperspb.FloatValue).GetValue(), nil
		case "google.protobuf.Int32Value":
			return value.Message().Interface().(*wrapperspb.Int32Value).GetValue(), nil
		case "google.protobuf.Int64Value":
			return value.Message().Interface().(*wrapperspb.Int64Value).GetValue(), nil
		case "google.protobuf.UInt32Value":
			return value.Message().Interface().(*wrapperspb.UInt32Value).GetValue(), nil
		case "google.protobuf.UInt64Value":
			return value.Message().Interface().(*wrapperspb.UInt64Value).GetValue(), nil
		case "google.protobuf.BoolValue":
			return value.Message().Interface().(*wrapperspb.BoolValue).GetValue(), nil
		case "google.protobuf.StringValue":
			return value.Message().Interface().(*wrapperspb.StringValue).GetValue(), nil
		case "google.protobuf.BytesValue":
			return value.Message().Interface().(*wrapperspb.BytesValue).GetValue(), nil
		case "google.protobuf.StructValue":
			data, err := value.Message().Interface().(*structpb.Struct).MarshalJSON()
			if err != nil {
				return nil, err
			}
			return string(data), nil
		case "google.type.Date":
			return nil, fmt.Errorf("TODO: implement support for google.type.Date")
		case "google.type.DateTime":
			return nil, fmt.Errorf("TODO: implement support for google.type.DateTime")
		case "google.type.LatLng":
			return nil, fmt.Errorf("TODO: implement support for google.type.LatLng")
		case "google.type.Time":
			return nil, fmt.Errorf("TODO: implement support for google.type.Time")
		default:
			return o.marshalMessage(value.Message())
		}
	default:
		return nil, fmt.Errorf("unsupported field type: %v", field.Name())
	}
}
