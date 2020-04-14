package main

import (
	"flag"
	"time"

	"github.com/dmundt/go-firmata"
)

func main() {
	port := flag.String("port", "COM1", "serial port")
	rel := flag.Int("rel", 1, "relay number")
	dur := flag.Int("dur", 1, "duration in [s]")

	flag.Parse()

	if *rel < 1 || *rel > 4 {
		panic("Relay pin number is outside of valid range [1..4]")
	}

	c, err := firmata.NewClient(*port, 57600)
	if err == nil {
		defer c.Close()

		c.Log.Info("Close relay #%d", *rel)
		c.SetPinMode(byte(*rel)+1, firmata.Output)
		time.Sleep(time.Duration(*dur) * time.Second)
		c.Log.Info("Open relay #%d", *rel)
		c.SetPinMode(byte(*rel)+1, firmata.Input)
		c.Log.Close()
	}
}
