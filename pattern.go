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
	Normal   uint8         // Normal value 0-255
	FadeTime time.Duration // Fadetime to state
	Duration time.Duration // Duration of state after FadeTime
}

// OffState helper
var OffState = State{Duration: time.Duration(10) * time.Millisecond}
