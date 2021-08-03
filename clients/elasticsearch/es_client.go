package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/FreeCodeUserJack/bookstore_utils/logger"
	"github.com/olivere/elastic/v7"
)


var (
	ESClient esClientInterface
)

type esClientInterface interface {
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	olivereClient, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	ESClient = &esClient{
		client: olivereClient,
	}
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()

	result, err := c.client.Index().Index(index).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}