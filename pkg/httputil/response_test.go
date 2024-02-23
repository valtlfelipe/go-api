package httputil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	data := map[string]string{"foo": "bar"}
	RespondSuccess(w, data)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, string(body), "{\"foo\":\"bar\"}\n")
}

func TestRespondError(t *testing.T) {

	w := httptest.NewRecorder()
	RespondError(w, http.StatusBadRequest, "invalid uuid")

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, string(body), "{\"error\":\"invalid uuid\"}\n")
}
