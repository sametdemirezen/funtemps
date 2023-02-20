package conv
import (
	
	
  "math"
)
/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahrenheit float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
    celsius := (fahrenheit - 32) * 5.0/9.0
	return math.Round(celsius *100)/100
}

// De andre konverteringsfunksjonene implementere her
// ...
func KelvinToCelsius(kelvin float64) float64 {
	if (kelvin >= -273.15){
    return 0;
  }
 
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


