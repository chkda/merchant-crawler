package crawler

import (
	"context"
	"os"
	"strings"

	"log/slog"

	"github.com/chkda/merchant-crawler/pkg/api/google"
	"github.com/chkda/merchant-crawler/pkg/db/mysql"
	"github.com/chkda/merchant-crawler/pkg/queue"
)

type Crawler struct {
	SQLConnector   *mysql.SQLConnector
	RabbitMQClient *queue.RabbitMQClient
	SearchAPI      *google.SearchAPI
}

func New(
	sqlConnector *mysql.SQLConnector,
	mqClient *queue.RabbitMQClient,
	searchAPI *google.SearchAPI,
) *Crawler {
	return &Crawler{
		SQLConnector:   sqlConnector,
		RabbitMQClient: mqClient,
		SearchAPI:      searchAPI,
	}
}

func (c *Crawler) Crawl() {
	ctx := context.Background()
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logWriter := slog.New(jsonHandler)
	unmatchedPatterns, err := c.SQLConnector.GetUnmatchedPattern(ctx)
	if err != nil {
		logWriter.Error(err.Error())
		return
	}
	for _, row := range unmatchedPatterns {
		go func() {
			err := c.ProcessPattern(ctx, row.Pattern)
			if err != nil {
				logWriter.Error(err.Error())
				return
			}
			logWriter.Info("Crawled query: " + string(row.Pattern))
		}()
	}

}

func (c *Crawler) ProcessPattern(ctx context.Context, pattern string) error {
	searchResults, err := c.SearchAPI.GetSearchResults(pattern)
	if err != nil {
		return err
	}
	firstResult := searchResults[0]
	merchantLink := firstResult.Link
	merchantTitle := firstResult.Title
	merchantDescription := firstResult.Snippet
	normalisedMerchant := getMerchantNameFromURL(firstResult.DisplayLink)
	msg := &queue.Message{
		Query:              pattern,
		MerchantLink:       merchantLink,
		NormalisedMerchant: normalisedMerchant,
		Title:              merchantTitle,
		Description:        merchantDescription,
	}
	err = c.RabbitMQClient.SendMessage(ctx, msg)
	return err

}

func getMerchantNameFromURL(url string) string {
	if strings.HasPrefix(url, "www.") {
		url = url[4:]
	}
	firstDotIndex := strings.Index(url, ".")
	return url[:firstDotIndex]
}
