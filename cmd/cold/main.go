package main

import (
	"flag"
	"time"

	"github.com/dmundt/go-firmata"
)

func tooglePin(c firmata.FirmataClient, rel int, dur int) {
	// Toggle specific relay.
	// /pin := rel + 1
	c.Log.Info("Close relay #%d", rel)
	// c.SetPinMode(byte(pin), firmata.Output)
	time.Sleep(time.Duration(dur) * time.Second)
	c.Log.Info("Open relay #%d", rel)
	// c.SetPinMode(byte(pin), firmata.Input)
}

func main() {
	port := flag.String("port", "COM1", "serial port")
	rel := flag.Int("pin", 1, "pin number [1..4]")
	dur := flag.Int("dur", 1, "duration in [s]")
	rate := flag.Int("rate", 57600, "baud rate")
	mode := flag.Bool("all", false, "switch all")

	flag.Parse()

	if *rel < 1 || *rel > 4 {
		panic("Relay pin number is outside of valid range [1..4]")
	}

	// Open firmata connection.
	c, err := firmata.NewClient(*port, *rate)
	if err == nil {
		defer c.Close()
		if !*mode {
			// Toggle specific relay.
			tooglePin(*c, *rel, *dur)
		} else {
			// Toggle all relays.
			for i := 1; i < 5; i++ {
				tooglePin(*c, i, *dur)
			}
		}
		c.Log.Close()
	}
}
