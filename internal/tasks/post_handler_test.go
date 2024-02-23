package tasks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostHandler(t *testing.T) {
	// Create a new instance of the mock store
	mockDB := new(MockDBClient)
	service := TaskService{dbClient: mockDB}

	testValue := "name-123"

	mockDB.On("Set", mock.Anything, testValue).Return()

	// Create a request to pass to our handler with the test UUID
	req, err := http.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "`+testValue+`"}`))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// We need to create a router that we can pass our request through so that the vars will be added to the context
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", service.postHandler)
	mux.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the body is what we expect
	var result map[string]string

	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, result["name"], testValue)
}
