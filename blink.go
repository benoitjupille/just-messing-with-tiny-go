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
    time.Sleep(time.Millisecond * 500)
    led.Low() // led off
    time.Sleep(time.Millisecond * 500)
  }
}
