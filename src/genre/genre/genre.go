package genre

import "context"

type Genre struct {
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

type Repository interface {
	GetGenre(ctx context.Context, slug string) (genre Genre, err error)
}
