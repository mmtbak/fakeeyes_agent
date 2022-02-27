package controller

import (
	"fmt"
	"testing"

	client "github.com/goodaye/fakeeyes_client_golang"
	"github.com/goodaye/wire"
)

func init() {
	var err error
	fesclient, err = client.NewClient("http://127.0.0.1:8080/")
	if err != nil {
		panic(err)
	}
	wire.Append(SVC{})
	err = wire.Init()
	if err != nil {
		panic(err)
	}
}

func TestDeviceRegister(t *testing.T) {

	// devinfo := request.DeviceInfo{
	// 	SN: "testdevice",
	// }
	info, err := localmachine.CollectDeviceInfo()
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
