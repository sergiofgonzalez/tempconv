package tempconv

import "fmt"

// Celsius temperature type declaration
type Celsius float64

// Fahrenheit temperature type declaration
type Fahrenheit float64

const (
	// AbsoluteZeroC is the absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC is the absolute zero temperature in Celsius
	FreezingC Celsius = 0

	// BoilingC is the absolute zero temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
