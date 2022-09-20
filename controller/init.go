package controller

import (
	"fakeeyes_agent/config"
	"fakeeyes_agent/machine"

	client "github.com/goodaye/fakeeyes_client_golang"
	"github.com/goodaye/wire"
)

var fesclient *client.Client

var deviceclient *client.Device

var localmachine machine.Machine

type SVC struct {
	wire.BaseService
}

func (s SVC) Init() error {
	var err error
	fesclient, err = client.NewClient(config.GlobalConfig.Fakeeyes.Server)
	if err != nil {
		return err
	}
	localmachine, err = DetectMachine()
	if err != nil {
		return err
	}

	return nil
}

func (s SVC) Start() error {
	var err error
	deviceclient, err = Register()
	if err != nil {
		return err
	}
	return Start()
}
