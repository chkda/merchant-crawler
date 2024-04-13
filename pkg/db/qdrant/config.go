package qdrant

type QdrantConfig struct {
	Host       string `json:"host"`
	Collection string `json:"collection"`
}

const (
	vectorSize = 8
)
