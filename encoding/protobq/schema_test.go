package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

func TestSchema(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		expected bigquery.Schema
	}{
		{
			name: "library.Book",
			msg:  &library.Book{},
			expected: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{Name: "author", Type: bigquery.StringFieldType},
				{Name: "title", Type: bigquery.StringFieldType},
				{Name: "read", Type: bigquery.BooleanFieldType},
			},
		},
		{
			name: "library.UpdateBookRequest",
			msg:  &library.UpdateBookRequest{},
			expected: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{
					Name: "book",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "name", Type: bigquery.StringFieldType},
						{Name: "author", Type: bigquery.StringFieldType},
						{Name: "title", Type: bigquery.StringFieldType},
						{Name: "read", Type: bigquery.BooleanFieldType},
					},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Schema(tt.msg)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
