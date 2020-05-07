package tempconv

import (
   "fmt"
   "flag"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
   AbsoluteZeroC Celsius = -273.15
   FreezingC     Celsius = 0
   BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g K, k") }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

type celsiusFlag struct {
   Celsius
}

func (f *celsiusFlag) Set (s string) error {
   var unit string
   var value float64
   fmt.Sscanf(s, "%f%s", &value, &unit)
   switch unit {
      case "C", "ºC":
         f.Celsius = Celsius(value)
	 return nil
      case "F", "ºF":
         f.Celsius = FToC(Fahrenheit(value))
	 return nil
      case "K":
         f.Celsius = KToC(Kelvin(value))
	 return nil
   }
   return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
   f := celsiusFlag{value}
   flag.CommandLine.Var(&f, name, usage)
   return &f.Celsius
}
