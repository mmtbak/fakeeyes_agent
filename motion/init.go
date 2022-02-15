package motion

import (
	"fakeeyes_agent/drivers"
	"fakeeyes_agent/drivers/raspberry"

	"github.com/goodaye/wire"
)

var driver drivers.Motion

type SVC struct {
	wire.BaseService
}

func (s SVC) Init() error {

	var err error
	driver, err = raspberry.NewRaspberry("http://127.0.0.1:5000")
	if err != nil {
		return err
	}
	err = driver.Init()
	return err
}

func (s SVC) Start() error {
	return nil
}
