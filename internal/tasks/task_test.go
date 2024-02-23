package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatPath(t *testing.T) {
	path := formatPath("123")

	assert.Equal(t, path, "task.123")
}
