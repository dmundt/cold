package main

import (
	"github.com/dmundt/go-firmata"
	"time"
)

func main() {
	c, err := firmata.NewClient("COM4", 57600)
	if err == nil {
		defer c.Close()

		c.DigitalWrite(2, true)
		c.DigitalWrite(3, true)
		c.DigitalWrite(4, true)
		c.DigitalWrite(5, true)

		time.Sleep(2 * time.Second)
	}
}
