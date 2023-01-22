package main

import (
	"machine"
	"time"
)

/*
		When signaled to start, the alarm struct will turn
	  on and off the LEDs and buzzer repeatedly. When signaled
	  to stop, the Alarm will turn them both off and remain off
	  until the next signal to begin an alarm cycle.
*/
type Alarm struct {
	buzzer  machine.Pin
	led1    machine.Pin
	led2    machine.Pin
	on      chan bool
	enabled bool
}

func (a *Alarm) Run(alarmCount int) {
	if !a.enabled {
		return
	}

	// a.buzzer.High()
	if alarmCount%2 == 0 {
		a.led1.High()
	} else {
		a.led2.High()
	}
	time.Sleep(100 * time.Millisecond)
	a.buzzer.Low()
	a.led1.Low()
	a.led2.Low()
	time.Sleep(100 * time.Millisecond)
}

func (a *Alarm) Off() {
	a.buzzer.Low()
	a.led1.Low()
	a.led2.Low()
}

func main() {

	machine.Serial.Configure(machine.UARTConfig{})
	machine.Serial.Write([]byte("Starting application..."))

	red := machine.D13
	blue := machine.D11

	motionSensor := machine.D2
	buzzer := machine.D9

	red.Configure(machine.PinConfig{Mode: machine.PinOutput})
	blue.Configure(machine.PinConfig{Mode: machine.PinOutput})
	motionSensor.Configure(machine.PinConfig{Mode: machine.PinInput})
	buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})

	alarm := Alarm{
		buzzer:  buzzer,
		led1:    red,
		led2:    blue,
		on:      make(chan bool, 2),
		enabled: true,
	}

	go func() {
		for on := range alarm.on {
			if on {
				machine.Serial.Write([]byte("Alarm!\n"))
				for i := 0; i <= 3; i++ {
					alarm.Run(i)
				}
			} else {
				alarm.Off()
			}
		}
	}()

	for {
		isMotion := motionSensor.Get()
		if isMotion {
			alarm.on <- true
		} else {
			alarm.on <- false
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
