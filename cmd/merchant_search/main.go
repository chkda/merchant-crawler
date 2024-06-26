package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/chkda/merchant-crawler/pkg/api/openai"
	"github.com/chkda/merchant-crawler/pkg/controllers/healthcheck"
	"github.com/chkda/merchant-crawler/pkg/controllers/search"
	"github.com/chkda/merchant-crawler/pkg/db/qdrant"
	"github.com/chkda/merchant-crawler/pkg/db/sql"
	"github.com/labstack/echo/v4"
)

type Config struct {
	HTTPPort      string               `json:"http_port"`
	OpenAIConfig  *openai.OpenAIConfig `json:"openai"`
	SQLConnConfig *sql.SQLConnConfig   `json:"sql_client"`
	QdrantConfig  *qdrant.QdrantConfig `json:"qdrant_config"`
}

var FILE_LOC = "/config/merchant_search/config.json"

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile, err := os.Open(currDir + FILE_LOC)
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

	openAIClient := openai.New(cfg.OpenAIConfig)

	qdrantClient, err := qdrant.New(cfg.QdrantConfig)
	if err != nil {
		panic(err)
	}

	searchController := search.New(sqlClient, qdrantClient, openAIClient)
	healthcheckController := healthcheck.New()
	serv := echo.New()
	serv.GET(search.Route, searchController.Handler)
	serv.GET(healthcheck.Route, healthcheckController.Handler)
	serv.Logger.Fatal(serv.Start(":" + cfg.HTTPPort))

}
