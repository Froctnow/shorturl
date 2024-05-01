package models

type URLFromFile struct {
	UUID      string `json:"uuid"`
	ShortURL  string `json:"short_url"`
	OriginURL string `json:"origin_url"`
}
