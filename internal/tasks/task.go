package tasks

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/valtlfelipe/go-api/pkg/db"
	"github.com/valtlfelipe/go-api/pkg/httputil"
)

type Task struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TaskService struct {
	dbClient db.DBInterface
}

func formatPath(id string) string {
	return fmt.Sprintf("task.%v", id)
}

func NewTasksService(mux *http.ServeMux, dbClient db.DBInterface) {
	service := &TaskService{dbClient: dbClient}

	httputil.WrapHandler(mux, "GET /tasks/{id}", service.getHandler)
	httputil.WrapHandler(mux, "POST /tasks", service.postHandler)
}
