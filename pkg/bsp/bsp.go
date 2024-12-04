package bsp

import (
	"machine"
	"time"

	"github.com/soypat/cyw43439"
)

type Hardware struct {
	Wifi *cyw43439.Device
}

func Initialize(i2cMode machine.I2CMode) *Hardware {
	boardSpecificInit()

	return &Hardware{
		Wifi: initWIFI(),
	}
}

func (hw *Hardware) HardReset() {
	hw.LED(false)
	machine.CPUReset()
}

func (hw *Hardware) HelloLED() {
	for i := 0; i < 10; i++ {
		hw.LED(true)
		time.Sleep(100 * time.Millisecond)
		hw.LED(false)
		time.Sleep(100 * time.Millisecond)
	}
	hw.LED(true)
}
