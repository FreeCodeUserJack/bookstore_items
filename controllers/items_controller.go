package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FreeCodeUserJack/bookstore_items/domain/items"
	"github.com/FreeCodeUserJack/bookstore_items/services"
	"github.com/FreeCodeUserJack/bookstore_items/utils/http_utils"
	"github.com/FreeCodeUserJack/bookstore_oauth-common/oauth"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
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
		http_utils.ResponseError(w, err)
		return
	}

	var itemRequest items.Item

	reqBody, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.ResponseError(w, respErr)
		return
	}
	defer r.Body.Close()

	jsonErr := json.Unmarshal(reqBody, &itemRequest)
	if jsonErr != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json")
		http_utils.ResponseError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, err := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}