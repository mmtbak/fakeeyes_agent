package raspberry

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	var cpuinfo = `Hardware	: BCM2711
	Revision	: c03112
	Serial		: 10000000f0b4fc60
	Model		: Raspberry Pi 4 Model B Rev 1.2`
	fmt.Println(cpuinfo)
}
