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

// Marshal writes the given proto.Message in BigQuery format using default options.
func Marshal(msg proto.Message) (map[string]bigquery.Value, error) {
	return MarshalOptions{}.Marshal(msg)
}

// InferSchema infers a BigQuery schema for the given proto.Message using default options.
func InferSchema(msg proto.Message) bigquery.Schema {
	return MarshalOptions{}.InferSchema(msg)
}

// MarshalOptions is a configurable BigQuery format marshaler.
type MarshalOptions struct {
}

// InferSchema infers a BigQuery schema for the given proto.Message using options in
// MarshalOptions.
func (o MarshalOptions) InferSchema(msg proto.Message) bigquery.Schema {
	return o.inferMessageSchema(msg.ProtoReflect().Descriptor())
}

// Marshal marshals the given proto.Message in the BigQuery format using options in
// MarshalOptions.
func (o MarshalOptions) Marshal(msg proto.Message) (map[string]bigquery.Value, error) {
	return o.marshalMessage(msg.ProtoReflect())
}

// inferMessageSchema infers the BigQuery schema for the given protoreflect.MessageDescriptor.
func (o MarshalOptions) inferMessageSchema(msg protoreflect.MessageDescriptor) bigquery.Schema {
	schema := make(bigquery.Schema, 0, msg.Fields().Len())
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if field.IsMap() {
			continue // TODO: support maps
		}
		fieldSchema := &bigquery.FieldSchema{
			Name:     string(field.Name()),
			Repeated: field.Cardinality() == protoreflect.Repeated,
		}
		switch field.Kind() {
		case protoreflect.DoubleKind,
			protoreflect.FloatKind:
			fieldSchema.Type = bigquery.FloatFieldType
		case protoreflect.Int64Kind,
			protoreflect.Uint64Kind,
			protoreflect.Int32Kind,
			protoreflect.Fixed64Kind,
			protoreflect.Fixed32Kind,
			protoreflect.Uint32Kind,
			protoreflect.Sfixed32Kind,
			protoreflect.Sfixed64Kind,
			protoreflect.Sint32Kind,
			protoreflect.Sint64Kind:
			fieldSchema.Type = bigquery.IntegerFieldType
		case protoreflect.BoolKind:
			fieldSchema.Type = bigquery.BooleanFieldType
		case protoreflect.StringKind:
			fieldSchema.Type = bigquery.StringFieldType
		case protoreflect.GroupKind:
			continue // ignore legacy proto2 group fields
		case protoreflect.BytesKind:
			fieldSchema.Type = bigquery.BytesFieldType
		case protoreflect.EnumKind:
			fieldSchema.Type = bigquery.StringFieldType
		case protoreflect.MessageKind:
			switch field.Message().FullName() {
			case "google.protobuf.Timestamp":
				fieldSchema.Type = bigquery.TimestampFieldType
			case "google.protobuf.Duration":
				fieldSchema.Type = bigquery.FloatFieldType
			case "google.protobuf.DoubleValue",
				"google.protobuf.FloatValue":
				fieldSchema.Type = bigquery.FloatFieldType
			case "google.protobuf.Int32Value",
				"google.protobuf.Int64Value",
				"google.protobuf.UInt32Value",
				"google.protobuf.UInt64Value":
				fieldSchema.Type = bigquery.IntegerFieldType
			case "google.protobuf.BoolValue":
				fieldSchema.Type = bigquery.BooleanFieldType
			case "google.protobuf.StringValue":
				fieldSchema.Type = bigquery.StringFieldType
			case "google.protobuf.BytesValue":
				fieldSchema.Type = bigquery.BytesFieldType
			case "google.protobuf.StructValue":
				fieldSchema.Type = bigquery.StringFieldType // JSON string
			default:
				fieldSchema.Type = bigquery.RecordFieldType
				fieldSchema.Schema = o.inferMessageSchema(field.Message())
			}
		}
		schema = append(schema, fieldSchema)
	}
	return schema
}

// marshalMessage marshals the given protoreflect.Message.
func (o MarshalOptions) marshalMessage(msg protoreflect.Message) (map[string]bigquery.Value, error) {
	row := make(map[string]bigquery.Value, msg.Descriptor().Fields().Len())
	var err error
	msg.Range(func(field protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		if field.IsMap() {
			return true // TODO
		}
		if field.IsList() {
			l := make([]bigquery.Value, 0, value.List().Len())
			for i := 0; i < value.List().Len(); i++ {
				f, errF := o.marshalValue(field, value.List().Get(i))
				if errF != nil {
					err = errF
					return false
				}
				l = append(l, f)
			}
			row[string(field.Name())] = l
			return true
		}
		column, errMarshal := o.marshalValue(field, value)
		if errMarshal != nil {
			err = errMarshal
			return false
		}
		row[string(field.Name())] = column
		return true
	})
	if err != nil {
		return nil, err
	}
	return row, nil
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
			return o.marshalMessage(value.Message())
		}
	default:
		return nil, fmt.Errorf("unsupported field type: %v", field.Name())
	}
}
