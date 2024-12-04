//go:build !picow
// +build !picow

package bsp

import (
	"machine"

	"github.com/soypat/cyw43439"
)

func initWIFI() *cyw43439.Device {
	return nil
}

func boardSpecificInit() {
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (hw *Hardware) LED(on bool) {
	machine.LED.Set(on)
}
