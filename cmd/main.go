package main

import (
	"context"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/valtlfelipe/go-api/internal/tasks"
	"github.com/valtlfelipe/go-api/pkg/config"
	"github.com/valtlfelipe/go-api/pkg/db"
)

var ctx = context.Background()

func main() {
	config := config.NewConfig()

	mux := http.NewServeMux()

	opt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	dbClient := db.NewDB(ctx, client)

	tasks.NewTasksService(ctx, mux, dbClient)

	log.Println("Starting on :8090 ...")

	if err := http.ListenAndServe("localhost:8090", mux); err != nil {
		panic(err)
	}
}
