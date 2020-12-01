package protobq_test

import (
	"fmt"
	"reflect"

	"cloud.google.com/go/bigquery"
	"github.com/einride/protobuf-bigquery-go/encoding/protobq"
	"google.golang.org/genproto/googleapis/example/library/v1"
)

func ExampleMarshal() {
	msg := &library.Book{
		Name:   "publishers/123/books/456",
		Author: "P.L. Travers",
		Title:  "Mary Poppins",
		Read:   true,
	}
	row, err := protobq.Marshal(msg)
	if err != nil {
		// TODO: Handle error.
	}
	expected := map[string]bigquery.Value{
		"name":   "publishers/123/books/456",
		"author": "P.L. Travers",
		"title":  "Mary Poppins",
		"read":   true,
	}
	fmt.Println(reflect.DeepEqual(expected, row))
	// Output: true
}

func ExampleInferSchema() {
	msg := &library.Book{}
	schema := protobq.InferSchema(msg)
	expected := bigquery.Schema{
		{Name: "name", Type: bigquery.StringFieldType},
		{Name: "author", Type: bigquery.StringFieldType},
		{Name: "title", Type: bigquery.StringFieldType},
		{Name: "read", Type: bigquery.BooleanFieldType},
	}
	fmt.Println(reflect.DeepEqual(expected, schema))
	// Output: true
}
