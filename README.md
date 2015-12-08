# Install libusb and libusb-compat (OSX Users)

Before using this library on OSX, you'll need to install the libusb and libusb-compat libraries.

This is easily done using the [Homebrew](http://brew.sh/) OSX package manager.
(Other OSes will need similar installation using their respective package managers)

```bash
brew install libusb libusb-compat
```

# blink1
--
    import "github.com/hink/go-blink1"


## Usage

```go
const (
	USBVendorID  = 10168
	USBProductID = 493
)
```
USB IDs

```go
var OffState = State{Duration: time.Duration(10) * time.Millisecond}
```
OffState helper

#### type Device

```go
type Device struct {
	Device          *libusb.Device // USB device
	DefaultFadeTime time.Duration  // Default time to fade between states
	CurrentState    State          // Current state of the Blink(1)
}
```

Device Thingm Blink(1) USB device

#### func  OpenNextDevice

```go
func OpenNextDevice() (device *Device, err error)
```
OpenNextDevice opens and returns the next available Blink(1) device

#### func (*Device) Close

```go
func (b *Device) Close()
```
Close communication channel to Blink(1)

#### func (*Device) RunPattern

```go
func (b *Device) RunPattern(pattern *Pattern) (err error)
```
RunPattern executes a predefined pattern

#### func (*Device) SetState

```go
func (b *Device) SetState(state State) (err error)
```
SetState sets the blink(1) to a specific state

#### type Pattern

```go
type Pattern struct {
	Repeat      uint          // How many times to repeat
	RepeatDelay time.Duration // Delay between repeats
	States      []State       // Slice of states to execute in pattern
}
```

Pattern is a Blink(1) blink pattern

#### type State

```go
type State struct {
	Red      uint8         // Red value 0-255
	Green    uint8         // Green value 0-255
	Blue     uint8         // Blue value 0-255
	Normal   uint8         // Normal value 0-255
	FadeTime time.Duration // Fadetime to state
	Duration time.Duration // Duration of state after FadeTime
}
```

State is a Blink(1) light state
