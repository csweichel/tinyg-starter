//go:build picow
// +build picow

package bsp

import (
	"github.com/soypat/cyw43439"
)

func boardSpecificInit() {
}

func initWIFI() *cyw43439.Device {
	dev := cyw43439.NewPicoWDevice()
	cfg := cyw43439.DefaultWifiConfig()
	// cfg.Logger = logger // Uncomment to see in depth info on wifi device functioning.
	err := dev.Init(cfg)
	if err != nil {
		panic(err)
	}
	return dev
}

func (hw *Hardware) LED(on bool) {
	hw.Wifi.GPIOSet(0, on)
}
