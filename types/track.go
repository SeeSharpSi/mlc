package types

type Track struct {
	Name     string `json:"name"`
	URLs     []URLs `json:"external_urls"`
	Explicit bool   `json:"explicit"`
}

type URLs struct {
	URL string `json:"spotify"`
}
