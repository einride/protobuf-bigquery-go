package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

func TestValue(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
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
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Marshal(tt.msg)
			assert.NilError(t, err)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
