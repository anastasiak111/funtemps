package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

import ("math")



// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	  result := (value - 32) * (5.0 / 9.0)
    return math.Round(result* 100) / 100
}


func CelsiusToFarhenheit(value float64) float64 {
    result := value*(9.0/5.0) + 32;
     return math.Round(result) 
}


func FarhenheitToKelvin(value float64) float64 {
    result := (value - 32 )*(5.0/9.0) + 273.15;
     return math.Round(result* 100) / 100
} 


func KelvinToFarhenheit(value float64) float64 {
      result :=  (value  - 273.15) * (9.0/5.0) +32
       return math.Round(result* 100) / 100
}

func KelvinToCelsius(value float64) float64 {
    result := value - 273.15
     return math.Round(result* 100) / 100
}

func CelsiusToKelvin(value float64) float64 {
    result := value + 273.15
     return math.Round(result* 100) / 100
}







