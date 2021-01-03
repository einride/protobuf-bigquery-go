package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/example/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
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
				Name: "name",
				Book: &library.Book{
					Name:   "name",
					Author: "author",
					Title:  "title",
					Read:   true,
				},
			},
			expected: map[string]bigquery.Value{
				"name": "name",
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
					false: "value1",
					true:  "value2",
				},
			},
			expected: map[string]bigquery.Value{
				"string_to_string": map[string]bigquery.Value{
					"key1": "value1",
				},
				"string_to_nested": map[string]bigquery.Value{
					"key1": map[string]bigquery.Value{
						"string_to_string": map[string]bigquery.Value{
							"key1": "value1",
						},
					},
				},
				"string_to_enum": map[string]bigquery.Value{
					"key1": "ENUM_VALUE1",
				},
				"int32_to_string": map[int64]bigquery.Value{
					1: "value1",
				},
				"int64_to_string": map[int64]bigquery.Value{
					1: "value1",
				},
				"uint32_to_string": map[uint64]bigquery.Value{
					1: "value1",
				},
				"bool_to_string": map[bool]bigquery.Value{
					false: "value1",
					true:  "value2",
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
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.opt.Marshal(tt.msg)
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
