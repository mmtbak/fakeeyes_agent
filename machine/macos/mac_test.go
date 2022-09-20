package macos

import (
	"fmt"
	"testing"
)

func TestMacMachine(t *testing.T) {

	var m MacOSDriver
	var err error
	m = MacOSDriver{}
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
