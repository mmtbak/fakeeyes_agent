package controller

import (
	"fakeeyes_agent/heartbeat"
	"fakeeyes_agent/motion"
	"fmt"
	"time"

	"github.com/goodaye/fakeeyes/protos/command"
	client "github.com/goodaye/fakeeyes_client_golang"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func Register() (dev *client.Device, err error) {

	devinfo, err := heartbeat.CollectDeviceStat()
	if err != nil {
		return
	}
	dev, err = fesclient.RegisterDevice(devinfo)
	return
}

func SendHeartBeat() {

	// var err error
	for {
		time.Sleep(10 * time.Minute)
		devinfo, err := heartbeat.CollectDeviceStat()
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = deviceclient.SendHeartBeat(devinfo)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Start() error {
	go SendHeartBeat()
	err := Connect()
	return err
}

func Connect() (err error) {
	conn, err := deviceclient.Connect()
	if err != nil {
		return
	}

	defer func() {
		conn.Close()
	}()
	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		switch mt {
		case websocket.TextMessage:
			PrintAction(p)
		case websocket.BinaryMessage:
			Motion(p)
		default:
			continue
		}
	}
	return nil
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
	motion.Motion(&op)
}
