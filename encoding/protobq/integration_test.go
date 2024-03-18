package protobq_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/public/v1"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func Test_Integration_PublicDataSets(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping integration test in short mode.")
	}
	if _, err := google.FindDefaultCredentials(context.Background()); err != nil {
		t.Skip("Skipping integration test, missing Google credentials.")
	}
	for _, tt := range []struct {
		ProjectID        string
		DatasetID        string
		TableID          string
		Limit            int
		Message          proto.Message
		MarshalOptions   protobq.MarshalOptions
		UnmarshalOptions protobq.UnmarshalOptions
	}{
		{
			ProjectID: "bigquery-public-data",
			DatasetID: "san_francisco_film_locations",
			TableID:   "film_locations",
			Limit:     10,
			Message:   &publicv1.FilmLocation{},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "london_bicycles",
			TableID:   "cycle_hire",
			Limit:     10,
			Message:   &publicv1.LondonBicycleRental{},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "san_francisco_transit_muni",
			TableID:   "stop_times",
			Limit:     10,
			Message:   &publicv1.SanFransiscoTransitStopTime{},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "london_bicycles",
			TableID:   "cycle_stations",
			Limit:     10,
			Message:   &publicv1.LondonBicycleStation{},
			UnmarshalOptions: protobq.UnmarshalOptions{
				DiscardUnknown: true, // Ignore non-snake case field "nbEmptyDocks".
			},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "noaa_historic_severe_storms",
			TableID:   "storms_2020",
			Limit:     10,
			Message:   &publicv1.HistoricSevereStorm{},
			MarshalOptions: protobq.MarshalOptions{
				Schema: protobq.SchemaOptions{UseDateTimeWithoutOffset: true},
			},
			UnmarshalOptions: protobq.UnmarshalOptions{
				Schema: protobq.SchemaOptions{UseDateTimeWithoutOffset: true},
			},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "geo_whos_on_first",
			TableID:   "geojson",
			Limit:     10,
			Message:   &publicv1.WhosOnFirstGeoJson{},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "crypto_dogecoin",
			TableID:   "transactions",
			Limit:     10,
			UnmarshalOptions: protobq.UnmarshalOptions{
				DiscardUnknown: true, // discard NUMERIC currency fields
			},
			Message: &publicv1.DogecoinTransaction{},
		},
	} {
		tt := tt
		t.Run(fmt.Sprintf("%s.%s.%s", tt.ProjectID, tt.DatasetID, tt.TableID), func(t *testing.T) {
			t.Parallel()
			client, err := bigquery.NewClient(context.Background(), tt.ProjectID)
			assert.NilError(t, err)
			t.Cleanup(func() {
				assert.NilError(t, client.Close())
			})
			messageLoader := &protobq.MessageLoader{
				Options: tt.UnmarshalOptions,
				Message: tt.Message,
			}
			rowIterator := client.Dataset(tt.DatasetID).Table(tt.TableID).Read(context.Background())
			for i := 0; i < tt.Limit; i++ {
				if err := rowIterator.Next(messageLoader); err != nil {
					if errors.Is(err, iterator.Done) {
						break
					}
					assert.NilError(t, err)
				}
				row, err := tt.MarshalOptions.Marshal(messageLoader.Message)
				assert.NilError(t, err)
				messageClone := proto.Clone(messageLoader.Message)
				assert.NilError(t, tt.UnmarshalOptions.Unmarshal(row, messageClone))
				assert.DeepEqual(t, messageLoader.Message, messageClone, protocmp.Transform())
			}
		})
	}
}
