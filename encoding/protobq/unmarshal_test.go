package protobq

import (
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/example/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestUnmarshalOptions_Unmarshal(t *testing.T) {
	for _, tt := range []struct {
		name          string
		row           map[string]bigquery.Value
		opt           UnmarshalOptions
		expected      proto.Message
		errorContains string
	}{
		{
			name: "library.Book",
			row: map[string]bigquery.Value{
				"name":   "name",
				"author": "author",
				"title":  "title",
				"read":   true,
			},
			expected: &library.Book{
				Name:   "name",
				Author: "author",
				Title:  "title",
				Read:   true,
			},
		},

		{
			name: "library.UpdateBookRequest",
			row: map[string]bigquery.Value{
				"book": map[string]bigquery.Value{
					"name":   "name",
					"author": "author",
					"title":  "title",
					"read":   true,
				},
			},
			expected: &library.UpdateBookRequest{
				Book: &library.Book{
					Name:   "name",
					Author: "author",
					Title:  "title",
					Read:   true,
				},
			},
		},

		{
			name: "expr.Value (bool)",
			row: map[string]bigquery.Value{
				"bool_value": true,
			},
			expected: &expr.Value{
				Kind: &expr.Value_BoolValue{
					BoolValue: true,
				},
			},
		},

		{
			name: "expr.Value (bool)",
			row: map[string]bigquery.Value{
				"double_value": float64(42),
			},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "error on unknown fields",
			row: map[string]bigquery.Value{
				"foo": "bar",
			},
			expected:      &expr.Value{},
			errorContains: "unknown field: foo",
		},

		{
			name: "discard unknown fields",
			row: map[string]bigquery.Value{
				"foo":          "bar",
				"double_value": float64(42),
			},
			opt: UnmarshalOptions{DiscardUnknown: true},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "enum values",
			row: map[string]bigquery.Value{
				"enum_value": "ENUM_VALUE1",
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "enum numbers",
			row: map[string]bigquery.Value{
				"enum_value": int64(1),
			},
			opt: UnmarshalOptions{
				Schema: SchemaOptions{
					UseEnumNumbers: true,
				},
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "wrappers",
			row: map[string]bigquery.Value{
				"float_value":  float64(1),
				"double_value": float64(2),
				"string_value": "foo",
				"bytes_value":  []byte("bar"),
				"int32_value":  int64(3),
				"int64_value":  int64(4),
				"uint32_value": uint64(5),
				"uint64_value": uint64(6),
				"bool_value":   true,
			},
			expected: &examplev1.ExampleWrappers{
				FloatValue:  wrapperspb.Float(1),
				DoubleValue: wrapperspb.Double(2),
				StringValue: wrapperspb.String("foo"),
				BytesValue:  wrapperspb.Bytes([]byte("bar")),
				Int32Value:  wrapperspb.Int32(3),
				Int64Value:  wrapperspb.Int64(4),
				Uint32Value: wrapperspb.UInt32(5),
				Uint64Value: wrapperspb.UInt64(6),
				BoolValue:   wrapperspb.Bool(true),
			},
		},

		{
			name: "primitive lists",
			row: map[string]bigquery.Value{
				"int64_list":  []bigquery.Value{int64(1), int64(2)},
				"string_list": []bigquery.Value{"a", "b"},
				"enum_list":   []bigquery.Value{"ENUM_VALUE1", "ENUM_VALUE2"},
			},
			expected: &examplev1.ExampleList{
				Int64List:  []int64{1, 2},
				StringList: []string{"a", "b"},
				EnumList: []examplev1.ExampleList_Enum{
					examplev1.ExampleList_ENUM_VALUE1,
					examplev1.ExampleList_ENUM_VALUE2,
				},
			},
		},

		{
			name: "well-known-type lists",
			row: map[string]bigquery.Value{
				"float_value_list": []bigquery.Value{float32(1), float32(2)},
			},
			expected: &examplev1.ExampleList{
				FloatValueList: []*wrapperspb.FloatValue{
					wrapperspb.Float(1), wrapperspb.Float(2),
				},
			},
		},

		{
			name: "lists",
			row: map[string]bigquery.Value{
				"int64_list":  []bigquery.Value{int64(1), int64(2)},
				"string_list": []bigquery.Value{"a", "b"},
				"enum_list":   []bigquery.Value{"ENUM_VALUE1", "ENUM_VALUE2"},
				"nested_list": []bigquery.Value{
					map[string]bigquery.Value{
						"string_list": []bigquery.Value{"a", "b"},
					},
					map[string]bigquery.Value{
						"string_list": []bigquery.Value{"c", "d"},
					},
				},
			},
			expected: &examplev1.ExampleList{
				Int64List:  []int64{1, 2},
				StringList: []string{"a", "b"},
				EnumList: []examplev1.ExampleList_Enum{
					examplev1.ExampleList_ENUM_VALUE1,
					examplev1.ExampleList_ENUM_VALUE2,
				},
				NestedList: []*examplev1.ExampleList_Nested{
					{StringList: []string{"a", "b"}},
					{StringList: []string{"c", "d"}},
				},
			},
		},

		{
			name: "primitive maps",
			row: map[string]bigquery.Value{
				"string_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": "b"},
				},
				"string_to_enum": []bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": "ENUM_VALUE1"},
				},
				"int32_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": int64(1), "value": "a"},
				},
				"int64_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": int64(2), "value": "a"},
				},
				"uint32_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": uint64(3), "value": "a"},
				},
				"bool_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": true, "value": "a"},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToString: map[string]string{"a": "b"},
				StringToEnum:   map[string]examplev1.ExampleMap_Enum{"a": examplev1.ExampleMap_ENUM_VALUE1},
				Int32ToString:  map[int32]string{1: "a"},
				Int64ToString:  map[int64]string{2: "a"},
				Uint32ToString: map[uint32]string{3: "a"},
				BoolToString:   map[bool]string{true: "a"},
			},
		},

		{
			name: "well-known-type maps",
			row: map[string]bigquery.Value{
				"string_to_float_value": []bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": float64(1)},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToFloatValue: map[string]*wrapperspb.FloatValue{
					"a": wrapperspb.Float(1),
				},
			},
		},

		{
			name: "nested maps",
			row: map[string]bigquery.Value{
				"string_to_nested": []bigquery.Value{
					map[string]bigquery.Value{
						"key": "a",
						"value": map[string]bigquery.Value{
							"string_to_string": []bigquery.Value{
								map[string]bigquery.Value{"key": "a", "value": "b"},
							},
						},
					},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToNested: map[string]*examplev1.ExampleMap_Nested{
					"a": {
						StringToString: map[string]string{
							"a": "b",
						},
					},
				},
			},
		},

		{
			name: "datetime (without offset)",
			row: map[string]bigquery.Value{
				"date_time": civil.DateTime{
					Date: civil.Date{
						Year:  2021,
						Month: time.February,
						Day:   1,
					},
					Time: civil.Time{
						Hour:       8,
						Minute:     30,
						Second:     1,
						Nanosecond: 2,
					},
				},
			},
			opt: UnmarshalOptions{
				Schema: SchemaOptions{
					UseDateTimeWithoutOffset: true,
				},
			},
			expected: &examplev1.ExampleDateTime{
				DateTime: &datetime.DateTime{
					Year:    2021,
					Month:   int32(time.February),
					Day:     1,
					Hours:   8,
					Minutes: 30,
					Seconds: 1,
					Nanos:   2,
				},
			},
		},

		{
			name: "oneof empty message, incorrect input",
			row: map[string]bigquery.Value{
				"oneof_fields_1": 2,
			},
			opt:           UnmarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			errorContains: "unexpected type int for oneof field oneof_fields_1, expected string",
			expected:      &examplev1.ExampleOneof{},
		},

		{
			name: "oneof empty messages",
			row: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_empty_message_1",
				"oneof_fields_2": "oneof_empty_message_2",
			},
			opt: UnmarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofEmptyMessage_1{
					OneofEmptyMessage_1: &examplev1.ExampleOneof_EmptyMessage{},
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofEmptyMessage_2{
					OneofEmptyMessage_2: &examplev1.ExampleOneof_EmptyMessage{},
				},
			},
		},

		{
			name: "one oneof empty message, one non-empty",
			row: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_empty_message_1",
				"oneof_fields_2": "oneof_message",
				"oneof_message": map[string]bigquery.Value{
					"string_value": "value",
				},
			},
			opt: UnmarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofEmptyMessage_1{
					OneofEmptyMessage_1: &examplev1.ExampleOneof_EmptyMessage{},
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofMessage{
					OneofMessage: &examplev1.ExampleOneof_Message{
						StringValue: "value",
					},
				},
			},
		},

		{
			name: "oneof non-empty message",
			row: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_bool_1",
				"oneof_bool_1":   true,
				"oneof_fields_2": "oneof_message",
				"oneof_message": map[string]bigquery.Value{
					"string_value": "value",
				},
			},
			opt: UnmarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofBool_1{
					OneofBool_1: true,
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofMessage{
					OneofMessage: &examplev1.ExampleOneof_Message{
						StringValue: "value",
					},
				},
			},
		},

		{
			name:     "oneof empty root message",
			row:      map[string]bigquery.Value{},
			opt:      UnmarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: &examplev1.ExampleOneof{},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := proto.Clone(tt.expected)
			proto.Reset(actual)
			if err := tt.opt.Unmarshal(tt.row, actual); tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tt.expected, actual, protocmp.Transform())
			}
		})
	}
}
