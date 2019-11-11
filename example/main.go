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

	neoPixelPin := "5"
	neoPixelCount := uint16(7)
	neoPixel := ws2812fxfirmata.NewNeopixelDriver(firmataAdaptor.Adaptor, neoPixelPin, neoPixelCount)

	fmt.Println("connecting")
	err := firmataAdaptor.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO: this is redundant but without it pixels don't show correctly. need to fix this
	neoPixel.SetConfig(5, 7)

	neoPixel.SetBrightness(uint8(255))
	for {
		time.Sleep(5 * time.Second)
		neoPixel.RandomMode()
	}

}
