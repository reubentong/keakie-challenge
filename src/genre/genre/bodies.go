package genre

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	GetGenreRequest struct {
		Slug string `json:"slug"`
	}
	GetGenreResponse struct {
		BpmRange         string `json:"bpm_range"`
		Description      string `json:"description"`
		Id               string `json:"id"`
		IsParent         bool   `json:"is_parent"`
		KeyInstruments   string `json:"key_instruments"`
		KeyLocation      string `json:"key_location"`
		ParentId         string `json:"parent_id"`
		ReleaseDate      string `json:"release_date"`
		ShortDescription string `json:"short_description"`
		Slug             string `json:"slug"`
		Title            string `json:"title"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGenreReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetGenreRequest
	vars := mux.Vars(r)

	req = GetGenreRequest{
		Slug: vars["slug"],
	}
	return req, nil
}
