// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius temperatures
type Celsius float64

// Fahrenheit temperatures
type Fahrenheit float64

// Kelvin temperatures
type Kelvin float64

// Key temperatures in Celsius
const (
	AbsoluteZeroC Celsius = -273.15 // The coldest possible temperature, in Celsius
	FreezingC     Celsius = 0       // Freezing temperature of water, in Celsius
	BoilingC      Celsius = 100     // Boiling temperature of water, in Celsius
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
