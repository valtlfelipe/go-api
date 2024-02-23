init:
	go mod tidy

run:
	nodemon --exec "go run" ./cmd/main.go --signal SIGTERM

test:
	go test -v ./...

build:
	go build cmd/main.go
