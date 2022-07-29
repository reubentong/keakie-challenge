package shows

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	GetShowRequest struct {
		Slug string `json:"slug"`
	}
	GetShowResponse struct {
		ShowResponse
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeShowReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetShowRequest
	vars := mux.Vars(r)

	req = GetShowRequest{
		Slug: vars["slug"],
	}
	return req, nil
}
