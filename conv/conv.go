package conv



// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	  result := (value - 32) * 5 / 9
    return result
}


func CelsiusToFarhenheit(value float64) float64 {
    result := value*(9.0/5.0) + 32;
     return result 
}


func FarhenheitToKelvin(value float64) float64 {
    result := (value - 32 )*(5.0/9.0) + 273.15;
     return result
} 


func KelvinToFarhenheit(value float64) float64 {
      result :=  (value  - 273.15) * (9.0/5.0) +32
       return result
}

func KelvinToCelsius(value float64) float64 {
    result := value - 273.15
     return result
}

func CelsiusToKelvin(value float64) float64 {
    result := value + 273.15
     return result
}











