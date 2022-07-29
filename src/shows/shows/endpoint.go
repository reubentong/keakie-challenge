package shows

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetShow endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{GetShow: makeGetShowEndpoint(s)}
}

func makeGetShowEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetShowRequest)
		show, err := s.GetShow(ctx, req.Slug)
		return GetShowResponse{ShowResponse{
			Assets: Asset{
				HeroL: HeroL{
					Jpg:  show.Assets.HeroL.Jpg,
					Webp: show.Assets.HeroL.Webp,
				},
			},
			Hosts:       show.Hosts,
			Description: show.Description,
			Format:      show.Format,
			Genres:      show.Genres,
			Id:          show.Id,
			IsApproved:  show.IsApproved,
			NodeType:    show.NodeType,
			Owners:      show.Owners,
			ReleaseDate: show.ReleaseDate,
			Slug:        show.Slug,
			Title:       show.Title,
		}}, err
	}
}
