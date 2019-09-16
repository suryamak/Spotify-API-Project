package spotify_objects

type PagingObject struct {
	Href string `json:"href"`
	Items []*Artist `json:"items"`
	Limit int `json:"limit"`
	Next string `json:"next"`
	Offset int `json:"offset"`
	Previous string `json:"previous"`
	Total int `json:"total"`
}
type Artist struct {
	External_urls *ExternalURL `json:"external_urls"`
	Followers *Followers `json:"followers"`
	Genres []string `json:"genres"`
	Href string `json:"href"`
	ID string `json:"id"`
	Images []*Image `json:"images"`
	Name string `json:"name"`
	Popularity int `json:"popularity"`
	Type string `json:"type"`
	URI string `json:"uri"`
	}

type ExternalURL struct {
	Key string `json:"{key}"`
	Value string `json:"{value}"`
}

type Followers struct {
	Href string `json:"href"`
	Total int `json:"total"`
}

type Image struct {
	Height int `json:"height"`
	URL string `json:"url"`
	Width int 	`json:"width"`
}
