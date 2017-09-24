package tempconv

//CToF konwertuje temperature w stopniach Celsjusza na stopnie Fahrenheita.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//CToK konwertuje temperature w stopniach Celsjusza na stopnie Kelvina.
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

//FToC konwertuje temperature w stopniach Fahrenheita na stopnie Celsjusza.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//FToK konwertuje temperature w stopniach Fahrenheita na stopnie Kelvina.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) - 273.15) }

//KToC konwertuje temperature w stopniach Kelvina na stopnie Celsjusza.
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

//KToF konwertuje temperature w stopniach Kelvina na stopnie Fahrenheita.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k*9/5 + 32) + 273.15) }
