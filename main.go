package tempconverter

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -237.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Returns the Celsius value in Float
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Returns the Fahrenheit value in Float
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// Returns the Kelvin value in Float
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

// Converts the celsius to fahrenheit
func (c Celsius) AsFahrenheit() Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// Converts the celsius to kelvin
func (c Celsius) AsKelvin() Kelvin { return Kelvin(c + 273.15) }

// Converts the fahrenheit to celsius
func (f Fahrenheit) AsCelsius() Celsius { return Celsius((f - 32) * 5 / 9) }

// Converts fahrenheit to kelvin
func (f Fahrenheit) AsKelvin() Kelvin { return Kelvin((f - 32) * 5 / 9 + 273.15) }

// Converts kelvin to celsius
func (k Kelvin) AsCelsius() Celsius { return Celsius(k - 273.15) }

// Converts kelvin to fahrenheit
func (k Kelvin) AsFahrenheit() Fahrenheit { return Fahrenheit((k - 273.15) * 9 / 5 + 32) }