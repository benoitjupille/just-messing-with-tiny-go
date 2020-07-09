package main

import (
	"machine"
	"time"
)

func main() {
	// chosing pin 13, the built in led on arduino cards
	var led machine.Pin = machine.Pin(13)

	// set output mode
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High() // led up
		delay(500) // wait
		led.Low() // led off
		delay(500) // wait
	}
}

// Put the execution to sleep during the given time in milliseconds
func delay(milliseconds int64) {
	time.Sleep(time.Duration(1000000 * milliseconds))
}