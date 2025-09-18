package protobq

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/goalsgame/protobuf-bigquery/internal/wkt"
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
			if err := o.loadListField(bqField, bqFieldSchema, field, message); err != nil {
				return err
			}
		case field.IsMap():
			if err := o.loadMapField(bqField, bqFieldSchema, field, message); err != nil {
				return err
			}
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
	if !bqFieldSchema.Repeated {
		return fmt.Errorf("%s: unsupported field schema for list field: not repeated", field.Name())
	}
	bqList, ok := bqField.([]bigquery.Value)
	if !ok {
		return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bqField)
	}
	isMessage := field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind
	switch {
	case isMessage && wkt.IsWellKnownType(string(field.Message().FullName())):
		return o.unmarshalWellKnownTypeListField(bqList, field, message)
	case isMessage:
		return o.loadMessageListField(bqList, bqFieldSchema, field, message)
	default:
		return o.unmarshalScalarListField(bqList, field, message)
	}
}

func (o UnmarshalOptions) loadMessageListField(
	bqListValue []bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	list := message.Mutable(field).List()
	for _, bqElement := range bqListValue {
		if bqFieldSchema.Type != bigquery.RecordFieldType {
			return fmt.Errorf(
				"%s: field schema has type %s but expected %s",
				field.Name(),
				bqFieldSchema.Type,
				bigquery.RecordFieldType,
			)
		}
		bqMessageElement, ok := bqElement.([]bigquery.Value)
		if !ok {
			return fmt.Errorf(
				"%s: unsupported BigQuery value for message: %v", field.Name(), bqMessageElement,
			)
		}
		listElementValue := list.NewElement()
		if err := o.loadMessage(bqMessageElement, bqFieldSchema.Schema, listElementValue.Message()); err != nil {
			return err
		}
		list.Append(listElementValue)
	}
	return nil
}

func (o UnmarshalOptions) loadMapField(
	bqField bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	bqMapField, ok := bqField.([]bigquery.Value)
	if !ok {
		return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bqField)
	}
	mapValue := field.MapValue()
	isMessage := mapValue.Kind() == protoreflect.MessageKind || mapValue.Kind() == protoreflect.GroupKind
	switch {
	case isMessage && wkt.IsWellKnownType(string(mapValue.Message().FullName())):
		return o.unmarshalWellKnownTypeValueMapField(bqMapField, field, message)
	case isMessage:
		return o.loadMessageValueMapField(bqMapField, bqFieldSchema, field, message)
	default:
		return o.unmarshalScalarValueMapField(bqMapField, field, message)
	}
}

func (o UnmarshalOptions) loadMessageValueMapField(
	bqMapField []bigquery.Value,
	bqFieldSchema *bigquery.FieldSchema,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	mapField := message.Mutable(field).Map()
	for _, bqMapEntry := range bqMapField {
		bqMapEntry, ok := bqMapEntry.(map[string]bigquery.Value)
		if !ok {
			return fmt.Errorf("%s: unsupported BigQuery value for map entry: %v", field.Name(), bqMapEntry)
		}
		mapEntryKey, err := o.unmarshalMapEntryKey(bqMapEntry)
		if err != nil {
			return err
		}
		bqMapEntryValue, ok := bqMapEntry["value"]
		if !ok {
			return fmt.Errorf("%s: map entry is missing value field", field.Name())
		}
		bqMapEntryMessageValue, ok := bqMapEntryValue.([]bigquery.Value)
		if !ok {
			return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bqMapEntryValue)
		}
		if len(bqFieldSchema.Schema) != 2 || bqFieldSchema.Schema[1].Name != "value" {
			return fmt.Errorf("%s: unsupported BigQuery schema for map entry", field.Name())
		}
		bqMapEntryValueSchema := bqFieldSchema.Schema[1].Schema
		mapEntryValue := mapField.NewValue()
		if err := o.loadMessage(bqMapEntryMessageValue, bqMapEntryValueSchema, mapEntryValue.Message()); err != nil {
			return err
		}
		mapField.Set(mapEntryKey, mapEntryValue)
	}
	return nil
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
