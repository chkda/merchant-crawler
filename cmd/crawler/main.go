package main

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/chkda/merchant-crawler/pkg/api/google"
	"github.com/chkda/merchant-crawler/pkg/controllers/healthcheck"
	"github.com/chkda/merchant-crawler/pkg/crawler"
	"github.com/chkda/merchant-crawler/pkg/db/sql"
	"github.com/chkda/merchant-crawler/pkg/queue"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

type Config struct {
	HTTPPort        string                `json:"http_port"`
	SearchAPIConfig *google.APIConfig     `json:"google_custom_search"`
	SQLConnConfig   *sql.SQLConnConfig    `json:"sql_client"`
	RabbitMQConfig  *queue.RabbitMQConfig `json:"rabbitmq_client"`
}

var FILE_LOC = "/config/crawler/config.json"

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

	rabbitmqClient, err := queue.New(cfg.RabbitMQConfig)
	if err != nil {
		panic(err)
	}

	searchAPI := google.New(cfg.SearchAPIConfig)

	crawler := crawler.New(sqlClient, rabbitmqClient, searchAPI)
	// crawler.Crawl()
	cron := gocron.NewScheduler(time.Local)
	_, err = cron.Cron("*/1 * * * *").Do(crawler.Crawl)
	if err != nil {
		panic(err)
	}
	cron.StartAsync()
	// block := make(chan struct{})
	// <-block
	healthcheckController := healthcheck.New()
	serv := echo.New()
	serv.GET(healthcheck.Route, healthcheckController.Handler)
	serv.Logger.Fatal(serv.Start(":" + cfg.HTTPPort))

}
