package main

import (
	"github.com/dmundt/go-firmata"
	"time"
)

func main() {
	c, err := firmata.NewClient("COM4", 57600)
	if err == nil {
		defer c.Close()

		c.SetPinMode(2, firmata.Output)
		c.SetPinMode(3, firmata.Output)
		c.DigitalWrite(2, true)
		c.DigitalWrite(3, true)

		time.Sleep(2 * time.Second)
	}
}
