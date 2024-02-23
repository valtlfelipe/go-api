run:
	nodemon --exec "go run" ./cmd/main.go --signal SIGTERM

test:
	go test -v ./...
