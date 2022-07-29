package genre

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetGenre endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{GetGenre: makeGetGenreEndpoint(s)}
}

func makeGetGenreEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGenreRequest)
		genre, err := s.GetGenre(ctx, req.Slug)
		return GetGenreResponse{
			Genre{
				BpmRange:         genre.BpmRange,
				Description:      genre.Description,
				Id:               genre.Id,
				IsParent:         genre.IsParent,
				KeyInstruments:   genre.KeyInstruments,
				KeyLocation:      genre.KeyLocation,
				ParentId:         genre.ParentId,
				ReleaseDate:      genre.ReleaseDate,
				ShortDescription: genre.ShortDescription,
				Slug:             genre.Slug,
				Title:            genre.Title,
			},
		}, err
	}
}
