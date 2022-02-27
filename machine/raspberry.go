package machine

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/goodaye/fakeeyes/protos/request"
)

// RaspberryOSREStat  MacOS RE State MacOS 正则表达式 编译好的表达式
var RaspberryOSREStat *regexp.Regexp
var RaspberryOSStatPattern = `\s*(?P<Key>\w.*)\s*:\s*(?P<Value>\w.*)\s+`

func init() {
	RaspberryOSREStat = regexp.MustCompile(RaspberryOSStatPattern)
}

type RaspberryMachine struct{}

func (r RaspberryMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	cpuinfo, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return
	}

	// systemresult
	matches := RaspberryOSREStat.FindAllStringSubmatch(string(cpuinfo), -1)
	// fmt.Println(match)

	for _, m := range matches {
		switch m[1] {
		case "Serial":
			info.SN = m[2]
		case "Model":
			info.ModelName = m[2]
		case "Serial Number (system)":
			info.SN = m[2]
		}
	}
	// 获得CPU信息
	cpucmd := exec.Command("lscpu")
	result, err := cpucmd.CombinedOutput()
	if err != nil {
		return
	}
	// cpu regex
	matches = RaspberryOSREStat.FindAllStringSubmatch(string(result), -1)

	for _, m := range matches {
		switch m[1] {
		case "Architectur":
			info.CPUArch = m[2]
		case "CPU(s)":
			var val int
			val, err = strconv.Atoi(m[2])
			if err != nil {
				return
			}
			info.CPUCores = val
		case "Socket(s)":
			var val int
			val, err = strconv.Atoi(m[2])
			if err != nil {
				return
			}
			info.CPUSocket = val
		case "Vendor ID":
			info.CPUVendor = m[2]

		case "Model":
			info.CPUModelID = m[2]
		case "Model name":
			info.CPUModel = m[2]
		case "CPU max MHz":
			info.CPUSpeed = m[2]

		}
	}

	return
}
