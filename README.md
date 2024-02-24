# go-api
[![go-test](https://github.com/valtlfelipe/go-api/actions/workflows/go-test.yml/badge.svg?branch=main)](https://github.com/valtlfelipe/go-api/actions/workflows/go-test.yml)

This is a simple example of writing an API in Go. I'm using this project to practice my knowledge of Go.

## Getting Started

### Requirements
- Go version 1.22 or higher
- [nodemon](https://www.npmjs.com/package/nodemon) package from npm (for development server)
- Redis server. [Upstash](https://upstash.com/) is free 😉

### Environment Variables
- `REDIS_URL`: (required) The URL of the redis instance to connect to.
- `PORT`: The http port the server will listen on. Defaults to `localhost:8090`.

To get started with this repo, follow these steps:

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

4. Run tests:

```sh
make test
```

5. Test coverage:

```sh
make test-coverage
# open coverage/coverage.html
```

Test coverage files will be available inside `coverage/` folder.

# API Endpoints

## Create a Task

### Request

- URL: `/tasks`
- Method: `POST`
- Headers:
  - Content-Type: application/json

#### Body

```json
{
  "name": "Task Title"
}
```

## Get a Task by Id

### Request

- URL: `/tasks/{id}`
- Method: `GET`

### Response

```json
{
  "id": "578f2d5e-c4cc-4ea8-b98d-47690e8bb6a5",
  "name": "Task Title"
}
```
