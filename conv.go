package tempconv

// CtoF converts a Celsius temperature into its Fahrenheit equivalent
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CtoK converts a Celsius temperature into its Kelvin equivalent
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FtoC converts a Fahrenheit temperature into its Celsius equivalent
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FtoK converts a Fahrenheit temperature into its Kelvin equivalent
func FtoK(f Fahrenheit) Kelvin {
	return Kelvin(273.15 + (f - 32) * 5 / 9)
}

// KtoC converts a Kelvin temperature into its Celsius equivalent
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KtoF converts a Kelvin temperature into its Fahrenheit equivalent
func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit(k - 273.15) * 9/5 + 32
}
