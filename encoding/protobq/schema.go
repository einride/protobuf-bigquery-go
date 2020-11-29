package protobq

import (
	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Schema(m proto.Message) bigquery.Schema {
	return schemaForMessageDescriptor(m.ProtoReflect().Descriptor())
}

func schemaForMessageDescriptor(message protoreflect.MessageDescriptor) bigquery.Schema {
	schema := make(bigquery.Schema, 0, message.Fields().Len())
	hasAddedOneOf := map[protoreflect.FullName]bool{}
	for i := 0; i < message.Fields().Len(); i++ {
		field := message.Fields().Get(i)
		if field.IsMap() {
			continue // TODO: support maps
		}
		if oneOf := field.ContainingOneof(); oneOf != nil && !hasAddedOneOf[oneOf.FullName()] {
			schema = append(schema, &bigquery.FieldSchema{
				Name: string(oneOf.Name()),
				Type: bigquery.StringFieldType,
			})
			hasAddedOneOf[oneOf.FullName()] = true
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
				fieldSchema.Schema = schemaForMessageDescriptor(field.Message())
			}
		}
		schema = append(schema, fieldSchema)
	}
	return schema
}
