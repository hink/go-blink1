package blink1

import (
	"errors"

	"time"

	"github.com/hink/go-blink1/libusb"
)

// USB IDs
const (
	USBVendorID  = 10168
	USBProductID = 493
)

var (
	errNoDevices = errors.New("No Blink(1) device found or all already in use")

	openDevices     = make(map[string]*Device)
	defaultFadeTime = 100
)

// Device Thingm Blink(1) USB device
type Device struct {
	Device          *libusb.Device // USB device
	DefaultFadeTime time.Duration  // Default time to fade between states
	CurrentState    State          // Current state of the Blink(1)
}

// OpenNextDevice opens and returns the next available Blink(1) device
func OpenNextDevice() (device *Device, err error) {
	// Refresh
	libusb.Init()

	// Enum devices and look for next Blink(1)
	for _, dev := range libusb.Enum() {
		if dev.Vid == USBVendorID && dev.Pid == USBProductID {
			if openDevices[dev.Device] == nil {
				d := libusb.Open(dev.Vid, dev.Pid, dev.Device)
				if d != nil {
					device = &Device{
						Device:          d,
						DefaultFadeTime: time.Duration(defaultFadeTime) * time.Millisecond,
					}
					openDevices[dev.Device] = device
					return
				}
			}
		}
	}
	err = errNoDevices
	return
}

// Close communication channel to Blink(1)
func (b *Device) Close() {
	delete(openDevices, b.Device.Device)
	_ = b.Device.Close()
}

// SetState sets the blink(1) to a specific state
func (b *Device) SetState(state State) (err error) {
	b.CurrentState = state
	bytesWritten := fadeToRgbBlink1(b, state.FadeTime, state.Red, state.Green, state.Blue, state.LED)
	if bytesWritten <= 0 {
		err = errors.New("Unable to write to blink(1)")
	}
	return
}

// RunPattern executes a predefined pattern
func (b *Device) RunPattern(pattern *Pattern) (err error) {
	if pattern.States == nil {
		pattern.States = []State{
			OffState,
		}
	} else {
		pattern.States = append(pattern.States, OffState)
	}

	if pattern.Repeat < 0 {
		pattern.Repeat = 0
	}

	for i := 0; i <= int(pattern.Repeat); i++ {
		for _, state := range pattern.States {
			err = b.SetState(state)
			if err != nil {
				// Try and fail silently so that successive patterns may succeed
				err = nil
				continue
			}
			if state.Duration != 0 {
				time.Sleep(state.Duration)
			}
		}
	}
	return
}
