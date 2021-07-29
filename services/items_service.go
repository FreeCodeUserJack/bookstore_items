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

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestError) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestError) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}