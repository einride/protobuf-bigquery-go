package protobq

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Unmarshal the bigquery.Value map into the given proto.Message.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func Unmarshal(row map[string]bigquery.Value, message proto.Message) error {
	return UnmarshalOptions{}.Unmarshal(row, message)
}

// Load the bigquery.Value list into the given proto.Message using the given bigquery.Schema.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func Load(row []bigquery.Value, schema bigquery.Schema, message proto.Message) error {
	return UnmarshalOptions{}.Load(row, schema, message)
}

// UnmarshalOptions is a configurable BigQuery format parser.
type UnmarshalOptions struct {
	// If AllowPartial is set, input for messages that will result in missing
	// required fields will not return an error.
	AllowPartial bool

	// If DiscardUnknown is set, unknown fields are ignored.
	DiscardUnknown bool
}

// Unmarshal reads the given BigQuery row and populates the given proto.Message using
// options in UnmarshalOptions object.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func (o UnmarshalOptions) Unmarshal(row map[string]bigquery.Value, message proto.Message) error {
	proto.Reset(message)
	if err := o.unmarshalMessage(row, message.ProtoReflect()); err != nil {
		return err
	}
	if o.AllowPartial {
		return nil
	}
	return proto.CheckInitialized(message)
}

// Load the bigquery.Value list into the given proto.Message using the given bigquery.Schema
// using options in UnmarshalOptions object.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func (o UnmarshalOptions) Load(row []bigquery.Value, schema bigquery.Schema, message proto.Message) error {
	proto.Reset(message)
	if err := o.loadMessage(row, schema, message.ProtoReflect()); err != nil {
		return err
	}
	if o.AllowPartial {
		return nil
	}
	return proto.CheckInitialized(message)
}

func (o UnmarshalOptions) loadMessage(
	bigQueryFields []bigquery.Value,
	schema bigquery.Schema,
	message protoreflect.Message,
) error {
	if len(bigQueryFields) != len(schema) {
		return fmt.Errorf("message has %d fields but schema has %d fields", len(bigQueryFields), len(schema))
	}
	for i, fieldSchema := range schema {
		bigQueryFieldValue := bigQueryFields[i]
		fieldName := protoreflect.Name(fieldSchema.Name)
		field := message.Descriptor().Fields().ByName(fieldName)
		if field == nil {
			if !o.DiscardUnknown && !message.Descriptor().ReservedNames().Has(fieldName) {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			continue
		}
		switch {
		case field.IsList():
			return fmt.Errorf("TODO: implement support for lists")
		case field.IsMap():
			return fmt.Errorf("TODO: implement support for maps")
		default:
			if err := o.loadSingularField(bigQueryFieldValue, fieldSchema, field, message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o UnmarshalOptions) unmarshalMessage(
	bigQueryFields map[string]bigquery.Value,
	message protoreflect.Message,
) error {
	for bigQueryFieldName, bigQueryFieldValue := range bigQueryFields {
		fieldName := protoreflect.Name(bigQueryFieldName)
		field := message.Descriptor().Fields().ByName(fieldName)
		if field == nil {
			if !o.DiscardUnknown && !message.Descriptor().ReservedNames().Has(fieldName) {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			continue
		}
		switch {
		case field.IsList():
			return fmt.Errorf("TODO: implement support for lists")
		case field.IsMap():
			return fmt.Errorf("TODO: implement support for maps")
		default:
			if err := o.unmarshalSingularField(bigQueryFieldValue, field, message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o UnmarshalOptions) loadSingularField(
	bigqueryValue bigquery.Value,
	fieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	if bigqueryValue == nil {
		return nil
	}
	if field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind {
		if fieldSchema.Type != bigquery.RecordFieldType {
			return fmt.Errorf("unsupported BigQuery type for message: %v", fieldSchema.Type)
		}
		bigqueryMessageValue, ok := bigqueryValue.([]bigquery.Value)
		if !ok {
			return fmt.Errorf("unsupported BigQuery value for message: %v", bigqueryMessageValue)
		}
		fieldValue := message.NewField(field)
		if err := o.loadMessage(bigqueryMessageValue, fieldSchema.Schema, fieldValue.Message()); err != nil {
			return err
		}
		message.Set(field, fieldValue)
	} else {
		fieldValue, err := o.unmarshalScalar(bigqueryValue, field)
		if err != nil {
			return err
		}
		message.Set(field, fieldValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalSingularField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	if bigqueryValue == nil {
		return nil
	}
	if field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind {
		bigqueryMessageValue, ok := bigqueryValue.(map[string]bigquery.Value)
		if !ok {
			return fmt.Errorf("unsupported BigQuery value for message: %v", bigqueryMessageValue)
		}
		fieldValue := message.NewField(field)
		if err := o.unmarshalMessage(bigqueryMessageValue, fieldValue.Message()); err != nil {
			return err
		}
		message.Set(field, fieldValue)
	} else {
		fieldValue, err := o.unmarshalScalar(bigqueryValue, field)
		if err != nil {
			return err
		}
		message.Set(field, fieldValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalScalar(
	bigQueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
) (protoreflect.Value, error) {
	switch field.Kind() {
	case protoreflect.BoolKind:
		if b, ok := bigQueryValue.(bool); ok {
			return protoreflect.ValueOfBool(b), nil
		}

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if n, ok := bigQueryValue.(int64); ok {
			return protoreflect.ValueOfInt32(int32(n)), nil
		}

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if n, ok := bigQueryValue.(int64); ok {
			return protoreflect.ValueOfInt64(n), nil
		}

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if n, ok := bigQueryValue.(int64); ok {
			return protoreflect.ValueOfUint32(uint32(n)), nil
		}

	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if n, ok := bigQueryValue.(int64); ok {
			return protoreflect.ValueOfUint64(uint64(n)), nil
		}

	case protoreflect.FloatKind:
		if n, ok := bigQueryValue.(float64); ok {
			return protoreflect.ValueOfFloat32(float32(n)), nil
		}

	case protoreflect.DoubleKind:
		if n, ok := bigQueryValue.(float64); ok {
			return protoreflect.ValueOfFloat64(n), nil
		}

	case protoreflect.StringKind:
		if s, ok := bigQueryValue.(string); ok {
			return protoreflect.ValueOfString(s), nil
		}

	case protoreflect.BytesKind:
		if b, ok := bigQueryValue.([]byte); ok {
			return protoreflect.ValueOfBytes(b), nil
		}

	case protoreflect.EnumKind:
		switch v := bigQueryValue.(type) {
		case string:
			if enumVal := field.Enum().Values().ByName(protoreflect.Name(v)); enumVal != nil {
				return protoreflect.ValueOfEnum(enumVal.Number()), nil
			}
		case int64:
			return protoreflect.ValueOfEnum(protoreflect.EnumNumber(int32(v))), nil
		}
	case protoreflect.MessageKind, protoreflect.GroupKind:
		// Fall through to return error, these should have been handled by the caller.
	}
	return protoreflect.Value{}, fmt.Errorf("invalid BigQuery value %v for kind %v", bigQueryValue, field.Kind())
}
