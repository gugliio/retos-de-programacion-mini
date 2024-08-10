package evenodd

/*
PAR O IMPAR
Crea un programa que compruebe si
un número entero es par o impar.
*/

func Execute(number int) string {
	if number%2 == 0 {
		return "par"
	}
	return "impar"
}
