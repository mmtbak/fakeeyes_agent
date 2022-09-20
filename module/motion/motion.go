package motion

import (
	"fakeeyes_agent/machine"

	"github.com/goodaye/fakeeyes/protos/command"
)

func Motion(m machine.Machine, op *command.DeviceOperation) {
	m.Motion(op)
}
