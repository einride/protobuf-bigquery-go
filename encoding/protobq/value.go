package protobq

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Marshal(m proto.Message) (map[string]bigquery.Value, error) {
	return marshalMessage(m.ProtoReflect())
}

func marshalMessage(m protoreflect.Message) (map[string]bigquery.Value, error) {
	v := make(map[string]bigquery.Value, m.Descriptor().Fields().Len())
	var err error
	m.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		if oneOf := field.ContainingOneof(); oneOf != nil {
			v[string(oneOf.Name())] = string(field.Name())
			return true
		}
		if field.IsMap() {
			return true // TODO
		}
		if field.IsList() {
			l := make([]bigquery.Value, 0, value.List().Len())
			for i := 0; i < value.List().Len(); i++ {
				f, errF := marshalField(field, value.List().Get(i))
				if errF != nil {
					err = errF
					return false
				}
				l = append(l, f)
			}
			v[string(field.Name())] = l
			return true
		}
		f, errF := marshalField(field, value)
		if errF != nil {
			err = errF
			return false
		}
		v[string(field.Name())] = f
		return true
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}

func marshalField(field protoreflect.FieldDescriptor, value protoreflect.Value) (bigquery.Value, error) {
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
		case "google.protobuf.Timestamp":
			return value.Message().Interface().(*timestamppb.Timestamp).AsTime(), nil
		case "google.protobuf.Duration":
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
		default:
			return marshalMessage(value.Message())
		}
	default:
		return nil, fmt.Errorf("unsupported field type: %v", field.Name())
	}
}
