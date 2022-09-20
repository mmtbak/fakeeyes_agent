package machine

import (
	"github.com/goodaye/fakeeyes/protos/command"
	"github.com/goodaye/fakeeyes/protos/request"
)

type Machine interface {
	Init() error
	Motion(*command.DeviceOperation) error
	CollectDeviceInfo() (request.DeviceInfo, error)
	Name() string
}
