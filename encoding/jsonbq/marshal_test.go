package jsonbq

import (
	"encoding/json"
	"testing"

	"cloud.google.com/go/bigquery"
	"gotest.tools/v3/assert"
)

func TestMarshalSchema(t *testing.T) {
	for _, tt := range []struct {
		name     string
		schema   bigquery.Schema
		expected []*Field
	}{
		{
			name: "library.Book",
			schema: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{Name: "author", Type: bigquery.StringFieldType},
				{Name: "title", Type: bigquery.StringFieldType},
				{Name: "read", Type: bigquery.BooleanFieldType},
			},
			expected: []*Field{
				{
					Name: "name",
					Type: bigquery.StringFieldType,
					Mode: ModeNullable,
				},
				{
					Name: "author",
					Type: bigquery.StringFieldType,
					Mode: ModeNullable,
				},
				{
					Name: "title",
					Type: bigquery.StringFieldType,
					Mode: ModeNullable,
				},
				{
					Name: "read",
					Type: bigquery.BooleanFieldType,
					Mode: ModeNullable,
				},
			},
		},
		{
			name: "library.UpdateBookRequest",
			schema: bigquery.Schema{
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
				{
					Name: "update_mask",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "paths", Type: bigquery.StringFieldType, Repeated: true},
					},
				},
			},
			expected: []*Field{
				{
					Name: "book",
					Type: bigquery.RecordFieldType,
					Mode: ModeNullable,
					Fields: []*Field{
						{
							Name: "name",
							Type: bigquery.StringFieldType,
							Mode: ModeNullable,
						},
						{
							Name: "author",
							Type: bigquery.StringFieldType,
							Mode: ModeNullable,
						},
						{
							Name: "title",
							Type: bigquery.StringFieldType,
							Mode: ModeNullable,
						},
						{
							Name: "read",
							Type: bigquery.BooleanFieldType,
							Mode: ModeNullable,
						},
					},
				},
				{
					Name: "update_mask",
					Type: bigquery.RecordFieldType,
					Mode: ModeNullable,
					Fields: []*Field{
						{
							Name: "paths",
							Type: bigquery.StringFieldType,
							Mode: ModeRepeated,
						},
					},
				},
			},
		},
		{
			name: "examplev1.ExampleDateTime (no offset)",
			schema: bigquery.Schema{
				{Name: "date_time", Type: bigquery.DateTimeFieldType},
			},
			expected: []*Field{
				{
					Name: "date_time",
					Type: bigquery.DateTimeFieldType,
					Mode: ModeNullable,
				},
			},
		},
		{
			name: "examplev1.ExampleDateTime (with offset)",
			schema: bigquery.Schema{
				{
					Name: "date_time",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "datetime", Type: bigquery.DateTimeFieldType},
						{Name: "utc_offset", Type: bigquery.FloatFieldType},
						{
							Name: "time_zone",
							Type: bigquery.RecordFieldType,
							Schema: bigquery.Schema{
								{Name: "id", Type: bigquery.StringFieldType},
								{Name: "version", Type: bigquery.StringFieldType},
							},
						},
					},
				},
			},
			expected: []*Field{
				{
					Name: "date_time",
					Type: bigquery.RecordFieldType,
					Mode: ModeNullable,
					Fields: []*Field{
						{
							Name: "datetime",
							Type: bigquery.DateTimeFieldType,
							Mode: ModeNullable,
						},
						{
							Name: "utc_offset",
							Type: bigquery.FloatFieldType,
							Mode: ModeNullable,
						},
						{
							Name: "time_zone",
							Type: bigquery.RecordFieldType,
							Mode: ModeNullable,
							Fields: []*Field{
								{
									Name: "id",
									Type: bigquery.StringFieldType,
									Mode: ModeNullable,
								},
								{
									Name: "version",
									Type: bigquery.StringFieldType,
									Mode: ModeNullable,
								},
							},
						},
					},
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalSchema(tt.schema)
			assert.NilError(t, err)
			want, err := json.Marshal(tt.expected)
			assert.NilError(t, err)
			assert.Equal(t, string(got), string(want))
		})
	}
}
