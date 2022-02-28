package vm

import (
	"fmt"
	"testing"
)

func TestDesktopDriver(t *testing.T) {

	var d = DesktopVM{}
	info, err := d.CollectDeviceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
