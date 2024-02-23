package tasks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) Get(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func (m *MockDBClient) Set(key, value string) {
	m.Called(key, value)
}

func TestGetHandler(t *testing.T) {
	// Create a new instance of the mock store
	mockDB := new(MockDBClient)
	service := TaskService{dbClient: mockDB}

	// Generate a valid UUID for our test
	testUUID := uuid.New()
	testKey := "task." + testUUID.String()
	testValue := "test task"

	// Set up our mock call expectations
	mockDB.On("Get", testKey).Return(testValue)

	// Create a request to pass to our handler with the test UUID
	req, err := http.NewRequest("GET", "/tasks/"+testUUID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// We need to create a router that we can pass our request through so that the vars will be added to the context
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks/{id}", service.getHandler)
	mux.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the body is what we expect
	expected := `{"id":"` + testUUID.String() + `","name":"` + testValue + `"}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
}
