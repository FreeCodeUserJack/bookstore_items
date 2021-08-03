package items

import (
	"errors"

	"github.com/FreeCodeUserJack/bookstore_items/clients/elasticsearch"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
)


const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestError {
	result, err := elasticsearch.ESClient.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	// update item id
	i.Id = result.Id
	return nil
}