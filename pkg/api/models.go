package api

const (
	BASE_URL = "https://www.googleapis.com/customsearch/v1"
)

type SearchResponse struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Kind             string `json:"kind"`
	Title            string `json:"title"`
	HTMLTitle        string `json:"htmlTitle"`
	Link             string `json:"link"`
	DisplayLink      string `json:"displayLink"`
	Snippet          string `json:"snippet"`
	HTMLSnippet      string `json:"htmlSnippet"`
	CacheId          string `json:"cacheId"`
	FormattedURL     string `json:"formattedUrl"`
	HTMLFormattedURL string `json:"htmlFormattedUrl"`
}
