package openai

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"encoding/json"
)

type OpenAIConfig struct {
	URL    string `json:"url"`
	ApiKey string `json:"api_key"`
	Model  string `json:"model"`
	EmbDim int    `json:"embedding_dimensions"`
}

type OpenAIAPI struct {
	URL        string
	ApiKey     string
	Model      string
	Dimensions int
}

func New(cfg *OpenAIConfig) *OpenAIAPI {
	return &OpenAIAPI{
		URL:        cfg.URL,
		ApiKey:     cfg.ApiKey,
		Model:      cfg.Model,
		Dimensions: cfg.EmbDim,
	}

}

func (c *OpenAIAPI) GetEmbedding(req *Request) ([]*EmbeddingData, error) {
	req.Model = c.Model
	req.Dimensions = c.Dimensions

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	request, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+c.ApiKey)
	request.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "__cf_bm=CDgubRWqao5rA_oaQlyK1_serKGjZBfSSJ3Mia_S1Ns-1713010726-1.0.1.1-nfszg48tp3VSQOTrC09kxGlp9C4dvdnLbi5NqGCT9pA.s_vXg.1HViNM12fGcfB1fLl1ef8D.yDcPz_29xWHwA; _cfuvid=y98uHDayycbQbvAbNhwvWUf390avkf3wmhAK0wE9T_A-1713010726634-0.0.1.1-604800000")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := &Response{}
	err = json.Unmarshal(body, results)
	if err != nil {
		return nil, err
	}

	return results.Data, nil
}
