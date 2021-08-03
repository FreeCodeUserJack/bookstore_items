package app

import (
	"net/http"
	"time"

	"github.com/FreeCodeUserJack/bookstore_items/clients/elasticsearch"
	"github.com/gorilla/mux"
)


var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}