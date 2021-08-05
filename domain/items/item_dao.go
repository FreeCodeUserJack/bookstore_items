package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/FreeCodeUserJack/bookstore_items/clients/elasticsearch"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
)


const (
	indexItems = "items"
	typeItem = "_doc"
)

func (i *Item) Save() rest_errors.RestError {
	result, err := elasticsearch.ESClient.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	// update item id
	i.Id = result.Id
	return nil
}

func (i *Item) Get() (rest_errors.RestError) {
	itemId := i.Id

	result, err := elasticsearch.ESClient.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}

	bytes, jsonErr := result.Source.MarshalJSON()
	if jsonErr != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	
	i.Id = itemId
	return nil
}