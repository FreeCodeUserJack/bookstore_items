package controllers

import (
	"fmt"
	"net/http"

	"github.com/FreeCodeUserJack/bookstore_items/domain/items"
	"github.com/FreeCodeUserJack/bookstore_items/services"
	"github.com/FreeCodeUserJack/bookstore_oauth-common/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error json to the user
		return
	}
	item := items.Item {
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: Return error json to the user
	}

	fmt.Println(result)

	// TODO: return created item as JSON with http status 201 created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}