package vm

import (
	"fmt"

	"github.com/goodaye/fakeeyes/protos/command"
	"github.com/goodaye/fakeeyes/protos/request"
)

var ErrMethodNotImplemented = fmt.Errorf("method not implemented")

type DesktopVM struct{}

func (DesktopVM) Init() error {

	return nil
}
func (DesktopVM) HealthCheck() error {

	return nil
}
func (d DesktopVM) Name() string {
	return "DesktopVM"
}

func (DesktopVM) Motion(*command.DeviceOperation) error {
	return ErrMethodNotImplemented
}
func (DesktopVM) CollectDeviceInfo() (request.DeviceInfo, error) {
	return request.DeviceInfo{}, ErrMethodNotImplemented
}
