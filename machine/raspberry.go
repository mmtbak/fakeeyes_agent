package machine

import (
	"github.com/goodaye/fakeeyes/protos/request"
)

type RaspberryMachine struct{}

func (r RaspberryMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	// cpuinfo, err := ioutil.ReadFile("/proc/cpuinfo")
	// if cpu

	return
}
