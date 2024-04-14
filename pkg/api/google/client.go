package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchAPI struct {
	URL string
}

type APIConfig struct {
	SearchEngineId string `json:"search_engine_id"`
	ApiKey         string `json:"api_key"`
}

func New(cfg *APIConfig) *SearchAPI {
	formattedUrl := fmt.Sprintf(BASE_URL+"?key=%s&cx=%s", cfg.ApiKey, cfg.SearchEngineId)
	return &SearchAPI{
		URL: formattedUrl,
	}
}

func (s *SearchAPI) GetSearchResults(query string) ([]*Item, error) {
	encodedQuery := url.QueryEscape(query)
	queryURL := fmt.Sprintf(s.URL+"&q=%s", encodedQuery)
	client := http.Client{}
	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	results := &SearchResponse{}
	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}
	return results.Items, nil
}
