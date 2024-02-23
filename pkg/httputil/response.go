package httputil

import (
	"encoding/json"
	"net/http"
)

// RespondSuccess sends a JSON response with a 200 OK status code to the client.
// The response parameter is marshaled into JSON and written to the response writer.
// It sets the Content-Type header to "application/json".
func RespondSuccess(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

// RespondError sends a JSON response with the specified status code to the client.
// The error message is marshaled into a JSON object with an "error" key.
// It sets the Content-Type header to "application/json".
func RespondError(w http.ResponseWriter, status int, error string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{"error": error})
}
