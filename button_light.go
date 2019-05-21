package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11")
	led := gpio.NewLedDriver(r, "7")
	lcd := gpio.NewJHD1313M1Driver(r)
	
	work := func() {
		fmt.Println("starting work!")

		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
			fmt.Printf("led %+v", led)
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
			fmt.Printf("led %+v", led)
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
}
