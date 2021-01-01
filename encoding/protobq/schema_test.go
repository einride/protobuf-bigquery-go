package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/example/v1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

func TestSchemaOptions_InferSchema(t *testing.T) {
	for _, tt := range []struct {
		name     string
		msg      proto.Message
		opt      SchemaOptions
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

		{
			name: "examplev1.ExampleMap",
			msg:  &examplev1.ExampleMap{},
			expected: bigquery.Schema{
				{
					Name:     "string_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "string_to_nested",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{
							Name: "value",
							Type: bigquery.RecordFieldType,
							Schema: bigquery.Schema{
								{
									Name:     "string_to_string",
									Type:     bigquery.RecordFieldType,
									Repeated: true,
									Schema: bigquery.Schema{
										{Name: "key", Type: bigquery.StringFieldType},
										{Name: "value", Type: bigquery.StringFieldType},
									},
								},
							},
						},
					},
				},

				{
					Name:     "string_to_enum",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "int32_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "int64_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "uint32_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.opt.InferSchema(tt.msg)
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
