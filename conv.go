package tempconv

// CtoF converts a Celsius temperature into its Fahrenheit equivalent
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converts a Fahrenheit temperature into its Celsius equivalent
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
