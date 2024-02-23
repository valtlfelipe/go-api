init:
	go mod tidy

run:
	nodemon --exec "go run" ./cmd/main.go --signal SIGTERM

test:
	go test -v ./...

build:
	go build cmd/main.go

test-coverage:
	mkdir -p coverage/
	go test -v ./... -coverprofile coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html

lint:
	golangci-lint run ./...
