package protobq_test

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/google/go-cmp/cmp"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/testing/protocmp"
)

func ExampleInferSchema() {
	msg := &library.Book{}
	schema := protobq.InferSchema(msg)
	expected := bigquery.Schema{
		{Name: "name", Type: bigquery.StringFieldType},
		{Name: "author", Type: bigquery.StringFieldType},
		{Name: "title", Type: bigquery.StringFieldType},
		{Name: "read", Type: bigquery.BooleanFieldType},
	}
	fmt.Println(cmp.Equal(expected, schema))
	// Output: true
}

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
	fmt.Println(cmp.Equal(expected, row))
	// Output: true
}

func ExampleUnmarshal() {
	row := map[string]bigquery.Value{
		"name":   "publishers/123/books/456",
		"author": "P.L. Travers",
		"title":  "Mary Poppins",
		"read":   true,
	}
	msg := &library.Book{}
	if err := protobq.Unmarshal(row, msg); err != nil {
		// TODO: Handle error.
	}
	expected := &library.Book{
		Name:   "publishers/123/books/456",
		Author: "P.L. Travers",
		Title:  "Mary Poppins",
		Read:   true,
	}
	fmt.Println(cmp.Equal(expected, msg, protocmp.Transform()))
	// Output: true
}

func ExampleLoad() {
	row := []bigquery.Value{
		"publishers/123/books/456",
		"P.L. Travers",
		"Mary Poppins",
		true,
	}
	schema := bigquery.Schema{
		{Name: "name", Type: bigquery.StringFieldType},
		{Name: "author", Type: bigquery.StringFieldType},
		{Name: "title", Type: bigquery.StringFieldType},
		{Name: "read", Type: bigquery.BooleanFieldType},
	}
	msg := &library.Book{}
	if err := protobq.Load(row, schema, msg); err != nil {
		// TODO: Handle error.
	}
	expected := &library.Book{
		Name:   "publishers/123/books/456",
		Author: "P.L. Travers",
		Title:  "Mary Poppins",
		Read:   true,
	}
	fmt.Println(cmp.Equal(expected, msg, protocmp.Transform()))
	// Output: true
}
