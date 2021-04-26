package jsonbq

import (
	"encoding/json"

	"cloud.google.com/go/bigquery"
)

// Different modes of a BigQuery field.
const (
	ModeRepeated string = "REPEATED"
	ModeRequired string = "REQUIRED"
	ModeNullable string = "NULLABLE"
)

// Field describes the schema of a field in BigQuery.
type Field struct {
	Name        string             `json:"name"`
	Type        bigquery.FieldType `json:"type"`
	Mode        string             `json:"mode"`
	Description string             `json:"description,omitempty"`
	Fields      []*Field           `json:"fields,omitempty"`
}

// MarshalSchema marshals a BigQuery schema to a valid BigQuery JSON schema.
// The JSON conforms to the format in this example:
// https://cloud.google.com/bigquery/docs/nested-repeated#example_schema
func MarshalSchema(schema bigquery.Schema) ([]byte, error) {
	fields := convertSchema(schema)
	return json.Marshal(fields)
}

func convertSchema(schema bigquery.Schema) []*Field {
	fields := make([]*Field, 0, len(schema))
	for _, fieldSchema := range schema {
		fields = append(fields, convertField(fieldSchema))
	}
	return fields
}

func convertField(fieldSchema *bigquery.FieldSchema) *Field {
	var field Field
	field.Name = fieldSchema.Name
	field.Type = fieldSchema.Type
	switch {
	case fieldSchema.Repeated:
		field.Mode = ModeRepeated
	case fieldSchema.Required:
		field.Mode = ModeRequired
	default:
		field.Mode = ModeNullable
	}
	field.Description = fieldSchema.Description
	field.Fields = convertSchema(fieldSchema.Schema)
	return &field
}
