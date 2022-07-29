package shows

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	GetShow(ctx context.Context, slug string) (show ShowResponse, err error)
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

func (s *service) GetShow(ctx context.Context, slug string) (show ShowResponse, err error) {
	logger := log.With(s.logger, "method", "GetShow")
	show, err = s.repository.GetShow(slug)
	if err != nil {
		level.Error(logger).Log("err", err)
		return ShowResponse{}, fmt.Errorf("error from repository: %v", err)
	}
	return show, err
}
