package tempconv

//CToF konwertuje temperature w stopniach Celsjusza na stopnie Fahrenheita.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC konwertuje temperature w stopniach Fahrenheita na stopnie Celsjusza.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
