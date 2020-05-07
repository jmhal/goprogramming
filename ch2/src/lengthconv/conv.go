package lengthconv

// MToF converts a meters length to feet.
func MToF(m Meters) Feet { return Feet(m * 3.28084) }

// FToM converts a Fahrenheit temperature to Celsius.
func FToM(f Feet) Meters { return Meters(f / 3.28084) }

