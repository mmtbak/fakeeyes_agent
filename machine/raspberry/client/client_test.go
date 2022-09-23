package client

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClient(t *testing.T) {
	addr := "http://127.0.0.1:5000"
	c, err := NewClient(addr)
	assert.Equal(t, err, nil)
	err = c.HealthCheck()
	assert.Equal(t, err, nil)

}
