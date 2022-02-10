package heartbeat

import (
	"fmt"
	"testing"
)

func TestCollectHost(t *testing.T) {
	info, err := CollectDeviceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}

func TestCollectMacOSStat(t *testing.T) {

	info, err := CollectDeviceStat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)

}
