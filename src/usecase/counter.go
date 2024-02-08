package usecase

import (
	"context"
	"errors"
	"github.com/isso-719/counter-api/src/domain"
	"github.com/isso-719/counter-api/src/repository"
	"regexp"
)

type IFCounterService interface {
	Increment(ctx context.Context, url string) (*domain.Counter, error)
}

type counterService struct {
	counterRepository repository.IFCounterRepository
}

func NewCounterService(counterRepository repository.IFCounterRepository) IFCounterService {
	return &counterService{
		counterRepository: counterRepository,
	}
}

func (c *counterService) Increment(ctx context.Context, url string) (*domain.Counter, error) {
	if url == "" {
		return nil, errors.New("url is required")
	}
	r := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9-_.]+.[a-zA-Z0-9-.]+$`)
	if !r.MatchString(url) {
		return nil, errors.New("url is invalid")
	}

	err := c.counterRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer c.counterRepository.CommitTx()

	count, err := c.counterRepository.TxRead(url)
	if err != nil {
		return nil, err
	}
	if count == nil {
		count = &domain.Counter{
			URL:   url,
			Count: 0,
		}
	}

	count.Count++

	_, err = c.counterRepository.TxWrite(url, count)
	if err != nil {
		return nil, err
	}

	return count, nil
}
