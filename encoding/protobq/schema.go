package protobq

import (
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// InferSchema infers a BigQuery schema for the given proto.Message using default options.
func InferSchema(msg proto.Message) bigquery.Schema {
	return SchemaOptions{}.InferSchema(msg)
}

// SchemaOptions contains configuration options for BigQuery schema inference.
type SchemaOptions struct {
	// UseEnumNumbers converts enum values to INTEGER types.
	UseEnumNumbers bool
	// UseDateTimeWithoutOffset converts google.type.DateTime values to DATETIME, discarding the optional time offset.
	UseDateTimeWithoutOffset bool
	// UseOneofFields adds an extra STRING field for oneof fields with the name of the oneof,
	// containing the name of the field that is set.
	UseOneofFields bool
	// UseOptionalFields causes any field not marked "optional" in the proto to be a REQUIRED field.
	// Lists and Maps are always NULLABLE as proto3 does not seem to support optional for those types.
	// Types inside maps are always marked as required.
	UseOptionalFields bool
}

// InferSchema infers a BigQuery schema for the given proto.Message using options in
// MarshalOptions.
func (o SchemaOptions) InferSchema(msg proto.Message) bigquery.Schema {
	return o.inferMessageSchema(msg.ProtoReflect().Descriptor())
}

// inferMessageSchema infers the BigQuery schema for the given protoreflect.MessageDescriptor.
func (o SchemaOptions) inferMessageSchema(msg protoreflect.MessageDescriptor) bigquery.Schema {
	schema := make(bigquery.Schema, 0, msg.Fields().Len())
	for i := 0; i < msg.Fields().Len(); i++ {
		fieldSchema := o.inferFieldSchema(msg.Fields().Get(i))
		if fieldSchema != nil {
			schema = append(schema, fieldSchema)
		}
	}
	if o.UseOneofFields {
		for i := 0; i < msg.Oneofs().Len(); i++ {
			oneof := msg.Oneofs().Get(i)
			schema = append(schema, o.inferOneofFieldSchema(oneof))
		}
	}
	return schema
}

func (o SchemaOptions) inferFieldSchema(field protoreflect.FieldDescriptor) *bigquery.FieldSchema {
	if field.IsMap() {
		return o.inferMapFieldSchema(field)
	}
	if (field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind) &&
		field.Message().FullName() == wkt.DateTime {
		return o.inferDateTimeFieldSchema(field)
	}
	fieldSchema := &bigquery.FieldSchema{
		Name:     string(field.Name()),
		Type:     o.inferFieldSchemaType(field),
		Repeated: field.IsList(),
		Required: o.UseOptionalFields && !field.HasOptionalKeyword() && !field.IsList(),
	}
	if fieldSchema.Type == bigquery.RecordFieldType && fieldSchema.Schema == nil {
		fieldSchema.Schema = o.inferMessageSchema(field.Message())
		if len(fieldSchema.Schema) == 0 {
			return nil
		}
	}
	return fieldSchema
}

func (o SchemaOptions) inferOneofFieldSchema(oneof protoreflect.OneofDescriptor) *bigquery.FieldSchema {
	fieldSchema := &bigquery.FieldSchema{
		Name: string(oneof.Name()),
		Type: bigquery.StringFieldType,
	}
	oneofFieldNames := make([]string, 0, oneof.Fields().Len())
	for i := 0; i < oneof.Fields().Len(); i++ {
		field := oneof.Fields().Get(i)
		oneofFieldNames = append(oneofFieldNames, string(field.Name()))
	}
	fieldSchema.Description = fmt.Sprintf("One of: %s.", strings.Join(oneofFieldNames, ", "))
	return fieldSchema
}

func (o SchemaOptions) inferDateTimeFieldSchema(field protoreflect.FieldDescriptor) *bigquery.FieldSchema {
	fieldSchema := &bigquery.FieldSchema{
		Name:     string(field.Name()),
		Repeated: field.IsList(),
		Required: o.UseOptionalFields && !field.HasOptionalKeyword() && !field.IsList(),
	}
	if o.UseDateTimeWithoutOffset {
		fieldSchema.Type = bigquery.DateTimeFieldType
	} else {
		fieldSchema.Type = bigquery.RecordFieldType
		fieldSchema.Schema = bigquery.Schema{
			{Name: "datetime", Type: bigquery.DateTimeFieldType},
			{Name: "utc_offset", Type: bigquery.FloatFieldType},
			{
				Name: "time_zone",
				Type: bigquery.RecordFieldType,
				Schema: bigquery.Schema{
					{Name: "id", Type: bigquery.StringFieldType},
					{Name: "version", Type: bigquery.StringFieldType},
				},
			},
		}
	}
	return fieldSchema
}

func (o SchemaOptions) inferFieldSchemaType(field protoreflect.FieldDescriptor) bigquery.FieldType {
	switch field.Kind() {
	case protoreflect.DoubleKind, protoreflect.FloatKind:
		return bigquery.FloatFieldType
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
		return bigquery.IntegerFieldType
	case protoreflect.BoolKind:
		return bigquery.BooleanFieldType
	case protoreflect.StringKind:
		return bigquery.StringFieldType
	case protoreflect.BytesKind:
		return bigquery.BytesFieldType
	case protoreflect.EnumKind:
		return o.inferEnumFieldType(field)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		switch field.Message().FullName() {
		case wkt.Timestamp:
			return bigquery.TimestampFieldType
		case wkt.Duration:
			return bigquery.FloatFieldType
		case wkt.DoubleValue, wkt.FloatValue:
			return bigquery.FloatFieldType
		case wkt.Int32Value, wkt.Int64Value, wkt.UInt32Value, wkt.UInt64Value:
			return bigquery.IntegerFieldType
		case wkt.BoolValue:
			return bigquery.BooleanFieldType
		case wkt.StringValue:
			return bigquery.StringFieldType
		case wkt.BytesValue:
			return bigquery.BytesFieldType
		case wkt.Struct:
			return bigquery.StringFieldType // JSON string
		case wkt.Date:
			return bigquery.DateFieldType
		case wkt.DateTime:
			if o.UseDateTimeWithoutOffset {
				return bigquery.DateTimeFieldType
			}
			return bigquery.RecordFieldType // to include explicit UTC offset or time zone
		case wkt.LatLng:
			return bigquery.GeographyFieldType
		case wkt.TimeOfDay:
			return bigquery.TimeFieldType
		}
	}
	return bigquery.RecordFieldType
}

func (o SchemaOptions) inferEnumFieldType(field protoreflect.FieldDescriptor) bigquery.FieldType {
	if o.UseEnumNumbers {
		return bigquery.IntegerFieldType
	}
	return bigquery.StringFieldType
}

func (o SchemaOptions) inferMapFieldSchema(field protoreflect.FieldDescriptor) *bigquery.FieldSchema {
	return &bigquery.FieldSchema{
		Name:     string(field.Name()),
		Repeated: true,
		Type:     bigquery.RecordFieldType,
		Schema: bigquery.Schema{
			o.inferFieldSchema(field.MapKey()),
			o.inferFieldSchema(field.MapValue()),
		},
	}
}
