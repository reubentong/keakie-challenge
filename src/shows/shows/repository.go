package shows

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
)

type repository struct {
	db     []Show
	logger log.Logger
}

func NewRepository(db []Show, logger log.Logger) Repository {
	return repository{
		db:     db,
		logger: log.With(logger, "repository"),
	}
}

func (r repository) GetShow(slug string) (show ShowResponse, err error) {
	// for verifying slug exists
	found := false

	var genres []Genre
	for _, element := range r.db {
		if element.Slug == slug {
			found = true
			for _, genre := range element.Genres {
				resp, err := http.Get("http://localhost:8080/genres/" + genre)
				if err != nil {
					level.Error(r.logger).Log("err", "error retrieving genres")
					return ShowResponse{}, fmt.Errorf("error from repository: %v", err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					level.Error(r.logger).Log("err", "error reading body")
					return ShowResponse{}, fmt.Errorf("error from repository: %v", err)

				}
				genreResponse := Genre{}
				err = json.Unmarshal(body, &genreResponse)
				if err != nil {
					level.Error(r.logger).Log("err", "error unmarshalling body")
					return ShowResponse{}, fmt.Errorf("error from repository: %v", err)
				}

				genres = append(genres, genreResponse)
			}

			show = ShowResponse{
				Assets: Asset{
					HeroL: HeroL{
						Jpg:  element.Assets.HeroL.Jpg,
						Webp: element.Assets.HeroL.Webp,
					}},
				Hosts:       element.Hosts,
				Description: element.Description,
				Format:      element.Format,
				Genres:      genres,
				Id:          element.Id,
				IsApproved:  element.IsApproved,
				NodeType:    element.NodeType,
				Owners:      element.Owners,
				ReleaseDate: element.ReleaseDate,
				Slug:        element.Slug,
				Title:       element.Title,
			}
			return
		}
	}

	if found == false {
		err := "slug does not exist"
		level.Error(r.logger).Log("err", err)
		return ShowResponse{}, fmt.Errorf(err)
	}

	return show, err
}
