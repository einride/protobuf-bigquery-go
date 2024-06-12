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
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestMarshalOptions_Marshal(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		opt      MarshalOptions
		expected map[string]bigquery.Value
	}{
		{
			name: "library.Book",
			msg: &library.Book{
				Name:   "name",
				Author: "author",
				Title:  "title",
				Read:   true,
			},
			expected: map[string]bigquery.Value{
				"name":   "name",
				"author": "author",
				"title":  "title",
				"read":   true,
			},
		},

		{
			name: "library.UpdateBookRequest",
			msg: &library.UpdateBookRequest{
				Book: &library.Book{
					Name:   "name",
					Author: "author",
					Title:  "title",
					Read:   true,
				},
			},
			expected: map[string]bigquery.Value{
				"book": map[string]bigquery.Value{
					"name":   "name",
					"author": "author",
					"title":  "title",
					"read":   true,
				},
			},
		},

		{
			name: "expr.Value (bool)",
			msg: &expr.Value{
				Kind: &expr.Value_BoolValue{
					BoolValue: true,
				},
			},
			expected: map[string]bigquery.Value{
				"bool_value": true,
			},
		},

		{
			name: "expr.Value (double)",
			msg: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
			expected: map[string]bigquery.Value{
				"double_value": float64(42),
			},
		},

		{
			name: "examplev1.ExampleMap",
			msg: &examplev1.ExampleMap{
				StringToString: map[string]string{
					"key1": "value1",
				},
				StringToNested: map[string]*examplev1.ExampleMap_Nested{
					"key1": {
						StringToString: map[string]string{
							"key1": "value1",
						},
					},
				},
				StringToEnum: map[string]examplev1.ExampleMap_Enum{
					"key1": examplev1.ExampleMap_ENUM_VALUE1,
				},
				Int32ToString: map[int32]string{
					1: "value1",
				},
				Int64ToString: map[int64]string{
					1: "value1",
				},
				Uint32ToString: map[uint32]string{
					1: "value1",
				},
				BoolToString: map[bool]string{
					true: "value1",
				},
			},
			expected: map[string]bigquery.Value{
				"string_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": "key1", "value": "value1"},
				},
				"string_to_nested": []bigquery.Value{
					map[string]bigquery.Value{
						"key": "key1",
						"value": map[string]bigquery.Value{
							"string_to_string": []bigquery.Value{
								map[string]bigquery.Value{"key": "key1", "value": "value1"},
							},
						},
					},
				},
				"string_to_enum": []bigquery.Value{
					map[string]bigquery.Value{"key": "key1", "value": "ENUM_VALUE1"},
				},
				"int32_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": int64(1), "value": "value1"},
				},
				"int64_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": int64(1), "value": "value1"},
				},
				"uint32_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": uint64(1), "value": "value1"},
				},
				"bool_to_string": []bigquery.Value{
					map[string]bigquery.Value{"key": true, "value": "value1"},
				},
			},
		},

		{
			name: "enum strings",
			msg: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
			expected: map[string]bigquery.Value{
				"enum_value": "ENUM_VALUE1",
			},
		},

		{
			name: "enum numbers",
			msg: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
			opt: MarshalOptions{
				Schema: SchemaOptions{
					UseEnumNumbers: true,
				},
			},
			expected: map[string]bigquery.Value{
				"enum_value": int64(1),
			},
		},

		{
			name: "primitive lists",
			msg: &examplev1.ExampleList{
				Int64List:  []int64{1, 2},
				StringList: []string{"a", "b"},
				EnumList:   []examplev1.ExampleList_Enum{examplev1.ExampleList_ENUM_VALUE1},
			},
			expected: map[string]bigquery.Value{
				"int64_list":  []bigquery.Value{int64(1), int64(2)},
				"string_list": []bigquery.Value{"a", "b"},
				"enum_list":   []bigquery.Value{"ENUM_VALUE1"},
			},
		},

		{
			name: "nested lists",
			msg: &examplev1.ExampleList{
				NestedList: []*examplev1.ExampleList_Nested{
					{StringList: []string{"a", "b"}},
					{StringList: []string{"c", "d"}},
				},
			},
			expected: map[string]bigquery.Value{
				"nested_list": []bigquery.Value{
					map[string]bigquery.Value{
						"string_list": []bigquery.Value{"a", "b"},
					},
					map[string]bigquery.Value{
						"string_list": []bigquery.Value{"c", "d"},
					},
				},
			},
		},

		{
			name: "wrappers",
			msg: &examplev1.ExampleWrappers{
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
			expected: map[string]bigquery.Value{
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
		},

		{
			name: "datetime (without offset)",
			msg: &examplev1.ExampleDateTime{
				DateTime: &datetime.DateTime{
					Year:    2021,
					Month:   int32(time.February),
					Day:     1,
					Hours:   8,
					Minutes: 30,
					Seconds: 1,
					Nanos:   2,
					TimeOffset: &datetime.DateTime_UtcOffset{
						UtcOffset: durationpb.New(2 * time.Hour),
					},
				},
			},
			opt: MarshalOptions{
				Schema: SchemaOptions{
					UseDateTimeWithoutOffset: true,
				},
			},
			expected: map[string]bigquery.Value{
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
		},

		{
			name: "oneof empty messages",
			msg: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofEmptyMessage_1{
					OneofEmptyMessage_1: &examplev1.ExampleOneof_EmptyMessage{},
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofEmptyMessage_2{
					OneofEmptyMessage_2: &examplev1.ExampleOneof_EmptyMessage{},
				},
			},
			opt: MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_empty_message_1",
				"oneof_fields_2": "oneof_empty_message_2",
			},
		},

		{
			name: "one oneof empty message, one non-empty",
			msg: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofEmptyMessage_1{
					OneofEmptyMessage_1: &examplev1.ExampleOneof_EmptyMessage{},
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofMessage{
					OneofMessage: &examplev1.ExampleOneof_Message{
						StringValue: "value",
					},
				},
			},
			opt: MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_empty_message_1",
				"oneof_fields_2": "oneof_message",
				"oneof_message": map[string]bigquery.Value{
					"string_value": "value",
				},
			},
		},

		{
			name: "oneof non-empty messages",
			msg: &examplev1.ExampleOneof{
				OneofFields_1: &examplev1.ExampleOneof_OneofBool_1{
					OneofBool_1: true,
				},
				OneofFields_2: &examplev1.ExampleOneof_OneofMessage{
					OneofMessage: &examplev1.ExampleOneof_Message{
						StringValue: "value",
					},
				},
			},
			opt: MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{
				"oneof_fields_1": "oneof_bool_1",
				"oneof_bool_1":   true,
				"oneof_fields_2": "oneof_message",
				"oneof_message": map[string]bigquery.Value{
					"string_value": "value",
				},
			},
		},

		{
			name:     "oneof empty root message",
			msg:      &examplev1.ExampleOneof{},
			opt:      MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{},
		},

		{
			name: "optional oneof and optional message",
			msg: &examplev1.ExampleOptional{
				OptMessage_2: &examplev1.ExampleOptMessage{
					StringValue: "opt_message_string_value",
				},
				OptOneofMessage_3: &examplev1.ExampleOptOneof{
					OneofFields_1: &examplev1.ExampleOptOneof_OneofMessage_2{
						OneofMessage_2: &examplev1.ExampleOptOneof_Message{
							StringValue: "opt_one_of_message_string_value",
						},
					},
				},
			},
			opt: MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{
				"opt_message_2": map[string]bigquery.Value{"string_value": string("opt_message_string_value")},
				"opt_oneof_message_3": map[string]bigquery.Value{
					"oneof_fields_1":  string("oneof_message_2"),
					"oneof_message_2": map[string]bigquery.Value{"string_value": string("opt_one_of_message_string_value")},
				},
			},
		},

		{
			name: "optional empty oneof and optional empty message",
			msg: &examplev1.ExampleOptional{
				OptMessage_2:      nil,
				OptOneofMessage_3: nil,
			},
			opt:      MarshalOptions{Schema: SchemaOptions{UseOneofFields: true}},
			expected: map[string]bigquery.Value{},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.opt.Marshal(tt.msg)
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
