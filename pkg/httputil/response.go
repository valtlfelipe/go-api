package httputil

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error  string `json:"error"`
	Status int    `json:"-"`
}

func RespondSuccess(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func RespondError(w http.ResponseWriter, result ResponseError) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(result.Status)

	json.NewEncoder(w).Encode(result)
}
