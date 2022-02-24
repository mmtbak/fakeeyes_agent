package machine

import (
	"os/exec"
	"regexp"

	"github.com/goodaye/fakeeyes/protos/request"
)

// REMacOSStat 编译好的表达式
var REMacOSStat *regexp.Regexp

func init() {

	REMacOSStat = regexp.MustCompile(REPatternMacOSState)
}

var REPatternMacOSState = `\s*(?P<Key>\w.*)\s*:\s*(?P<Value>\w.*)\s+`

type MacMachine struct{}

func (m MacMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	// 硬件信息
	// cmdstr := "/usr/sbin/system_profiler SPHardwareDataType"
	hwcmd := exec.Command("/usr/sbin/system_profiler", "SPHardwareDataType")
	result, err := hwcmd.CombinedOutput()
	if err != nil {
		return
	}
	// systemresult
	matches := REMacOSStat.FindAllStringSubmatch(string(result), -1)
	// fmt.Println(match)

	for _, m := range matches {
		switch m[1] {
		case "Model Name":
			info.ModelName = m[2]
		case "Model Identifier":
			info.ModelID = m[2]
		case "Processor Name":
			info.ProcessorName = m[2]
		case "Processor Speed":
			info.ProcessorSpeed = m[2]
		case "Serial Number (system)":
			info.SN = m[2]
		case "Hardware UUID":
			info.HardwareUUID = m[2]
		}
	}

	// OS信息
	softcmd := exec.Command("/usr/sbin/system_profiler", "SPSoftwareDataType")
	result, err = softcmd.CombinedOutput()
	if err != nil {
		return
	}
	// systemresult
	matches = REMacOSStat.FindAllStringSubmatch(string(result), -1)
	// fmt.Println(matches)
	for _, m := range matches {
		switch m[1] {
		case "Computer Name":
			info.Name = m[2]
		case "System Version":
			info.OSName = m[2]
		case "Kernel Version":
			info.OSVersion = m[2]
		}
	}

	return

}
