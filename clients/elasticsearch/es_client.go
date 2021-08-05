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
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	olivereClient, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(logger.GetLogger()),
		elastic.SetInfoLog(logger.GetLogger()),
	)

	if err != nil {
		panic(err)
	}

	ESClient = &esClient{
		client: olivereClient,
	}
}

func (c *esClient) Index(index, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()

	result, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Get(index, docType, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}

	return result, nil
}