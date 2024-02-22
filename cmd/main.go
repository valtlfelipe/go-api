package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/valtlfelipe/go-api/internal/tasks"
	"github.com/valtlfelipe/go-api/pkg/config"
	"github.com/valtlfelipe/go-api/pkg/db"
)

var ctx = context.Background()

func main() {
	config := config.NewConfig()

	mux := http.NewServeMux()

	// setup redis client
	redisOpt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(redisOpt)
	dbClient := db.NewDB(ctx, redisClient)

	// setup routes and handlers
	tasks.NewTasksService(mux, dbClient)

	// start the http server
	log.Printf("Starting on %s ...\n", config.Port)

	srv := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	// Create a context that listens for the SIGINT and SIGTERM signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for the context to be canceled (when SIGINT or SIGTERM is received)
	<-ctx.Done()

	// Create a deadline for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %s\n", err)
		os.Exit(1)
	}

	log.Println("Server shutdown gracefully")
}
