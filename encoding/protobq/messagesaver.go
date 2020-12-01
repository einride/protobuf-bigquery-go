package protobq

import (
	"cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/proto"
)

type MessageSaver struct {
	Message  proto.Message
	InsertID string
}

var _ bigquery.ValueSaver = &MessageSaver{}

func (m *MessageSaver) Save() (map[string]bigquery.Value, string, error) {
	row, err := Marshal(m.Message)
	if err != nil {
		return nil, "", err
	}
	return row, m.InsertID, nil
}
