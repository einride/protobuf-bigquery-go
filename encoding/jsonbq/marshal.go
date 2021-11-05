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

// MarshalSchema marshals a BigQuery schema to a valid BigQuery JSON schema based on
// default MarshalOptions. The JSON conforms to the format in this example:
// https://cloud.google.com/bigquery/docs/nested-repeated#example_schema
func MarshalSchema(schema bigquery.Schema) ([]byte, error) {
	return MarshalOptions{}.MarshalSchema(schema)
}

// MarshalOptions is a configurable BigQuery schema json marshaller.
type MarshalOptions struct {
	// Indent specifies the set of indentation characters to use in a formatted
	// output such that every entry is preceded by Indent and terminated by a newline.
	// Leave Indent empty to disable formatted output.
	Indent string
}

// MarshalSchema marshals a BigQuery schema to a valid BigQuery JSON schema based on
// the given MarshalOptions. The JSON conforms to the format in this example:
// https://cloud.google.com/bigquery/docs/nested-repeated#example_schema
func (o MarshalOptions) MarshalSchema(schema bigquery.Schema) ([]byte, error) {
	fields := convertSchema(schema)
	if o.Indent != "" {
		return json.MarshalIndent(fields, "", o.Indent)
	}
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
