package drivers

import "github.com/goodaye/fakeeyes/protos/command"

type Motion interface {
	Init() error
	HealthCheck() error
	Motion(*command.DeviceOperation) error
}
