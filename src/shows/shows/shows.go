package shows

import (
	"time"
)

type Show struct {
	Assets      Asset     `json:"assets"`
	Hosts       []string  `json:"hosts"`
	Description string    `json:"description"`
	Format      string    `json:"format"`
	Genres      []string  `json:"genres"`
	Id          string    `json:"id"`
	IsApproved  bool      `json:"is_approved"`
	NodeType    string    `json:"node_type"`
	Owners      []string  `json:"owners"`
	ReleaseDate time.Time `json:"release_date"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
}

type ShowResponse struct {
	Assets      Asset     `json:"assets"`
	Hosts       []string  `json:"hosts"`
	Description string    `json:"description"`
	Format      string    `json:"format"`
	Genres      []Genre   `json:"genres"`
	Id          string    `json:"id"`
	IsApproved  bool      `json:"is_approved"`
	NodeType    string    `json:"node_type"`
	Owners      []string  `json:"owners"`
	ReleaseDate time.Time `json:"release_date"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
}

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

type Asset struct {
	HeroL HeroL `json:"hero_l"`
}

type HeroL struct {
	Jpg  string `json:"jpg"`
	Webp string `json:"webp"`
}

type Repository interface {
	GetShow(slug string) (shows ShowResponse, err error)
}
