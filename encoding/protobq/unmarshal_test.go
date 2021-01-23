package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/example/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
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
				"name": "name",
				"book": map[string]bigquery.Value{
					"name":   "name",
					"author": "author",
					"title":  "title",
					"read":   true,
				},
			},
			expected: &library.UpdateBookRequest{
				Name: "name",
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
