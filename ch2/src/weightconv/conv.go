package weightconv

// KToP converts a kilograms weight to pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462) }

// FToC converts a Fahrenheit temperature to Celsius.
func PToK(p Pounds) Kilograms { return Kilograms(p / 2.20462) }

