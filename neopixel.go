package ws2812fxfirmata

import (
	"strconv"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/firmata"
)

const (
	// serial commands
	NEOPIXEL_CMD = 0x51
	// turn strip off
	NEOPIXEL_OFF = 0x00
	// configure the strip
	NEOPIXEL_CONFIG = 0x01
	// show currently set pixels
	NEOPIXEL_SHOW = 0x02
	// set the color value of pixel n using 32bit packed color value
	NEOPIXEL_SET_PIXEL = 0x03
	NEOPIXEL_SET_STRIP = 0x04
	// TODO: shift all the pixels n places along the strip
	NEOPIXEL_SHIFT = 0x05

	PIXEL_RANDOM_MODE = 0x06
	PIXEL_SET_MODE    = 0x07
	// TODO: resort these
	NEOPIXEL_ON             = 0x08
	NEOPIXEL_SET_COLOR      = 0x09
	NEOPIXEL_SET_BRIGHTNESS = 0x10

	// WS2812FX neopixel modes
	// a list of modes can be found in the WS2812FX repo:
	// https://github.com/kitesurfer1404/WS2812FX/blob/837b7dd843f0e0c38347c13e4b4594d6dc603162/src/WS2812FX.h#L120-L180
	FX_MODE_STATIC                 = 0
	FX_MODE_BLINK                  = 1
	FX_MODE_BREATH                 = 2
	FX_MODE_COLOR_WIPE             = 3
	FX_MODE_COLOR_WIPE_INV         = 4
	FX_MODE_COLOR_WIPE_REV         = 5
	FX_MODE_COLOR_WIPE_REV_INV     = 6
	FX_MODE_COLOR_WIPE_RANDOM      = 7
	FX_MODE_RANDOM_COLOR           = 8
	FX_MODE_SINGLE_DYNAMIC         = 9
	FX_MODE_MULTI_DYNAMIC          = 10
	FX_MODE_RAINBOW                = 11
	FX_MODE_RAINBOW_CYCLE          = 12
	FX_MODE_SCAN                   = 13
	FX_MODE_DUAL_SCAN              = 14
	FX_MODE_FADE                   = 15
	FX_MODE_THEATER_CHASE          = 16
	FX_MODE_THEATER_CHASE_RAINBOW  = 17
	FX_MODE_RUNNING_LIGHTS         = 18
	FX_MODE_TWINKLE                = 19
	FX_MODE_TWINKLE_RANDOM         = 20
	FX_MODE_TWINKLE_FADE           = 21
	FX_MODE_TWINKLE_FADE_RANDOM    = 22
	FX_MODE_SPARKLE                = 23
	FX_MODE_FLASH_SPARKLE          = 24
	FX_MODE_HYPER_SPARKLE          = 25
	FX_MODE_STROBE                 = 26
	FX_MODE_STROBE_RAINBOW         = 27
	FX_MODE_MULTI_STROBE           = 28
	FX_MODE_BLINK_RAINBOW          = 29
	FX_MODE_CHASE_WHITE            = 30
	FX_MODE_CHASE_COLOR            = 31
	FX_MODE_CHASE_RANDOM           = 32
	FX_MODE_CHASE_RAINBOW          = 33
	FX_MODE_CHASE_FLASH            = 34
	FX_MODE_CHASE_FLASH_RANDOM     = 35
	FX_MODE_CHASE_RAINBOW_WHITE    = 36
	FX_MODE_CHASE_BLACKOUT         = 37
	FX_MODE_CHASE_BLACKOUT_RAINBOW = 38
	FX_MODE_COLOR_SWEEP_RANDOM     = 39
	FX_MODE_RUNNING_COLOR          = 40
	FX_MODE_RUNNING_RED_BLUE       = 41
	FX_MODE_RUNNING_RANDOM         = 42
	FX_MODE_LARSON_SCANNER         = 43
	FX_MODE_COMET                  = 44
	FX_MODE_FIREWORKS              = 45
	FX_MODE_FIREWORKS_RANDOM       = 46
	FX_MODE_MERRY_CHRISTMAS        = 47
	FX_MODE_FIRE_FLICKER           = 48
	FX_MODE_FIRE_FLICKER_SOFT      = 49
	FX_MODE_FIRE_FLICKER_INTENSE   = 50
	FX_MODE_CIRCUS_COMBUSTUS       = 51
	FX_MODE_HALLOWEEN              = 52
	FX_MODE_BICOLOR_CHASE          = 53
	FX_MODE_TRICOLOR_CHASE         = 54
	FX_MODE_ICU                    = 55
	FX_MODE_CUSTOM                 = 56
	FX_MODE_CUSTOM_0               = 56
	FX_MODE_CUSTOM_1               = 57
	FX_MODE_CUSTOM_2               = 58
	FX_MODE_CUSTOM_3               = 59
)

// NeopixelDriver represents a connection to a NeoPixel
type NeopixelDriver struct {
	name       string
	pin        string
	pixelCount uint16
	connection *firmata.Adaptor
	gobot.Eventer
}

// NewNeopixelDriver returns a new NeopixelDriver
func NewNeopixelDriver(a *firmata.Adaptor, pin string, pixelCount uint16) *NeopixelDriver {
	neo := &NeopixelDriver{
		name:       gobot.DefaultName("Neopixel"),
		connection: a,
		pin:        pin,
		pixelCount: pixelCount,
		Eventer:    gobot.NewEventer(),
	}

	return neo
}

// Start starts up the NeopixelDriver
func (neo *NeopixelDriver) Start() (err error) {
	i, _ := strconv.Atoi(neo.pin)
	return neo.SetConfig(uint8(i), neo.pixelCount)
}

// Halt stops the NeopixelDriver
func (neo *NeopixelDriver) Halt() (err error) {
	return neo.Off()
}

// Name returns the Driver's name
func (neo *NeopixelDriver) Name() string { return neo.name }

// SetName sets the Driver's name
func (neo *NeopixelDriver) SetName(n string) { neo.name = n }

// Pin returns the Driver's pin
func (neo *NeopixelDriver) Pin() string { return neo.pin }

// Connection returns the Driver's Connection
func (neo *NeopixelDriver) Connection() gobot.Connection { return neo.connection }

// Off turns off all the Neopixels in the strip
func (neo *NeopixelDriver) Off() error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_OFF})
}

// On turns on all the Neopixels in the strip
func (neo *NeopixelDriver) On() error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_ON})
}

// Show activates all the Neopixels in the strip
func (neo *NeopixelDriver) Show() error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_SHOW})
}

func (neo *NeopixelDriver) RandomMode() error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, PIXEL_RANDOM_MODE})
}

func (neo *NeopixelDriver) SetMode(mode uint8) error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, PIXEL_SET_MODE, byte(mode)})
}

func (neo *NeopixelDriver) SetBrightness(brightness uint8) error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_SET_BRIGHTNESS, byte(brightness)})
}

// SetPixel sets the color of one specific Neopixel in the strip
func (neo *NeopixelDriver) SetPixel(pix uint16, color uint32) error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_SET_PIXEL,
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
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_SET_COLOR,
		byte(color & 0x7F),
		byte((color >> 7) & 0x7F),
		byte((color >> 14) & 0x7F),
		byte((color >> 21) & 0x7F),
	})
}

// SetConfig sets the config info for the Neopixel strip
func (neo *NeopixelDriver) SetConfig(pin uint8, len uint16) error {
	return neo.connection.WriteSysex([]byte{NEOPIXEL_CMD, NEOPIXEL_CONFIG,
		byte(pin),
		byte(len & 0x7F),
		byte((len >> 7) & 0x7F),
	})
}
