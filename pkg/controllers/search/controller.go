package search

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/chkda/merchant-crawler/pkg/api/openai"
	"github.com/chkda/merchant-crawler/pkg/db/qdrant"
	"github.com/chkda/merchant-crawler/pkg/db/sql"
	"github.com/labstack/echo/v4"
)

const Route = "/search"

type Controller struct {
	SQLClient   *sql.SQLConnector
	VectorStore *qdrant.Qdrant
	Model       *openai.OpenAIAPI
}

func New(
	sqlClient *sql.SQLConnector,
	vectorStore *qdrant.Qdrant,
	openaiClient *openai.OpenAIAPI,
) *Controller {
	return &Controller{
		SQLClient:   sqlClient,
		VectorStore: vectorStore,
		Model:       openaiClient,
	}
}

func (c *Controller) Handler(e echo.Context) error {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logWriter := slog.New(jsonHandler)
	resp := &Response{}
	query := e.QueryParam("q")
	embeddings, err := c.Model.GetEmbedding(query)
	if err != nil {
		logWriter.Error(err.Error())
		resp.Message = "unable to generate embeddings"
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if len(embeddings) == 0 {
		err = errors.New("could not create embedding")
		logWriter.Error(err.Error())
		resp.Message = "unable to generate embeddings"
		return e.JSON(http.StatusInternalServerError, resp)
	}
	ctx := context.Background()
	embedding := embeddings[0]
	matchedItem, err := c.VectorStore.Search(ctx, embedding.Embedding)
	if err != nil {
		logWriter.Error(err.Error())
		resp.Message = "search failure"
		return e.JSON(http.StatusInternalServerError, resp)
	}

	itemMetaMap := matchedItem.GetPayload()
	merchantName, ok := itemMetaMap["name"]
	if !ok {
		err = errors.New("could not find merchant name in vector payload")
		logWriter.Error(err.Error())
		resp.Message = "search failure"
		return e.JSON(http.StatusInternalServerError, resp)
	}
	resp.MerchantName = merchantName.GetStringValue()
	return e.JSON(http.StatusOK, resp)
}
