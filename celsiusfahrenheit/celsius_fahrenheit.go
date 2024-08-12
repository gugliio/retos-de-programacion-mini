package celsiusfahrenheit

/*
CONVERSOR DE TEMPERATURAS
Crea un conversor entre
grados Celsius y Fahrenheit.
*/

func Execute(temp float64, tempType string) float64 {
	if tempType == "fahrenheit" {
		return (temp - 32) * 5 / 9
	}
	if tempType == "celsius" {
		return temp*9/5 + 32
	}
	return 0
}
