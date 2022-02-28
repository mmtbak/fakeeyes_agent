package motion

import (
	"fakeeyes_agent/drivers"

	"github.com/goodaye/fakeeyes/protos/command"
)

func Motion(m drivers.Machine, op *command.DeviceOperation) {
	m.Motion(op)
}
