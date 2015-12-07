package blink1

/*
	#cgo LDFLAGS: -lusb
	#include <usb.h>
*/
import "C"

import (
	"time"

	"github.com/hink/go-blink1/libusb"
)

func fadeToRgbBlink1(device *Device, fadeTime time.Duration, red, green, blue, normal int) (bytesWritten int) {
	bytesWritten = libusb.SendBlink1Command(device.Device, toMs(fadeTime), red, blue, green, normal)
	return
}
