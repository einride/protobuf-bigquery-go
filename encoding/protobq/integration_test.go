package protobq_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/public/v1"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"gotest.tools/v3/assert"
)

func Test_Integration_PublicDataSets(t *testing.T) {
	if _, err := google.FindDefaultCredentials(context.Background()); err != nil {
		t.Skip("Skipping integration test, missing Google credentials.")
	}
	for _, tt := range []*publicDataSetIntegrationTest{
		{
			ProjectID: "bigquery-public-data",
			DatasetID: "san_francisco_film_locations",
			TableID:   "film_locations",
			Limit:     100,
			Message:   &publicv1.FilmLocation{},
		},

		{
			ProjectID: "bigquery-public-data",
			DatasetID: "hacker_news",
			TableID:   "stories",
			Limit:     100,
			Message:   &publicv1.HackerNewsStory{},
		},
	} {
		tt.Run(t)
	}
}

type publicDataSetIntegrationTest struct {
	ProjectID        string
	DatasetID        string
	TableID          string
	Limit            int
	Message          proto.Message
	MarshalOptions   protobq.MarshalOptions
	UnmarshalOptions protobq.UnmarshalOptions
}

func (tt *publicDataSetIntegrationTest) Run(t *testing.T) {
	t.Run(fmt.Sprintf("%s.%s.%s", tt.ProjectID, tt.DatasetID, tt.TableID), func(t *testing.T) {
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
			proto.Reset(messageClone)
			assert.NilError(t, tt.UnmarshalOptions.Unmarshal(row, messageClone))
			assert.DeepEqual(t, messageLoader.Message, messageClone, protocmp.Transform())
		}
	})
}
