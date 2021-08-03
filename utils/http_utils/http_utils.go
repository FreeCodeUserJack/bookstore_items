package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/FreeCodeUserJack/bookstore_utils/rest_errors"
)


func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err rest_errors.RestError) {
	ResponseJson(w, err.Status(), err)
}