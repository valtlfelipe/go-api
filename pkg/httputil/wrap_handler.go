package httputil

import (
	"net/http"

	"github.com/valtlfelipe/go-api/pkg/httputil/middlewares"
)

// WrapHandler takes an http.ServeMux, a path, and an http.HandlerFunc,
// then applies a set of middlewares to the handler before registering it
// to the ServeMux on the specified path.
func WrapHandler(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	mux.Handle(path,
		middlewares.LogRequest(
			middlewares.EnforceJSON(
				http.HandlerFunc(handler))))
}
