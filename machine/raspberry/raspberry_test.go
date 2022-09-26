package raspberry

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestParse(t *testing.T) {

	var cpuinfo = `Hardware	: BCM2711
	Revision	: c03112
	Serial		: 10000000f0b4fc60
	Model		: Raspberry Pi 4 Model B Rev 1.2`
	fmt.Println(cpuinfo)
}

func init() {
	driverpath = "../../drivers"
}

func TestMontion(t *testing.T) {

	var err error
	wd, err := os.Getwd()
	assert.Equal(t, err, nil)

	fmt.Println("wd :", wd)
	m := RaspberryMachine{}
	go func() {
		err = m.StartLocalServer()
		if err != nil {
			fmt.Println(err)
		}
	}()
	time.Sleep(1 * time.Second)
	err = m.client.HealthCheck()
	assert.Equal(t, err, nil)
}
