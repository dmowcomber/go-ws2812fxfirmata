package main

import (
	"fmt"
	"time"

	"github.com/dmowcomber/go-ws2812fxfirmata"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	arduinoAddress := "192.168.1.113:3030"
	firmataAdaptor := firmata.NewTCPAdaptor(arduinoAddress)

	ws2812fx := ws2812fxfirmata.NewNeopixelDriver(firmataAdaptor.Adaptor)

	fmt.Println("connecting")
	err := firmataAdaptor.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ws2812fxPin := uint8(5)
	ws2812fxCount := uint16(7)
	ws2812fx.SetConfig(ws2812fxPin, ws2812fxCount)

	ws2812fx.SetBrightness(uint8(1))

	for {
		time.Sleep(5 * time.Second)
		ws2812fx.ModeCycle()
	}

}
