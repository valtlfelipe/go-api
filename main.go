package main

import (
	"context"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/valtlfelipe/go-api/src/config"
	"github.com/valtlfelipe/go-api/src/db"
	"github.com/valtlfelipe/go-api/src/tasks"
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
