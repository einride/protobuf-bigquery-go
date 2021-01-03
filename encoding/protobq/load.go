package protobq

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Load the bigquery.Value list into the given proto.Message using the given bigquery.Schema.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func Load(row []bigquery.Value, schema bigquery.Schema, message proto.Message) error {
	return UnmarshalOptions{}.Load(row, schema, message)
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
		if wkt.IsWellKnownType(string(field.Message().FullName())) {
			if err := o.unmarshalWellKnownTypeField(bigqueryValue, field, message); err != nil {
				return fmt.Errorf("%s: %w", field.Name(), err)
			}
			return nil
		}
		if fieldSchema.Type != bigquery.RecordFieldType {
			return fmt.Errorf("%s: unsupported BigQuery type for message: %v", field.Name(), fieldSchema.Type)
		}
		bigqueryMessageValue, ok := bigqueryValue.([]bigquery.Value)
		if !ok {
			return fmt.Errorf("unsupported BigQuery value for message: %v", bigqueryMessageValue)
		}
		fieldValue := message.NewField(field)
		if err := o.loadMessage(bigqueryMessageValue, fieldSchema.Schema, fieldValue.Message()); err != nil {
			return fmt.Errorf("%s: %w", field.Name(), err)
		}
		message.Set(field, fieldValue)
	} else {
		fieldValue, err := o.unmarshalScalar(bigqueryValue, field)
		if err != nil {
			return fmt.Errorf("%s: %w", field.Name(), err)
		}
		message.Set(field, fieldValue)
	}
	return nil
}
