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
		Genre
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
