package tasks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/valtlfelipe/go-api/pkg/db"
)

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaskService struct {
	ctx      context.Context
	dbClient *db.DB
}

func formatPath(id string) string {
	return fmt.Sprintf("task.%v", id)
}

func NewTasksService(ctx context.Context, mux *http.ServeMux, dbClient *db.DB) {
	service := &TaskService{ctx: ctx, dbClient: dbClient}

	mux.HandleFunc("GET /task/{id}", service.getHandler)
}
