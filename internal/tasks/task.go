package tasks

import (
	"fmt"
	"net/http"

	"github.com/valtlfelipe/go-api/pkg/db"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaskService struct {
	dbClient *db.DB
}

func formatPath(id string) string {
	return fmt.Sprintf("task.%v", id)
}

func NewTasksService(mux *http.ServeMux, dbClient *db.DB) {
	service := &TaskService{dbClient: dbClient}

	mux.HandleFunc("GET /task/{id}", service.getHandler)
}
