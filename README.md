# Protobuf + BigQuery + Go

Go utilities for saving protocol buffers to BigQuery.

**Work in progress:** This library is under active development.

## Examples

### Schema inference

```go
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
```

### Marshaling

```go
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
```

### Unmarshaling

```go
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
```
