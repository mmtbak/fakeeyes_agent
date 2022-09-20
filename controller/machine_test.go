package controller

import (
	"fakeeyes_agent/machine/macos"
	"fmt"
	"runtime"
	"testing"
)

func TestGOOS(t *testing.T) {
	osname := runtime.GOOS
	fmt.Println(osname)
}

func TestMacMachine(t *testing.T) {
	mac := macos.MacOSDriver{}
	info, err := mac.CollectDeviceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
