package app

import (
	"net/http"

	"github.com/FreeCodeUserJack/bookstore_items/controllers"
)


func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}