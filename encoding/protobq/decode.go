package protobq

import (
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"go.einride.tech/protobuf-bigquery/internal/wkt"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	// Schema contains the schema options.
	Schema SchemaOptions

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

func (o UnmarshalOptions) unmarshalSingularField(
	bigqueryValue bigquery.Value,
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
		bigqueryMessageValue, ok := bigqueryValue.(map[string]bigquery.Value)
		if !ok {
			return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bigqueryMessageValue)
		}
		fieldValue := message.NewField(field)
		if err := o.unmarshalMessage(bigqueryMessageValue, fieldValue.Message()); err != nil {
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

func (o UnmarshalOptions) unmarshalWellKnownTypeField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	switch field.Message().FullName() {
	case wkt.Timestamp:
		return o.unmarshalTimestampField(bigqueryValue, field, message)
	case wkt.Duration:
		return o.unmarshalDurationField(bigqueryValue, field, message)
	case wkt.TimeOfDay:
		return o.unmarshalTimeOfDayField(bigqueryValue, field, message)
	case wkt.Date:
		return o.unmarshalDateField(bigqueryValue, field, message)
	case wkt.LatLng:
		return o.unmarshalLatLngField(bigqueryValue, field, message)
	case wkt.Struct:
		return o.unmarshalStructField(bigqueryValue, field, message)
	default:
		return fmt.Errorf("unsupported well-known-type: %s", field.Message().FullName())
	}
}

func (o UnmarshalOptions) unmarshalTimestampField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	t, ok := bigqueryValue.(time.Time)
	if !ok {
		return fmt.Errorf("unsupported BigQuery value for %s: %v", wkt.Timestamp, bigqueryValue)
	}
	message.Set(field, protoreflect.ValueOfMessage(timestamppb.New(t).ProtoReflect()))
	return nil
}

func (o UnmarshalOptions) unmarshalDurationField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	var duration time.Duration
	switch bigqueryValue := bigqueryValue.(type) {
	case int64:
		duration = time.Duration(bigqueryValue) * time.Second
	case float64:
		duration = time.Duration(bigqueryValue * float64(time.Second))
	default:
		return fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Duration, bigqueryValue)
	}
	message.Set(field, protoreflect.ValueOfMessage(durationpb.New(duration).ProtoReflect()))
	return nil
}

func (o UnmarshalOptions) unmarshalTimeOfDayField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	t, ok := bigqueryValue.(civil.Time)
	if !ok {
		return fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.TimeOfDay, bigqueryValue)
	}
	message.Set(field, protoreflect.ValueOfMessage((&timeofday.TimeOfDay{
		Hours:   int32(t.Hour),
		Minutes: int32(t.Minute),
		Seconds: int32(t.Second),
		Nanos:   int32(t.Nanosecond),
	}).ProtoReflect()))
	return nil
}

func (o UnmarshalOptions) unmarshalDateField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	d, ok := bigqueryValue.(civil.Date)
	if !ok {
		return fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Date, bigqueryValue)
	}
	message.Set(field, protoreflect.ValueOfMessage((&date.Date{
		Year:  int32(d.Year),
		Month: int32(d.Month),
		Day:   int32(d.Day),
	}).ProtoReflect()))
	return nil
}

func (o UnmarshalOptions) unmarshalLatLngField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	s, ok := bigqueryValue.(string)
	if !ok {
		return fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.LatLng, bigqueryValue)
	}
	latLng := &latlng.LatLng{}
	if _, err := fmt.Sscanf(s, "POINT(%f %f)", &latLng.Longitude, &latLng.Latitude); err != nil {
		return fmt.Errorf("invalid GEOGRAPHY value for %s: %#v: %w", wkt.LatLng, bigqueryValue, err)
	}
	message.Set(field, protoreflect.ValueOfMessage(latLng.ProtoReflect()))
	return nil
}

func (o UnmarshalOptions) unmarshalStructField(
	bigqueryValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	s, ok := bigqueryValue.(string)
	if !ok {
		return fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Struct, bigqueryValue)
	}
	var structValue structpb.Struct
	if err := structValue.UnmarshalJSON([]byte(s)); err != nil {
		return fmt.Errorf("invalid BigQuery value for %s: %#v: %w", wkt.Struct, bigqueryValue, err)
	}
	message.Set(field, protoreflect.ValueOfMessage(structValue.ProtoReflect()))
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
	return protoreflect.Value{}, fmt.Errorf("invalid BigQuery value %#v for kind %v", bigQueryValue, field.Kind())
}
