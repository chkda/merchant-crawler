package main

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/chkda/merchant-crawler/pkg/api/google"
	"github.com/chkda/merchant-crawler/pkg/crawler"
	"github.com/chkda/merchant-crawler/pkg/db/sql"
	"github.com/chkda/merchant-crawler/pkg/queue"
	"github.com/go-co-op/gocron"
)

type Config struct {
	SearchAPIConfig *google.APIConfig     `json:"google_custom_search"`
	SQLConnConfig   *sql.SQLConnConfig    `json:"sql_client"`
	RabbitMQConfig  *queue.RabbitMQConfig `json:"rabbitmq_client"`
}

var FILE_LOC = "/home/chhaya/my_files/ml/fold/crawler/config/crawler/config.json"

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

	searchAPI := google.New(cfg.SearchAPIConfig)

	crawler := crawler.New(sqlClient, rabbitmqClient, searchAPI)
	// crawler.Crawl()
	cron := gocron.NewScheduler(time.Local)
	_, err = cron.Cron("*/1 * * * *").Do(crawler.Crawl)
	if err != nil {
		panic(err)
	}
	cron.StartAsync()
	block := make(chan struct{})
	<-block

}
