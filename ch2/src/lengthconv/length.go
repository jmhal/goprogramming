// Package tempconv performs Meters and Feet conversions.
package lengthconv

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string { return fmt.Sprintf("%g ft", f) }
