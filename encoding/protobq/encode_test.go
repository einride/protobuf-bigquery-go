package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
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
			name: "expr.Value (bool)",
			msg: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
			expected: map[string]bigquery.Value{
				"double_value": float64(42),
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
