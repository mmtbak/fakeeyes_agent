package motion

import "github.com/goodaye/fakeeyes/protos/command"

func Motion(op *command.DeviceOperation) {
	driver.Motion(op)
}
