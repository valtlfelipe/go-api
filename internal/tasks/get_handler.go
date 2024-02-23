package tasks

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/valtlfelipe/go-api/pkg/httputil"
)

func (service *TaskService) getHandler(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.RespondError(w, httputil.ResponseError{
			Error:  "invalid uuid",
			Status: http.StatusBadRequest,
		})
		return
	}

	// service.dbClient.Set(formatPath(id.String()), "teste-123")

	val := service.dbClient.Get(formatPath(id.String()))

	if val == "" {
		httputil.RespondError(w, httputil.ResponseError{
			Error:  "not found",
			Status: http.StatusNotFound,
		})
		return
	}

	httputil.RespondSuccess(w, Task{
		Id:   id,
		Name: val,
	})
}
