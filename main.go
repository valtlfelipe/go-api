package main

import (
	"context"
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

	err = http.ListenAndServe("localhost:8090", mux)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Running on :8090")
}
