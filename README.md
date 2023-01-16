# arduino-security

Code for an alarm using Arudino and TinyGo. The motion sensor will detect motion and the alarm will go off, flashing blue and red lights. The buzzer will also sound, alerting of an intruder.

![An image of the completed module](https://hjc-public.s3.amazonaws.com/arduino-security.jpeg)

## Requirements

1. <a href="https://tinygo.org/">Tinygo</a>
2. <a href="https://store.arduino.cc/products/arduino-uno-rev3">Arduino Uno</a>, power supply, and USB-B data-transfer cable
3. Jumper Cables
4. Leds (x2)
5. 330Ω Resistor (x2)
6. <a href="https://www.amazon.com/dp/B07KZW86YR">PIR Infrared Sensor</a>
7. <a href="https://www.amazon.com/dp/B07VK1GJ9X">Buzzers</a>

## Build and Flash Arduino

```bash
go mod tidy
tinygo flash --target=arduino --scheduler=tasks --port=/dev/cu.usbmodem101
```
