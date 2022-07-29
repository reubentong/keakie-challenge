package genre

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type repository struct {
	db     []Genre
	logger log.Logger
}

func NewRepository(db []Genre, logger log.Logger) Repository {
	return repository{
		db:     db,
		logger: log.With(logger, "repository"),
	}
}

func (r repository) GetGenre(ctx context.Context, slug string) (genre Genre, err error) {
	for _, element := range r.db {
		if element.Slug == slug {
			genre = element
			return genre, nil
		} else {
			level.Error(r.logger).Log("Error finding slug", err)
			return Genre{}, err
		}
	}

	return genre, nil
}
