package embeddings

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/chkda/merchant-crawler/pkg/api/openai"
	"github.com/chkda/merchant-crawler/pkg/db/qdrant"
)

type EmbeddingsGenerator struct {
	Model       *openai.OpenAIAPI
	VectorStore *qdrant.Qdrant
}

func New(
	client *openai.OpenAIAPI,
	vectorStore *qdrant.Qdrant,
) *EmbeddingsGenerator {
	return &EmbeddingsGenerator{
		Model:       client,
		VectorStore: vectorStore,
	}
}

func (c *EmbeddingsGenerator) ProcessText(
	ctx context.Context,
	text string,
	merchantName string,
	merchantLink string,
) {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logWriter := slog.New(jsonHandler)
	embeddings, err := c.Model.GetEmbedding(text)
	if err != nil {
		logWriter.Error(err.Error())
		return
	}

	if len(embeddings) == 0 {
		err = errors.New("could not create embedding")
		logWriter.Error(err.Error())
		return
	}

	embedding := embeddings[0]
	item := &qdrant.QdrantItem{
		NormalisedMerchant: merchantName,
		Link:               merchantLink,
		Vector:             embedding.Embedding,
	}
	err = c.VectorStore.Upsert(ctx, item)
	if err != nil {
		logWriter.Error(err.Error())
		return
	}
}
