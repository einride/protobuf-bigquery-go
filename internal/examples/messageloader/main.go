package main

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/public/v1"
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
