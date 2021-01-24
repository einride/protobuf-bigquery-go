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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Unmarshal the bigquery.Value map into the given proto.Message.
// It will clear the message first before setting the fields. If it returns an error,
// the given message may be partially set.
func Unmarshal(bqMessage map[string]bigquery.Value, message proto.Message) error {
	return UnmarshalOptions{}.Unmarshal(bqMessage, message)
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
func (o UnmarshalOptions) Unmarshal(bqMessage map[string]bigquery.Value, message proto.Message) error {
	proto.Reset(message)
	if err := o.unmarshalMessage(bqMessage, message.ProtoReflect()); err != nil {
		return err
	}
	if o.AllowPartial {
		return nil
	}
	return proto.CheckInitialized(message)
}

func (o UnmarshalOptions) unmarshalMessage(
	bqMessage map[string]bigquery.Value,
	message protoreflect.Message,
) error {
	for bqFieldName, bqField := range bqMessage {
		fieldName := protoreflect.Name(bqFieldName)
		field := message.Descriptor().Fields().ByName(fieldName)
		if field == nil {
			if !o.DiscardUnknown && !message.Descriptor().ReservedNames().Has(fieldName) {
				return fmt.Errorf("unknown field: %s", fieldName)
			}
			continue
		}
		switch {
		case field.IsList():
			if err := o.unmarshalListField(bqField, field, message); err != nil {
				return err
			}
		case field.IsMap():
			if err := o.unmarshalMapField(bqField, field, message); err != nil {
				return err
			}
		default:
			value, err := o.unmarshalSingularField(bqField, field, message)
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

func (o UnmarshalOptions) unmarshalListField(
	bqValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	bqListValue, ok := bqValue.([]bigquery.Value)
	if !ok {
		return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bqValue)
	}
	isMessage := field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind
	switch {
	case isMessage && wkt.IsWellKnownType(string(field.Message().FullName())):
		return o.unmarshalWellKnownTypeListField(bqListValue, field, message)
	case isMessage:
		return o.unmarshalMessageListField(bqListValue, field, message)
	default:
		return o.unmarshalScalarListField(bqListValue, field, message)
	}
}

func (o UnmarshalOptions) unmarshalWellKnownTypeListField(
	bqListValue []bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	list := message.Mutable(field).List()
	for _, bqListElementValue := range bqListValue {
		value, err := o.unmarshalWellKnownTypeField(bqListElementValue, field)
		if err != nil {
			return err
		}
		list.Append(value)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalMessageListField(
	bqListValue []bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	list := message.Mutable(field).List()
	for _, bqListElementValue := range bqListValue {
		bqListElementMessageValue, ok := bqListElementValue.(map[string]bigquery.Value)
		if !ok {
			return fmt.Errorf(
				"%s: unsupported BigQuery value for message: %v", field.Name(), bqListElementMessageValue,
			)
		}
		listElementValue := list.NewElement()
		if err := o.unmarshalMessage(bqListElementMessageValue, listElementValue.Message()); err != nil {
			return err
		}
		list.Append(listElementValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalScalarListField(
	bqListValue []bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) error {
	list := message.Mutable(field).List()
	for _, bqListElementValue := range bqListValue {
		value, err := o.unmarshalScalar(bqListElementValue, field)
		if err != nil {
			return err
		}
		list.Append(value)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalMapField(
	bqField bigquery.Value,
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
		return o.unmarshalMessageValueMapField(bqMapField, field, message)
	default:
		return o.unmarshalScalarValueMapField(bqMapField, field, message)
	}
}

func (o UnmarshalOptions) unmarshalScalarValueMapField(
	bqMapField []bigquery.Value,
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
		mapEntryValue, err := o.unmarshalScalar(bqMapEntryValue, field.MapValue())
		if err != nil {
			return err
		}
		mapField.Set(mapEntryKey, mapEntryValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalWellKnownTypeValueMapField(
	bqMapField []bigquery.Value,
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
		mapEntryValue, err := o.unmarshalWellKnownTypeField(bqMapEntryValue, field.MapValue())
		if err != nil {
			return err
		}
		mapField.Set(mapEntryKey, mapEntryValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalMessageValueMapField(
	bqMapField []bigquery.Value,
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
		bqMapEntryMessageValue, ok := bqMapEntryValue.(map[string]bigquery.Value)
		if !ok {
			return fmt.Errorf("%s: unsupported BigQuery value for message: %v", field.Name(), bqMapEntryValue)
		}
		mapEntryValue := mapField.NewValue()
		if err := o.unmarshalMessage(bqMapEntryMessageValue, mapEntryValue.Message()); err != nil {
			return err
		}
		mapField.Set(mapEntryKey, mapEntryValue)
	}
	return nil
}

func (o UnmarshalOptions) unmarshalMapEntryKey(
	bqMapEntry map[string]bigquery.Value,
) (protoreflect.MapKey, error) {
	bqMapEntryKey, ok := bqMapEntry["key"]
	if !ok {
		return protoreflect.MapKey{}, fmt.Errorf("map entry is missing key field")
	}
	return protoreflect.ValueOf(bqMapEntryKey).MapKey(), nil
}

func (o UnmarshalOptions) unmarshalSingularField(
	bqValue bigquery.Value,
	field protoreflect.FieldDescriptor,
	message protoreflect.Message,
) (protoreflect.Value, error) {
	if bqValue == nil {
		return protoreflect.ValueOf(nil), nil
	}
	if field.Kind() == protoreflect.MessageKind || field.Kind() == protoreflect.GroupKind {
		if wkt.IsWellKnownType(string(field.Message().FullName())) {
			return o.unmarshalWellKnownTypeField(bqValue, field)
		}
		bqMessage, ok := bqValue.(map[string]bigquery.Value)
		if !ok {
			return protoreflect.ValueOf(nil), fmt.Errorf(
				"%s: unsupported BigQuery value for message: %v", field.Name(), bqMessage,
			)
		}
		fieldValue := message.NewField(field)
		if err := o.unmarshalMessage(bqMessage, fieldValue.Message()); err != nil {
			return protoreflect.ValueOf(nil), fmt.Errorf("%s: %w", field.Name(), err)
		}
		return fieldValue, nil
	}
	return o.unmarshalScalar(bqValue, field)
}

func (o UnmarshalOptions) unmarshalWellKnownTypeField(
	bqValue bigquery.Value,
	field protoreflect.FieldDescriptor,
) (protoreflect.Value, error) {
	var result proto.Message
	var err error
	switch field.Message().FullName() {
	case wkt.Timestamp:
		result, err = o.unmarshalTimestamp(bqValue)
	case wkt.Duration:
		result, err = o.unmarshalDuration(bqValue)
	case wkt.TimeOfDay:
		result, err = o.unmarshalTimeOfDay(bqValue)
	case wkt.Date:
		result, err = o.unmarshalDate(bqValue)
	case wkt.LatLng:
		result, err = o.unmarshalLatLng(bqValue)
	case wkt.Struct:
		result, err = o.unmarshalStruct(bqValue)
	case wkt.DoubleValue:
		result, err = o.unmarshalDoubleValue(bqValue)
	case wkt.FloatValue:
		result, err = o.unmarshalFloatValue(bqValue)
	case wkt.Int32Value:
		result, err = o.unmarshalInt32Value(bqValue)
	case wkt.Int64Value:
		result, err = o.unmarshalInt64Value(bqValue)
	case wkt.UInt32Value:
		result, err = o.unmarshalUInt32Value(bqValue)
	case wkt.UInt64Value:
		result, err = o.unmarshalUInt64Value(bqValue)
	case wkt.BoolValue:
		result, err = o.unmarshalBoolValue(bqValue)
	case wkt.StringValue:
		result, err = o.unmarshalStringValue(bqValue)
	case wkt.BytesValue:
		result, err = o.unmarshalBytesValue(bqValue)
	default:
		result, err = nil, fmt.Errorf("unsupported well-known-type: %s", field.Message().FullName())
	}
	if err != nil {
		return protoreflect.ValueOf(nil), err
	}
	return protoreflect.ValueOf(result.ProtoReflect()), nil
}

func (o UnmarshalOptions) unmarshalTimestamp(bqValue bigquery.Value) (*timestamppb.Timestamp, error) {
	t, ok := bqValue.(time.Time)
	if !ok {
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %v", wkt.Timestamp, bqValue)
	}
	return timestamppb.New(t), nil
}

func (o UnmarshalOptions) unmarshalDuration(bqValue bigquery.Value) (*durationpb.Duration, error) {
	var duration time.Duration
	switch bigqueryValue := bqValue.(type) {
	case int64:
		duration = time.Duration(bigqueryValue) * time.Second
	case float64:
		duration = time.Duration(bigqueryValue * float64(time.Second))
	default:
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Duration, bigqueryValue)
	}
	return durationpb.New(duration), nil
}

func (o UnmarshalOptions) unmarshalTimeOfDay(bqValue bigquery.Value) (*timeofday.TimeOfDay, error) {
	t, ok := bqValue.(civil.Time)
	if !ok {
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.TimeOfDay, bqValue)
	}
	return &timeofday.TimeOfDay{
		Hours:   int32(t.Hour),
		Minutes: int32(t.Minute),
		Seconds: int32(t.Second),
		Nanos:   int32(t.Nanosecond),
	}, nil
}

func (o UnmarshalOptions) unmarshalDate(bqValue bigquery.Value) (*date.Date, error) {
	d, ok := bqValue.(civil.Date)
	if !ok {
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Date, bqValue)
	}
	return &date.Date{
		Year:  int32(d.Year),
		Month: int32(d.Month),
		Day:   int32(d.Day),
	}, nil
}

func (o UnmarshalOptions) unmarshalLatLng(bqValue bigquery.Value) (*latlng.LatLng, error) {
	s, ok := bqValue.(string)
	if !ok {
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.LatLng, bqValue)
	}
	latLng := &latlng.LatLng{}
	if _, err := fmt.Sscanf(s, "POINT(%f %f)", &latLng.Longitude, &latLng.Latitude); err != nil {
		return nil, fmt.Errorf("invalid GEOGRAPHY value for %s: %#v: %w", wkt.LatLng, bqValue, err)
	}
	return latLng, nil
}

func (o UnmarshalOptions) unmarshalStruct(bqValue bigquery.Value) (*structpb.Struct, error) {
	s, ok := bqValue.(string)
	if !ok {
		return nil, fmt.Errorf("unsupported BigQuery value for %s: %#v", wkt.Struct, bqValue)
	}
	var structValue structpb.Struct
	if err := structValue.UnmarshalJSON([]byte(s)); err != nil {
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v: %w", wkt.Struct, bqValue, err)
	}
	return &structValue, nil
}

func (o UnmarshalOptions) unmarshalDoubleValue(bqValue bigquery.Value) (*wrapperspb.DoubleValue, error) {
	switch bqValue := bqValue.(type) {
	case float32:
		return wrapperspb.Double(float64(bqValue)), nil
	case float64:
		return wrapperspb.Double(bqValue), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.DoubleValue, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalFloatValue(bqValue bigquery.Value) (*wrapperspb.FloatValue, error) {
	switch bqValue := bqValue.(type) {
	case float32:
		return wrapperspb.Float(bqValue), nil
	case float64:
		return wrapperspb.Float(float32(bqValue)), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.FloatValue, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalInt32Value(bqValue bigquery.Value) (*wrapperspb.Int32Value, error) {
	switch bqValue := bqValue.(type) {
	case int32:
		return wrapperspb.Int32(bqValue), nil
	case int64:
		return wrapperspb.Int32(int32(bqValue)), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.Int32Value, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalInt64Value(bqValue bigquery.Value) (*wrapperspb.Int64Value, error) {
	switch bqValue := bqValue.(type) {
	case int32:
		return wrapperspb.Int64(int64(bqValue)), nil
	case int64:
		return wrapperspb.Int64(bqValue), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.Int64Value, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalUInt32Value(bqValue bigquery.Value) (*wrapperspb.UInt32Value, error) {
	switch bqValue := bqValue.(type) {
	case uint32:
		return wrapperspb.UInt32(bqValue), nil
	case uint64:
		return wrapperspb.UInt32(uint32(bqValue)), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.UInt32Value, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalUInt64Value(bqValue bigquery.Value) (*wrapperspb.UInt64Value, error) {
	switch bqValue := bqValue.(type) {
	case uint32:
		return wrapperspb.UInt64(uint64(bqValue)), nil
	case uint64:
		return wrapperspb.UInt64(bqValue), nil
	default:
		return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.UInt64Value, bqValue)
	}
}

func (o UnmarshalOptions) unmarshalBoolValue(bqValue bigquery.Value) (*wrapperspb.BoolValue, error) {
	if bqValue, ok := bqValue.(bool); ok {
		return wrapperspb.Bool(bqValue), nil
	}
	return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.BoolValue, bqValue)
}

func (o UnmarshalOptions) unmarshalStringValue(bqValue bigquery.Value) (*wrapperspb.StringValue, error) {
	if bqValue, ok := bqValue.(string); ok {
		return wrapperspb.String(bqValue), nil
	}
	return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.StringValue, bqValue)
}

func (o UnmarshalOptions) unmarshalBytesValue(bqValue bigquery.Value) (*wrapperspb.BytesValue, error) {
	if bqValue, ok := bqValue.([]byte); ok {
		return wrapperspb.Bytes(bqValue), nil
	}
	return nil, fmt.Errorf("invalid BigQuery value for %s: %#v", wkt.BytesValue, bqValue)
}

func (o UnmarshalOptions) unmarshalScalar(
	bqValue bigquery.Value,
	field protoreflect.FieldDescriptor,
) (protoreflect.Value, error) {
	switch field.Kind() {
	case protoreflect.BoolKind:
		if b, ok := bqValue.(bool); ok {
			return protoreflect.ValueOfBool(b), nil
		}

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		if n, ok := bqValue.(int64); ok {
			return protoreflect.ValueOfInt32(int32(n)), nil
		}

	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		if n, ok := bqValue.(int64); ok {
			return protoreflect.ValueOfInt64(n), nil
		}

	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		if n, ok := bqValue.(int64); ok {
			return protoreflect.ValueOfUint32(uint32(n)), nil
		}

	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		if n, ok := bqValue.(int64); ok {
			return protoreflect.ValueOfUint64(uint64(n)), nil
		}

	case protoreflect.FloatKind:
		if n, ok := bqValue.(float64); ok {
			return protoreflect.ValueOfFloat32(float32(n)), nil
		}

	case protoreflect.DoubleKind:
		if n, ok := bqValue.(float64); ok {
			return protoreflect.ValueOfFloat64(n), nil
		}

	case protoreflect.StringKind:
		if s, ok := bqValue.(string); ok {
			return protoreflect.ValueOfString(s), nil
		}

	case protoreflect.BytesKind:
		if b, ok := bqValue.([]byte); ok {
			return protoreflect.ValueOfBytes(b), nil
		}

	case protoreflect.EnumKind:
		return o.unmarshalEnumScalar(bqValue, field)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		// Fall through to return error, these should have been handled by the caller.
	}
	return protoreflect.Value{}, fmt.Errorf("invalid BigQuery value %#v for kind %v", bqValue, field.Kind())
}

func (o UnmarshalOptions) unmarshalEnumScalar(
	bqValue bigquery.Value,
	field protoreflect.FieldDescriptor,
) (protoreflect.Value, error) {
	if o.Schema.UseEnumNumbers {
		v, ok := bqValue.(int64)
		if !ok {
			return protoreflect.Value{}, fmt.Errorf(
				"invalid BigQuery value %#v for enum %s number", bqValue, field.Enum().FullName(),
			)
		}
		return protoreflect.ValueOfEnum(protoreflect.EnumNumber(int32(v))), nil
	}
	v, ok := bqValue.(string)
	if !ok {
		return protoreflect.Value{}, fmt.Errorf(
			"invalid BigQuery value %#v for enum %s", bqValue, field.Enum().FullName(),
		)
	}
	enumVal := field.Enum().Values().ByName(protoreflect.Name(v))
	if enumVal == nil {
		return protoreflect.Value{}, fmt.Errorf(
			"unknown enum value %#v for enum %s", bqValue, field.Enum().FullName(),
		)
	}
	return protoreflect.ValueOfEnum(enumVal.Number()), nil
}
