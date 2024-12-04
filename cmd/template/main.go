package main

import (
	"machine"

	"github.com/csweichel/splitflap/pkg/bsp"
	"github.com/csweichel/splitflap/pkg/cmdline"
)

func main() {
	hw := bsp.Initialize(machine.I2CModeController)
	hw.HelloLED()

	go cmdline.Repl(map[string]cmdline.CommandFunc{
		"reset": func(args []string) error {
			hw.HardReset()
			return nil
		},
	})

	<-make(chan struct{})
}
