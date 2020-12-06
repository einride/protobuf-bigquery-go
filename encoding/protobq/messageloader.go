package protobq

import (
	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/proto"
)

// MessageLoader implements bigquery.ValueLoader for a proto.Message.
// The message is converted from a BigQuery row using the provided UnmarshalOptions.
type MessageLoader struct {
	// Options to use for unmarshaling the Message.
	Options UnmarshalOptions

	// Message to load.
	Message proto.Message
}

var _ bigquery.ValueLoader = &MessageLoader{}

// Load implements bigquery.ValueLoader.
func (m *MessageLoader) Load(row []bigquery.Value, schema bigquery.Schema) error {
	return m.Options.Load(row, schema, m.Message)
}
