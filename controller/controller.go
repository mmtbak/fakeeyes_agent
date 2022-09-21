package controller

import (
	"fakeeyes_agent/module/motion"
	"fmt"

	"github.com/goodaye/fakeeyes/protos/command"
	client "github.com/goodaye/fakeeyes_client_golang"
	"google.golang.org/protobuf/proto"
)

func Register() (dev *client.Device, err error) {

	devinfo, err := localmachine.CollectDeviceInfo()
	if err != nil {
		return
	}
	dev, err = fesclient.RegisterDevice(devinfo)
	return
}

func Start() error {
	go SendHeartBeat()
	err := Connect()
	return err
}

func PrintAction(p []byte) {

	fmt.Println(string(p))
}

func Motion(p []byte) {
	var op command.DeviceOperation

	err := proto.Unmarshal(p, &op)
	if err != nil {
		return
	}
	motion.Motion(localmachine, &op)
}
