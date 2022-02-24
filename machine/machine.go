package machine

import (
	"fmt"
	"os"
	"runtime"

	"github.com/goodaye/fakeeyes/protos/request"
)

type Machine interface {
	CollectDeviceInfo() (request.DeviceInfo, error)
}

func DetectMachine() (m Machine, err error) {

	osname := runtime.GOOS
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	if osname == DevicePlatform.Darwin {
		m = MacMachine{}
	} else if osname == DevicePlatform.Linux && hostname == "raspberry" {
		m = RaspberryMachine{}
	} else if osname == DevicePlatform.Linux {
		m = X86LinuxMachine{}
	} else {
		err = fmt.Errorf("can't detect machine environment")
	}
	return m, err
}
