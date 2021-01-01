# Protobuf + BigQuery + Go

Add-ons to [cloud.google.com/bigquery][google-cloud-go-bigquery] for
first-class protobuf support.

**Work in progress:** This library is under active development.

[google-cloud-go-bigquery]: https://pkg.go.dev/cloud.google.com/go/bigquery

## Installing

```bash
$ go get -u go.einride.tech/protobuf-bigquery
```

## Features

### Support for Well-Known Types (`google.protobuf`)

| Protobuf                    | BigQuery        |
| --------------------------- | --------------- |
| google.protobuf.Timestamp   | TIMESTAMP       |
| google.protobuf.Duration    | FLOAT (seconds) |
| google.protobuf.DoubleValue | FLOAT           |
| google.protobuf.FloatValue  | FLOAT           |
| google.protobuf.Int32Value  | INTEGER         |
| google.protobuf.Int64Value  | INTEGER         |
| google.protobuf.Uint32Value | INTEGER         |
| google.protobuf.Uint64Value | INTEGER         |
| google.protobuf.BoolValue   | BOOLEAN         |
| google.protobuf.StringValue | STRING          |
| google.protobuf.BytesValue  | BYTES           |
| google.protobuf.StructValue | STRING (JSON)   |

_[Reference ≫][well-known-types]_

[well-known-types]: https://developers.google.com/protocol-buffers/docs/reference/google.protobuf

### Support for API Common Protos (`google.type`)

| Protobuf             | BigQuery  |
| -------------------- | --------- |
| google.type.Date     | DATE      |
| google.type.DateTime | TIMESTAMP |
| google.type.LatLng   | GEOGRAPHY |
| google.type.Time     | TIME      |

_[Reference ≫][api-common-protos]_

[api-common-protos]: https://github.com/googleapis/api-common-protos

## Examples

### `protobq.MessageLoader`

Loads BigQuery rows into protobuf messages.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/example/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/prototext"
)

func main() {
	ctx := context.Background()
	// Read from the public "film locations" BigQuery dataset into a proto message.
	const (
		project = "bigquery-public-data"
		dataset = "san_francisco_film_locations"
		table   = "film_locations"
	)
	// Connect to BigQuery.
	client, err := bigquery.NewClient(ctx, project)
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	// Load BigQuery rows into a FilmLocation message.
	messageLoader := &protobq.MessageLoader{
		Message: &examplev1.FilmLocation{},
	}
	// Iterate rows in table.
	rowIterator := client.Dataset(dataset).Table(table).Read(ctx)
	for {
		// Load next row into the FilmLocation message.
		if err := rowIterator.Next(messageLoader); err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			panic(err) // TODO: Handle error.
		}
		// Print the message.
		fmt.Println(prototext.Format(messageLoader.Message))
	}
}
```

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
