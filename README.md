# go-api
[![Go](https://github.com/valtlfelipe/go-api/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/valtlfelipe/go-api/actions/workflows/go.yml)

This is a simple example of writing an API in Go. I'm using this project to practice my knowledge of Go.

## Getting Started

### Requirements
- Go version 1.22 or higher
- [nodemon](https://www.npmjs.com/package/nodemon) package from npm (for development server)

### Required Environment Variables
- `REDIS_URL`: The URL of the redis instance to connect to.
- `PORT`: The http port the server will listen on. Defaults to `localhost:8090`.

To get started with this API, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/valtlfelipe/go-api.git
```

2. Install dependencies:

```sh
cd go-api
go mod tidy
```

3. Run locally:

```sh
make run
```

The API should now be running at `http://localhost:8090`.

# API Endpoints

_TODO_
