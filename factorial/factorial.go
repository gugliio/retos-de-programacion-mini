package factorial

/*
FACTORIAL
Crea una función que calcule
el factorial de un número.
*/

func Execute(number int) int {
	factorial := 1
	for i := 1; i < number; i++ {
		factorial = factorial + (i * factorial)
	}
	return factorial
}
