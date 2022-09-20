package raspberry

import (
	"fmt"
	"testing"

	"github.com/goodaye/fakeeyes/protos/command"
)

var rasp *RaspberryMachine

func init() {

	var err error
	var address = "http://127.0.0.1:5000"
	rasp, err = NewRaspberryDriver(address)
	if err != nil {
		panic(err)
	}
}
func TestHealthCheck(t *testing.T) {
	err := rasp.HealthCheck()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("health ok ")
}

func TestMotion(t *testing.T) {
	op := &command.DeviceOperation{
		Opcode: int32(command.OperateCode_MoveBack),
	}
	err := rasp.Motion(op)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("move ok ")
}
