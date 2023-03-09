package conv
import (
	
  "math"
)

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahrenheit float64) float64 {
    celsius := (fahrenheit - 32) * 5.0/9.0
	return math.Round(celsius *100)/100
}


func KelvinToCelsius(kelvin float64) float64 {
  
	return math.Round((kelvin - 273.15)*100)/100
}

func CelsiusToFahrenheit(celsius float64) float64{

  fahrenheit := celsius * (9.0/5.0) + 32
  return math.Round(fahrenheit*100)/100
}

func CelsiusToKelvin(celsius float64) float64 {

  kelvin := celsius + 273.15
  return math.Round(kelvin*100)/100

}

func FarhenheitToKelvin(fahrenheit float64) float64 {

  kelvin := (fahrenheit - 32) * 5 / 9 + 273.15
  return math.Round(kelvin *100) / 100

}

func KelvinToFahrenheit(kelvin float64) float64 {

  fahrenheit := (kelvin - 273.15) * 9 / 5 + 32
  return math.Round(fahrenheit *100) / 100

}


