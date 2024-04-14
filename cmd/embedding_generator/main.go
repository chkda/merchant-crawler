package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/chkda/merchant-crawler/pkg/api/openai"
	"github.com/chkda/merchant-crawler/pkg/controllers/healthcheck"
	"github.com/chkda/merchant-crawler/pkg/db/qdrant"
	"github.com/chkda/merchant-crawler/pkg/db/sql"
	"github.com/chkda/merchant-crawler/pkg/embeddings"
	"github.com/chkda/merchant-crawler/pkg/queue"
	"github.com/labstack/echo/v4"
)

type Config struct {
	HTTPPort       string                `json:"http_port"`
	OpenAIConfig   *openai.OpenAIConfig  `json:"openai"`
	SQLConnConfig  *sql.SQLConnConfig    `json:"sql_client"`
	RabbitMQConfig *queue.RabbitMQConfig `json:"rabbitmq_client"`
	QdrantConfig   *qdrant.QdrantConfig  `json:"qdrant_config"`
}

var FILE_LOC = "/home/chhaya/my_files/ml/fold/crawler/config/embedding_generator/config.json"

func main() {
	configFile, err := os.Open(FILE_LOC)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	fileBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	cfg := &Config{}
	err = json.Unmarshal(fileBytes, cfg)
	if err != nil {
		panic(err)
	}

	sqlClient, err := sql.New(cfg.SQLConnConfig)
	if err != nil {
		panic(err)
	}

	rabbitmqClient, err := queue.New(cfg.RabbitMQConfig)
	if err != nil {
		panic(err)
	}

	openAIClient := openai.New(cfg.OpenAIConfig)

	qdrantClient, err := qdrant.New(cfg.QdrantConfig)
	if err != nil {
		panic(err)
	}

	generator := embeddings.New(openAIClient, qdrantClient, sqlClient)

	ctx := context.Background()
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logWriter := slog.New(jsonHandler)

	msgChan, err := rabbitmqClient.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	go func() {
		for msg := range msgChan {
			merchantInfo := &queue.Message{}
			err = json.Unmarshal(msg.Body, merchantInfo)
			if err != nil {
				logWriter.Error(err.Error())
				continue
			}
			generator.ProcessText(ctx,
				merchantInfo.NormalisedMerchant,
				merchantInfo.MerchantLink,
				merchantInfo.ID)
			logWriter.Info("Message processed")

		}
	}()
	// block := make(chan struct{})
	// <-block
	healthcheckController := healthcheck.New()
	serv := echo.New()
	serv.GET(healthcheck.Route, healthcheckController.Handler)
	serv.Logger.Fatal(serv.Start(":" + cfg.HTTPPort))
}
