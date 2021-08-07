package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/FreeCodeUserJack/bookstore_items/domain/items"
	"github.com/FreeCodeUserJack/bookstore_items/domain/queries"
	"github.com/FreeCodeUserJack/bookstore_items/services"
	"github.com/FreeCodeUserJack/bookstore_items/utils/http_utils"
	"github.com/FreeCodeUserJack/bookstore_oauth-common/oauth"
	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.ResponseError(w, respErr)
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

	itemRequest.Seller = sellerId

	result, err := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.ResponseError(w, err)
		return
	}

	http_utils.ResponseJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.ResponseError(w, err)
	}

	http_utils.ResponseJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.ResponseError(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if err != nil {
		http_utils.ResponseError(w, searchErr)
		return
	}

	http_utils.ResponseJson(w, http.StatusOK, items)
}