package services

import (
	"net/http"

	"github.com/FreeCodeUserJack/bookstore_items/domain/items"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestError)
	Get(string) (*items.Item, rest_errors.RestError)
}

type itemsService struct {
}

func (s *itemsService) Create(itemReq items.Item) (*items.Item, rest_errors.RestError) {
	err := itemReq.Save()
	if err != nil {
		return nil, err
	}
	return &itemReq, nil
}

func (s *itemsService) Get(key string) (*items.Item, rest_errors.RestError) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}