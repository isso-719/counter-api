package infra

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/isso-719/counter-api/src/infra/config"
)

type DB struct {
	*datastore.Client
}

type TX struct {
	*datastore.Transaction
}

func CreateDatastoreClient(ctx context.Context) (*DB, error) {
	projectID := config.LoadDatastoreConfig().ProjectID

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}

	return &DB{client}, nil
}
