package client

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/goodaye/fakeeyes/protos/command"
)

func TestClient(t *testing.T) {
	addr := "http://127.0.0.1:5000"
	c, err := NewClient(addr)
	assert.Equal(t, err, nil)
	err = c.HealthCheck()
	assert.Equal(t, err, nil)

	op := &command.DeviceOperation{
		Opcode:    1,
		Opvalue:   2,
		Opmessage: "test value ",
	}
	err = c.Motion(op)
	assert.Equal(t, err, nil)

}
