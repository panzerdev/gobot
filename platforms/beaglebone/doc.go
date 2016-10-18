/*
Package beaglebone provides the Gobot adaptor for the Beaglebone Black.

Installing:

	go get github.com/hybridgroup/platforms/gobot/beaglebone

Example:

	package main

	import (
		"time"

		"github.com/hybridgroup/gobot"
		"github.com/hybridgroup/gobot/drivers/gpio"
		"github.com/hybridgroup/gobot/platforms/beaglebone"
	)

	func main() {
		beagleboneAdaptor := beaglebone.NewAdaptor()
		led := gpio.NewLedDriver(beagleboneAdaptor, "P9_12")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{beagleboneAdaptor},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For more information refer to the beaglebone README:
https://github.com/hybridgroup/gobot/blob/master/platforms/beaglebone/README.md
*/
package beaglebone
