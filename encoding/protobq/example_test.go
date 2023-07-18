package protobq_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/bigquery"
	"github.com/google/go-cmp/cmp"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/public/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/protobuf/encoding/prototext"
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

func ExampleMarshal() {
	msg := &library.Book{
		Name:   "publishers/123/books/456",
		Author: "P.L. Travers",
		Title:  "Mary Poppins",
		Read:   true,
	}
	row, err := protobq.Marshal(msg)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
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
