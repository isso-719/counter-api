package infra

import (
	"cloud.google.com/go/datastore"
	"context"
	"errors"
	"github.com/isso-719/counter-api/src/domain"
	"github.com/isso-719/counter-api/src/repository"
)

type counterRepository struct {
	db *DB
	tx *datastore.Transaction
}

func NewCounterRepository(client *DB) repository.IFCounterRepository {
	return &counterRepository{
		db: client,
		tx: nil,
	}
}

func (r *counterRepository) BeginTx(ctx context.Context) error {
	tx, err := r.db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *counterRepository) CommitTx() error {
	if r.tx == nil {
		return errors.New("no transaction started")
	}
	_, err := r.tx.Commit()
	r.tx = nil
	return err
}

func (r *counterRepository) TxRead(key string) (*domain.Counter, error) {
	if r.tx == nil {
		return nil, errors.New("no transaction started")
	}

	k := datastore.NameKey("Counter", key, nil)
	count := &domain.Counter{}
	err := r.tx.Get(k, count)
	if errors.Is(err, datastore.ErrNoSuchEntity) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (r *counterRepository) TxWrite(key string, value *domain.Counter) (interface{}, error) {
	if r.tx == nil {
		return nil, errors.New("no transaction started")
	}

	k := datastore.NameKey("Counter", key, nil)

	_, err := r.tx.Put(k, value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
