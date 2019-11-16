package ws2812fxfirmata

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/firmata"
)

const (
	neopixelCmd           = 0x51
	neopixelSetBrightness = 0x00
	neopixelSetColor      = 0x01
	neopixelSetPixel      = 0x02
	neopixelSetStrip      = 0x03
	neopixelShift         = 0x04
	neopixelSetMode       = 0x05
	neopixelModeCycle     = 0x06
	neopixelStart         = 0x07
	neopixelStop          = 0x08
	neopixelConfig        = 0x09

	// WS2812FX neopixel modes
	// a list of modes can be found in the WS2812FX repo:
	// https://github.com/kitesurfer1404/WS2812FX/blob/837b7dd843f0e0c38347c13e4b4594d6dc603162/src/WS2812FX.h#L120-L180
	FXModeStatic               = 0
	FXModeBlink                = 1
	FXModeBreath               = 2
	FXModeColorWipe            = 3
	FXModeColorWipeInv         = 4
	FXModeColorWipeRev         = 5
	FXModeColorWipeRevInv      = 6
	FXModeColorWipeRandom      = 7
	FXModeRandomColor          = 8
	FXModeSingleDynamic        = 9
	FXModeMultiDynamic         = 10
	FXModeRainbow              = 11
	FXModeRainbowCycle         = 12
	FXModeScan                 = 13
	FXModeDualScan             = 14
	FXModeFade                 = 15
	FXModeTheaterChase         = 16
	FXModeTheaterChaseRainbow  = 17
	FXModeRunningLights        = 18
	FXModeTwinkle              = 19
	FXModeTwinkleRandom        = 20
	FXModeTwinkleFade          = 21
	FXModeTwinkleFadeRandom    = 22
	FXModeSparkle              = 23
	FXModeFlashSparkle         = 24
	FXModeHyperSparkle         = 25
	FXModeStrobe               = 26
	FXModeStrobeRainbow        = 27
	FXModeMultiStrobe          = 28
	FXModeBlinkRainbow         = 29
	FXModeChaseWhite           = 30
	FXModeChaseColor           = 31
	FXModeChaseRandom          = 32
	FXModeChaseRainbow         = 33
	FXModeChaseFlash           = 34
	FXModeChaseFlashRandom     = 35
	FXModeChaseRainbowWhite    = 36
	FXModeChaseBlackout        = 37
	FXModeChaseBlackoutRainbow = 38
	FXModeColorSweepRandom     = 39
	FXModeRunningColor         = 40
	FXModeRunningRedBlue       = 41
	FXModeRunningRandom        = 42
	FXModeLarsonScanner        = 43
	FXModeComet                = 44
	FXModeFireworks            = 45
	FXModeFireworksRandom      = 46
	FXModeMerryChristmas       = 47
	FXModeFireFlicker          = 48
	FXModeFireFlickerSoft      = 49
	FXModeFireFlickerIntense   = 50
	FXModeCircusCombustus      = 51
	FXModeHalloween            = 52
	FXModeBicolorChase         = 53
	FXModeTricolorChase        = 54
	FXModeIcu                  = 55
	FXModeCustom               = 56
	FXModeCustom0              = 56
	FXModeCustom1              = 57
	FXModeCustom2              = 58
	FXModeCustom3              = 59
)

// NeopixelDriver represents a connection to a NeoPixel
type NeopixelDriver struct {
	name       string
	connection *firmata.Adaptor
	gobot.Eventer
}

// NewNeopixelDriver returns a new NeopixelDriver
func NewNeopixelDriver(a *firmata.Adaptor) *NeopixelDriver {
	neo := &NeopixelDriver{
		name:       gobot.DefaultName("Neopixel"),
		connection: a,
		Eventer:    gobot.NewEventer(),
	}

	return neo
}

// SetConfig sets the pin and pixel count for the Neopixel strip
func (neo *NeopixelDriver) SetConfig(pin uint8, pixelCount uint16) error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelConfig,
		byte(pin),
		byte(pixelCount & 0x7F),
		byte((pixelCount >> 7) & 0x7F),
	})
}

// Stop turns off all the Neopixels in the strip
func (neo *NeopixelDriver) Stop() error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelStop})
}

// Start turns on all the Neopixels in the strip
func (neo *NeopixelDriver) Start() error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelStart})
}

// ModeCycle cycles through modes
func (neo *NeopixelDriver) ModeCycle() error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelModeCycle})
}

// SetMode sets the mode
func (neo *NeopixelDriver) SetMode(mode uint8) error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelSetMode, byte(mode)})
}

// SetBrightness sets the brightness
func (neo *NeopixelDriver) SetBrightness(brightness uint8) error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelSetBrightness, byte(brightness)})
}

// SetPixel sets the color of one specific Neopixel in the strip
func (neo *NeopixelDriver) SetPixel(pix uint16, color uint32) error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelSetPixel,
		byte(pix & 0x7F),
		byte((pix >> 7) & 0x7F),
		byte(color & 0x7F),
		byte((color >> 7) & 0x7F),
		byte((color >> 14) & 0x7F),
		byte((color >> 21) & 0x7F),
	})
}

// SetColor sets the color of the strip
func (neo *NeopixelDriver) SetColor(color uint32) error {
	return neo.connection.WriteSysex([]byte{neopixelCmd, neopixelSetColor,
		byte(color & 0x7F),
		byte((color >> 7) & 0x7F),
		byte((color >> 14) & 0x7F),
		byte((color >> 21) & 0x7F),
	})
}
