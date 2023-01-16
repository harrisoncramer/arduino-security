package main

// import (
// 	"machine"
// 	"time"
// )
//
// func main() {
// 	buzzer := machine.D9
// 	buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})
// 	buzzer.High()
// 	for {
// 		time.Sleep(time.Millisecond * 300)
// 		buzzer.High()
// 		time.Sleep(time.Millisecond * 300)
// 		buzzer.Low()
// 	}
// }

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
