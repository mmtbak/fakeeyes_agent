package machine

import "github.com/goodaye/fakeeyes/protos/request"

type X86LinuxMachine struct{}

func (l X86LinuxMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	return
}
