package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/example/v1"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/public/v1"
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

				{
					Name:     "bool_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.BooleanFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "string_to_float_value",
					Repeated: true,
					Type:     bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.FloatFieldType},
					},
				},
			},
		},

		{
			name: "examplev1.ExampleEnum",
			msg:  &examplev1.ExampleEnum{},
			expected: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.StringFieldType},
			},
		},

		{
			name: "examplev1.ExampleEnum (UseEnumNumbers)",
			msg:  &examplev1.ExampleEnum{},
			opt: SchemaOptions{
				UseEnumNumbers: true,
			},
			expected: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.IntegerFieldType},
			},
		},

		{
			name: "examplev1.ExampleWrappers",
			msg:  &examplev1.ExampleWrappers{},
			expected: bigquery.Schema{
				{Name: "float_value", Type: bigquery.FloatFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "string_value", Type: bigquery.StringFieldType},
				{Name: "bytes_value", Type: bigquery.BytesFieldType},
				{Name: "int32_value", Type: bigquery.IntegerFieldType},
				{Name: "int64_value", Type: bigquery.IntegerFieldType},
				{Name: "uint32_value", Type: bigquery.IntegerFieldType},
				{Name: "uint64_value", Type: bigquery.IntegerFieldType},
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
			},
		},

		{
			name: "examplev1.ExampleDateTime (no offset)",
			msg:  &examplev1.ExampleDateTime{},
			opt: SchemaOptions{
				UseDateTimeWithoutOffset: true,
			},
			expected: bigquery.Schema{
				{Name: "date_time", Type: bigquery.DateTimeFieldType},
			},
		},

		{
			name: "examplev1.ExampleDateTime (with offset)",
			msg:  &examplev1.ExampleDateTime{},
			expected: bigquery.Schema{
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
		},

		{
			name: "examplev1.ExampleOneof",
			msg:  &examplev1.ExampleOneof{},
			opt: SchemaOptions{
				UseOneofFields: true,
			},
			expected: bigquery.Schema{
				{
					Name: "oneof_bool_1",
					Type: bigquery.BooleanFieldType,
				},
				{
					Name: "oneof_message",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{
							Name: "string_value",
							Type: bigquery.StringFieldType,
						},
					},
				},
				{
					Name:        "oneof_fields_1",
					Type:        bigquery.StringFieldType,
					Description: "One of: oneof_empty_message_1, oneof_bool_1.",
				},
				{
					Name:        "oneof_fields_2",
					Type:        bigquery.StringFieldType,
					Description: "One of: oneof_empty_message_2, oneof_message.",
				},
			},
		},

		{
			name: "examplev1.ExampleAnnotations",
			msg:  &examplev1.ExampleAnnotations{},
			opt: SchemaOptions{
				UseModeFromFieldBehavior: true,
			},
			expected: bigquery.Schema{
				{
					Name:     "required_int",
					Type:     bigquery.IntegerFieldType,
					Required: true,
				},
				{
					Name:     "optional_int",
					Type:     bigquery.IntegerFieldType,
					Required: false,
				},
			},
		},

		{
			name: "publicv1.WhosOnFirstGeoJson",
			msg:  &publicv1.WhosOnFirstGeoJson{},
			expected: bigquery.Schema{
				{
					Name:     "geoid",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "id",
					Type:     bigquery.IntegerFieldType,
					Required: false,
				},
				{
					Name:     "body",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "geometry_type",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "bounding_box",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "geom",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "last_modified",
					Type:     bigquery.IntegerFieldType,
					Required: false,
				},
				{
					Name:     "last_modified_timestamp",
					Type:     bigquery.TimestampFieldType,
					Required: false,
				},
			},
		},

		{
			name: "publicv1.WhosOnFirstGeoJson (with enable `UseJSONStructs`)",
			opt: SchemaOptions{
				UseJSONStructs: true,
			},
			msg: &publicv1.WhosOnFirstGeoJson{},
			expected: bigquery.Schema{
				{
					Name:     "geoid",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "id",
					Type:     bigquery.IntegerFieldType,
					Required: false,
				},
				{
					Name:     "body",
					Type:     bigquery.JSONFieldType,
					Required: false,
				},
				{
					Name:     "geometry_type",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "bounding_box",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "geom",
					Type:     bigquery.StringFieldType,
					Required: false,
				},
				{
					Name:     "last_modified",
					Type:     bigquery.IntegerFieldType,
					Required: false,
				},
				{
					Name:     "last_modified_timestamp",
					Type:     bigquery.TimestampFieldType,
					Required: false,
				},
			},
		},
		{
			name: "examplev1.ExampleStruct (with enable `UseJSONValues`)",
			opt: SchemaOptions{
				UseJSONValues: true,
			},
			msg: &examplev1.ExampleStruct{},
			expected: bigquery.Schema{
				{
					Name:     "value",
					Type:     bigquery.JSONFieldType,
					Required: false,
				},
			},
		},
		{
			name: "examplev1.ExampleOptional (with enable `UseOneofFields`)",
			opt: SchemaOptions{
				UseOneofFields: true,
			},
			msg: &examplev1.ExampleOptional{},
			expected: bigquery.Schema{
				{
					Name:     "opt",
					Type:     bigquery.FloatFieldType,
					Required: false,
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
