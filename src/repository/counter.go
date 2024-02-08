package repository

//go:generate mockgen -source=counter.go -destination=counter_mock.go -package=repository

import (
	"context"
	"github.com/isso-719/counter-api/src/domain"
)

type IFCounterRepository interface {
	BeginTx(ctx context.Context) error
	CommitTx() error
	TxRead(key string) (*domain.Counter, error)
	TxWrite(key string, value *domain.Counter) (interface{}, error)
}
