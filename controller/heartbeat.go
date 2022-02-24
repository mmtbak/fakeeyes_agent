package controller

import (
	"fmt"
	"time"
)

func SendHeartBeat() {

	// var err error
	for {
		time.Sleep(10 * time.Minute)
		devinfo, err := localmachine.CollectDeviceInfo()
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
