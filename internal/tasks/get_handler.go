package tasks

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/valtlfelipe/go-api/pkg/httputil"
)

func (service *TaskService) getHandler(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		httputil.RespondError(w, http.StatusBadRequest, "invalid uuid")
		return
	}

	val := service.dbClient.Get(formatPath(id.String()))

	if val == "" {
		httputil.RespondError(w, http.StatusNotFound, "not found")
		return
	}

	httputil.RespondSuccess(w, Task{
		Id:   id,
		Name: val,
	})
}
