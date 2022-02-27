package machine

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/goodaye/fakeeyes/protos/request"
)

type Machine interface {
	CollectDeviceInfo() (request.DeviceInfo, error)
	// 设备类型名字
	// Name() strings
}

func DetectMachine() (m Machine, err error) {

	osname := runtime.GOOS
	arch := runtime.GOARCH
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	return
	// }
	if osname == DevicePlatform.Darwin {
		// 如果是mac电脑
		m = MacMachine{}
	} else if osname == DevicePlatform.Linux && strings.HasPrefix(arch, "arm") {
		// 判断是Arm Linux设备
		m = RaspberryMachine{}
	} else if osname == DevicePlatform.Linux {
		// 判断是普通Linux
		m = X86LinuxMachine{}
	} else {
		err = fmt.Errorf("can't detect machine environment")
	}
	return m, err
}
