package machine

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/goodaye/fakeeyes/protos/request"
)

// MacOSREStat  MacOS RE State MacOS 正则表达式 编译好的表达式
var MacOSREStat *regexp.Regexp
var MacOSStatPattern = `\s*(?P<Key>\w.*)\s*:\s*(?P<Value>\w.*)\s+`

func init() {
	MacOSREStat = regexp.MustCompile(MacOSStatPattern)
}

type MacMachine struct{}

func (m MacMachine) CollectDeviceInfo() (info request.DeviceInfo, err error) {

	// 硬件信息
	// cmdstr := "/usr/sbin/system_profiler SPHardwareDataType"
	hwcmd := exec.Command("/usr/sbin/system_profiler", "SPHardwareDataType", "-json")
	result, err := hwcmd.CombinedOutput()
	if err != nil {
		return
	}
	var cmdoutputerr = fmt.Errorf("command output error ")
	var hwinfo = map[string]interface{}{}
	err = json.Unmarshal(result, &hwinfo)
	if err != nil {
		return
	}
	hwresult, ok := hwinfo["SPHardwareDataType"]
	if !ok {
		err = cmdoutputerr
		return
	}
	hwdata, ok := hwresult.([]interface{})
	if !ok {
		err = cmdoutputerr
		return
	}
	if len(hwdata) == 0 {
		err = cmdoutputerr
		return
	}
	hwdatamap := hwdata[0].(map[string]interface{})

	for key, value := range hwdatamap {
		switch key {
		case "cpu_type":
			info.CPUModel = value.(string)
		case "current_processor_speed":
			info.CPUSpeed = value.(string)
		case "machine_name":
			info.ModelName = value.(string)
		case "machine_model":
			info.ModelID = value.(string)
		case "number_processors":
			info.CPUCores = int(value.(float64))
		case "platform_UUID":
			info.HardwareUUID = value.(string)
		case "serial_number":
			info.SN = value.(string)
		}
	}

	// // systemresult
	// matches := MacOSREStat.FindAllStringSubmatch(string(result), -1)
	// // fmt.Println(match)

	// for _, m := range matches {
	// 	switch m[1] {
	// 	case "Model Name":
	// 		info.ModelName = m[2]
	// 	case "Model Identifier":
	// 		info.ModelID = m[2]
	// 	case "Processor Name":
	// 		info.CPUModel = m[2]
	// 	case "Processor Speed":
	// 		info.CPUSpeed = m[2]
	// 	case "Serial Number (system)":
	// 		info.SN = m[2]
	// 	case "Hardware UUID":
	// 		info.HardwareUUID = m[2]
	// 	}
	// }

	// OS信息
	softcmd := exec.Command("/usr/sbin/system_profiler", "SPSoftwareDataType", "-json")
	result, err = softcmd.CombinedOutput()
	if err != nil {
		return
	}

	var osinfo = map[string]interface{}{}
	err = json.Unmarshal(result, &osinfo)
	if err != nil {
		return
	}
	osresult, ok := osinfo["SPSoftwareDataType"]
	if !ok {
		err = cmdoutputerr
		return
	}
	osdata, ok := osresult.([]interface{})
	if !ok {
		err = cmdoutputerr
		return
	}
	if len(osdata) == 0 {
		err = cmdoutputerr
		return
	}
	osdatamap := osdata[0].(map[string]interface{})
	for key, value := range osdatamap {
		switch key {
		case "kernel_version":
			info.OSVersion = value.(string)
		case "local_host_name":
			info.Name = value.(string)
		case "os_version":
			info.OSName = value.(string)
			// case "uptime":
			// 	预留
		}
	}

	// // systemresult
	// matches = MacOSREStat.FindAllStringSubmatch(string(result), -1)
	// // fmt.Println(matches)
	// for _, m := range matches {
	// 	switch m[1] {
	// 	case "Computer Name":
	// 		info.Name = m[2]
	// 	case "System Version":
	// 		info.OSName = m[2]
	// 	case "Kernel Version":
	// 		info.OSVersion = m[2]
	// 	}
	// }
	return

}

func (m MacMachine) Name() {
}
