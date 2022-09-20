package x86linux

import (
	"fakeeyes_agent/machine/vm"

	"github.com/goodaye/fakeeyes/protos/request"
)

type X86LinuxMachine struct {
	vm.DesktopVM
}

func (l X86LinuxMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	return
}
