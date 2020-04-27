package main

import (
	"flag"
	"time"

	"github.com/dmundt/go-firmata"
)

func main() {
	port := flag.String("port", "COM1", "serial port")
	rel := flag.Int("pin", 1, "pin number [1..4]")
	dur := flag.Int("dur", 1, "duration in [s]")
	rate := flag.Int("rate", 57600, "baud rate")

	flag.Parse()

	if *rel < 1 || *rel > 4 {
		panic("Relay pin number is outside of valid range [1..4]")
	}

	// Open firmata connection.
	c, err := firmata.NewClient(*port, *rate)
	if err == nil {
		defer c.Close()

		// Toggle specific relay.
		pin := *rel + 1
		c.Log.Info("Close relay #%d", *rel)
		c.SetPinMode(byte(pin), firmata.Output)
		time.Sleep(time.Duration(*dur) * time.Second)
		c.Log.Info("Open relay #%d", *rel)
		c.SetPinMode(byte(pin), firmata.Input)
		c.Log.Close()
	}
}
