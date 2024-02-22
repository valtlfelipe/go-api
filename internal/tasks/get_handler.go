package tasks

import (
	"encoding/json"
	"net/http"
)

func (service *TaskService) getHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	val := service.dbClient.Get(formatPath(id))

	w.Header().Set("Content-Type", "application/json")

	if val == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(Task{
		Id:   id,
		Name: val,
	})
}
