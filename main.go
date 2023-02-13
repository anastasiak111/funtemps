package main

import (
	"flag"
	"fmt"
	"github.com/ANASTASIK111/funtemps/conv"
)

var fahr float64
var celsius float64 
var kelvin float64 
var out string
var funfacts string
var t string 

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "0.0", "temperaturskala, C, K eller F for funfacts")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	/*
	fmt.Println(fahr, out, funfacts)
	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())
	fmt.Println(isFlagPassed("out"))
	*/

	
	if isFlagPassed("out"){
		if out == "C" && isFlagPassed("F") {
		fmt.Println(conv.FarhenheitToCelsius(fahr))	
	    } else if out == "F" && isFlagPassed("C") {
		fmt.Println(conv.CelsiusToFarhenheit(celsius))
	   }
	}

	if isFlagPassed("out"){
		if out == "K" && isFlagPassed("C") {
		fmt.Println(conv.KelvinToCelsius(celsius))	
	    } else if out == "C" && isFlagPassed("K") {
		fmt.Println(conv.CelsiusToKelvin(kelvin))
	   }
	}

	if isFlagPassed("out"){
		if out == "K" && isFlagPassed("F") {
		fmt.Println(conv.FarhenheitToKelvin(kelvin))	
	    } else if out == "F" && isFlagPassed("K") {
		fmt.Println(conv.KelvinToFarhenheit(fahr))
	   }
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
