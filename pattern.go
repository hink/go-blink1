package blink1

import "time"

// Pattern is a Blink(1) blink pattern
type Pattern struct {
	Repeat      uint          // How many times to repeat
	RepeatDelay time.Duration // Delay between repeats
	States      []State       // Slice of states to execute in pattern
}

// State is a Blink(1) light state
type State struct {
	Red      uint8         // Red value 0-255
	Green    uint8         // Green value 0-255
	Blue     uint8         // Blue value 0-255
	LED      uint8         // which LED to address (0=all, 1=1st LED, 2=2nd LED)
	FadeTime time.Duration // Fadetime to state
	Duration time.Duration // Duration of state after FadeTime
}

// OffState helper
var OffState = State{Duration: time.Duration(10) * time.Millisecond}

// LED helper constants, used to target specific LED's on the sides of Blink(1).
const (
	LEDAll uint8 = iota
	LED1
	LED2
)
