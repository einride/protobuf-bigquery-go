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
func Load(bqMessage []bigquery.Value, bqSchema bigquery.Schema, message proto.Message) error {
	return UnmarshalOptions{}.Load(bqMessage, bqSchema, message)
}

// Load the bigquery.Value list into the given proto.Message using the given bigquery.Schema
// using options in UnmarshalOptions object.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func (o UnmarshalOptions) Load(bqMessage []bigquery.Value, bqSchema bigquery.Schema, message proto.Message) error {
	proto.Reset(message)
	if err := o.loadMessage(bqMessage, bqSchema, message.ProtoReflect()); err != nil {
		return err
	}
	if o.AllowPartial {
		return nil
	}
	return proto.CheckInitialized(message)
}

func (o UnmarshalOptions) loadMessage(
	bqMessage []bigquery.Value,
	bqSchema bigquery.Schema,
	message protoreflect.Message,
) error {
	if len(bqMessage) != len(bqSchema) {
		return fmt.Errorf("message has %d fields but schema has %d fields", len(bqMessage), len(bqSchema))
	}
	for i, bqFieldSchema := range bqSchema {
		bqField := bqMessage[i]
		fieldName := protoreflect.Name(bqFieldSchema.Name)
		field := message.Descriptor().Fields().ByName(fieldName)
		if field == nil {
			if !o.DiscardUnknown && !message.Descriptor().ReservedNames().Has(fieldName) {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			continue
		}
		switch {
		case field.IsList():
			return o.loadListField(bqField, bqFieldSchema, field, message)
		case field.IsMap():
			return o.loadMapField(bqField, bqFieldSchema, field, message)
		default:
			value, err := o.loadSingularField(bqField, bqFieldSchema, field, message)
			if err != nil {
				return err
			}
			if value.IsValid() {
				message.Set(field, value)
			}
		}
	}
	return nil
}

func (o UnmarshalOptions) loadListField(
	bqField bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	return fmt.Errorf("TODO: implement support for lists")
}

func (o UnmarshalOptions) loadMapField(
	bqField bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	return fmt.Errorf("TODO: implement support for maps")
}

func (o UnmarshalOptions) loadSingularField(
	bqField bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) (protoreflect.Value, error) {
	if bqField == nil {
		return protoreflect.ValueOf(nil), nil
	}
	if field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind {
		if wkt.IsWellKnownType(string(field.Message().FullName())) {
			return o.unmarshalWellKnownTypeField(bqField, field)
		}
		if bqFieldSchema.Type != bigquery.RecordFieldType {
			return protoreflect.ValueOf(nil), fmt.Errorf(
				"%s: unsupported BigQuery type for message: %v", field.Name(), bqFieldSchema.Type,
			)
		}
		bqMessage, ok := bqField.([]bigquery.Value)
		if !ok {
			return protoreflect.ValueOf(nil), fmt.Errorf("unsupported BigQuery value for message: %v", bqMessage)
		}
		fieldValue := message.NewField(field)
		if err := o.loadMessage(bqMessage, bqFieldSchema.Schema, fieldValue.Message()); err != nil {
			return protoreflect.ValueOf(nil), fmt.Errorf("%s: %w", field.Name(), err)
		}
		return fieldValue, nil
	}
	return o.unmarshalScalar(bqField, field)
}
