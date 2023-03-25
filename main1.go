/*package main

import (
	"flag"
	"fmt"
	"github.com/sametdemirezen/funtemps/conv"
	"github.com/sametdemirezen/funtemps/funfacts"
	
)

var fahr float64
var celsius float64
var kelvin float64
var out string
var funFacts string
var t string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "temp", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funFacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "f", "temperaturen relevant med flagget")

}

func main() {

	flag.Parse()
	//Konverterings metoder	
	if isFlagPassed("F") {
		if out == "C" {
			celsius := conv.FarhenheitToCelsius(fahr)
			fmt.Printf("%.2f°F er %.2f°C", fahr, celsius)
		}
		if out == "K" {
		fahr := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.2f °F er %.2f °K", fahr, kelvin)
		}
	} else if isFlagPassed("C"){
		if out == "F" {
			fahr := conv.CelsiusToFahrenheit(celsius)
			fmt.Printf("%.2f°C er %.2f°F", celsius, fahr)
		}
		if out == "K" {
			kelvin :=conv.CelsiusToKelvin(celsius)
			fmt.Printf("%.2f°C er %.2f°K", celsius, kelvin)
		}
	} else if isFlagPassed("K"){
		if out == "C" {
			cels := conv.KelvinToCelsius(kelvin)
			fmt.Printf("%.2f°K er %.2f°C", kelvin, cels)
		}
		if out == "K" {
			fahr := conv.KelvinToFahrenheit(kelvin)
			fmt.Printf("%.2f°K er %.2f°F", kelvin, fahr)
		}
	}

	//Funfacts metoder
	if funFacts == "terra" && isFlagPassed("funfacts") {
		
		terraFact := funfacts.GetFunFacts(funFacts)

		if t == "C" {
			fmt.Printf("%s %s C \n%s %s C \n%s %s °C", terraFact[0], "56.7", terraFact[1], "-89.4", terraFact[2], "9118")				
		}else if t == "K" {
			fmt.Printf("%s %f K\n %s %f \n%s %f K", terraFact[0], conv.CelsiusToKelvin(56.7), terraFact[1], conv.CelsiusToKelvin(-89.4), terraFact[2], conv.CelsiusToKelvin(9118))
		} else if t == "F" {
			fmt.Printf("%s %s °F \n%s %f \n%s %f °F", terraFact[0], "134", terraFact[1], conv.CelsiusToFahrenheit(-89.4), terraFact[2], conv.CelsiusToFahrenheit(9392))
		}

	}

	if funFacts == "sun" && isFlagPassed("funfacts") {
		
		sunFact := funfacts.GetFunFacts(funFacts)

		if t == "C" {
			fmt.Printf("%s %s C.\n %s %f C", sunFact[0], "15000000", sunFact[1], conv.KelvinToCelsius(5778))				
		}else if t == "K" {
			fmt.Printf("%s %f K.\n %s %s K", sunFact[0], conv.CelsiusToKelvin(15000000), sunFact[1],"5778")
		} else if t == "F" {
			fmt.Printf("%s %f °F. \n%s %f", sunFact[0], conv.CelsiusToFahrenheit(15000000), sunFact[1], conv.KelvinToFahrenheit(5778))
		}

	}

	if funFacts == "luna" && isFlagPassed("funfacts") {
		
		lunaFact := funfacts.GetFunFacts(funFacts)

		if t == "C" {
			fmt.Printf("%s %s C.\n %s %s C", lunaFact[0], "-183", lunaFact[1], "106")				
		}else if t == "K" {
			fmt.Printf("%s %f K.\n %s %f K", lunaFact[0], conv.CelsiusToKelvin(-183), lunaFact[1],conv.CelsiusToKelvin(106))
		} else if t == "F" {
			fmt.Printf("%s %f °F. \n%s %f", lunaFact[0], conv.CelsiusToFahrenheit(-183), lunaFact[1], conv.CelsiusToFahrenheit(106))
		}

	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
