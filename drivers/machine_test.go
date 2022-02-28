package drivers

import (
	"fakeeyes_agent/drivers/macos"
	"fmt"
	"testing"
)

func TestMacMachine(t *testing.T) {

	var m Machine
	var err error
	m = macos.MacOSDriver{}
	err = m.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	info, err := m.CollectDeviceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)

}
