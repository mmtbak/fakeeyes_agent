package controller

import (
	"fakeeyes_agent/heartbeat"
	"fmt"
	"testing"

	client "github.com/goodaye/fakeeyes_client_golang"
)

func init() {
	fesclient, _ = client.NewClient("http://127.0.0.1:8080/")
}

func TestDeviceRegister(t *testing.T) {

	// devinfo := request.DeviceInfo{
	// 	SN: "testdevice",
	// }
	info, err := heartbeat.CollectDeviceStat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
	devcie, err := fesclient.RegisterDevice(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(devcie)
}
