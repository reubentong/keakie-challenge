package genre

import (
	"context"
	"fmt"
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
	// for verifying slug exists
	found := false
	for _, element := range r.db {
		if element.Slug == slug {
			genre = element
			found = true
			return
		}
	}

	if found == false {
		err := "slug does not exist"
		level.Error(r.logger).Log("err", err)
		return Genre{}, fmt.Errorf(err)
	}
	return genre, err
}
