package protobq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func TestMessageLoader_Load(t *testing.T) {
	messageLoader := MessageLoader{
		Message: &library.Book{},
	}
	row := []bigquery.Value{
		"name",
		"author",
		"title",
		true,
	}
	schema := bigquery.Schema{
		{Name: "name", Type: bigquery.StringFieldType},
		{Name: "author", Type: bigquery.StringFieldType},
		{Name: "title", Type: bigquery.StringFieldType},
		{Name: "read", Type: bigquery.BooleanFieldType},
	}
	assert.NilError(t, messageLoader.Load(row, schema))
	expected := &library.Book{
		Name:   "name",
		Author: "author",
		Title:  "title",
		Read:   true,
	}
	assert.DeepEqual(t, expected, messageLoader.Message, protocmp.Transform())
}
