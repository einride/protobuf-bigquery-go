# Protobuf + BigQuery + Go

[![PkgGoDev](https://pkg.go.dev/badge/go.einride.tech/protobuf-bigquery)](https://pkg.go.dev/go.einride.tech/protobuf-bigquery)
[![GoReportCard](https://goreportcard.com/badge/go.einride.tech/protobuf-bigquery)](https://goreportcard.com/report/go.einride.tech/protobuf-bigquery)
[![Codecov](https://codecov.io/gh/einride/protobuf-bigquery-go/branch/master/graph/badge.svg)](https://codecov.io/gh/einride/protobuf-bigquery-go)

Seamlessly save and load protocol buffers to and from BigQuery using Go.

This library provides add-ons to
[cloud.google.com/bigquery](https://pkg.go.dev/cloud.google.com/go/bigquery) for
first-class protobuf support using
[protobuf reflection](https://blog.golang.org/protobuf-apiv2).

## Installing

```bash
$ go get -u go.einride.tech/protobuf-bigquery
```

## Examples

### `protobq.InferSchema`

BigQuery schema inference for arbitrary protobuf messages.

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

### `protobq.MessageSaver`

An implementation of
[bigquery.ValueSaver](https://pkg.go.dev/cloud.google.com/go/bigquery#ValueSaver)
that saves arbitrary protobuf messages to BigQuery.

```go
func ExampleMessageSaver() {
	ctx := context.Background()
	// Write protobuf messages to a BigQuery table.
	projectID := flag.String("project", "", "BigQuery project to write to.")
	datasetID := flag.String("dataset", "", "BigQuery dataset to write to.")
	tableID := flag.String("table", "", "BigQuery table to write to.")
	create := flag.Bool("create", false, "Flag indicating whether to create the table.")
	flag.Parse()
	// Connect to BigQuery.
	client, err := bigquery.NewClient(ctx, *projectID)
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	table := client.Dataset(*datasetID).Table(*tableID)
	// Create the table by inferring the BigQuery schema from the protobuf schema.
	if *create {
		if err := table.Create(ctx, &bigquery.TableMetadata{
			Schema: protobq.InferSchema(&publicv1.FilmLocation{}),
		}); err != nil {
			panic(err) // TODO: Handle error.
		}
	}
	// Insert the protobuf messages.
	inserter := table.Inserter()
	for i, filmLocation := range []*publicv1.FilmLocation{
		{Title: "Dark Passage", ReleaseYear: 1947, Locations: "Filbert Steps"},
		{Title: "D.O.A", ReleaseYear: 1950, Locations: "Union Square"},
		{Title: "Flower Drum Song", ReleaseYear: 1961, Locations: "Chinatown"},
	} {
		if err := inserter.Put(ctx, &protobq.MessageSaver{
			Message:  filmLocation,
			InsertID: strconv.Itoa(i), // include an optional insert ID
		}); err != nil {
			panic(err) // TODO: Handle error.
		}
	}
}
```

### `protobq.MessageLoader`

An implementation of
[bigquery.ValueLoader](https://pkg.go.dev/cloud.google.com/go/bigquery#ValueLoader)
that loads arbitrary protobuf messages from BigQuery.

```go
func ExampleMessageLoader() {
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
		Message: &publicv1.FilmLocation{},
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

*[Reference ≫](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)*

### Support for API Common Protos (`google.type`)

| Protobuf              | BigQuery             |
| --------------------- | -------------------- |
| google.type.Date      | DATE                 |
| google.type.DateTime  | RECORD (or DATETIME) |
| google.type.LatLng    | GEOGRAPHY            |
| google.type.TimeOfDay | TIME                 |

*[Reference ≫](https://github.com/googleapis/api-common-protos)*
