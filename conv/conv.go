package conv

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahrenheit float64) float64 {
     
	return (fahrenheit - 32) * 5.0/9.0
}


func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func CelsiusToFahrenheit(celsius float64) float64{
  return celsius * (9.0/5.0) + 32
}

func CelsiusToKelvin(celsius float64) float64 {
  return celsius + 273.15
}

func FarhenheitToKelvin(fahrenheit float64) float64 {
  return (fahrenheit - 32) * 5 / 9 + 273.15
}

func KelvinToFahrenheit(kelvin float64) float64 {
  return (kelvin - 273.15) * 9 / 5 + 32
}


