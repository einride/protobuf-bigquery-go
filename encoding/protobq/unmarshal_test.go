package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/example/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
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
