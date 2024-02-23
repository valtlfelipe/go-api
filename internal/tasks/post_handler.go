package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/valtlfelipe/go-api/pkg/httputil"
)

func (service *TaskService) postHandler(w http.ResponseWriter, r *http.Request) {
	var t Task

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		httputil.RespondError(w, http.StatusBadRequest, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	if t.Name == "" {
		httputil.RespondError(w, http.StatusBadRequest, "name is required")
		return
	}

	id := uuid.New()

	service.dbClient.Set(formatPath(id.String()), t.Name)

	httputil.RespondSuccess(w, Task{
		Id:   id,
		Name: t.Name,
	})
}
