package genre

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	GetGenre(ctx context.Context, slug string) (genre Genre, err error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) *service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) GetGenre(ctx context.Context, slug string) (genre Genre, err error) {
	logger := log.With(s.logger, "method", "GetGenre")
	genre, err = s.repository.GetGenre(ctx, slug)
	if err != nil {
		level.Error(logger).Log("err", err)
		return Genre{}, fmt.Errorf("error from repository: %v", err)
	}

	logger.Log("Get genre", slug)

	return genre, err
}
