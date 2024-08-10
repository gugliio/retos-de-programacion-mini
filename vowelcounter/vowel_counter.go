package vowelcounter

import "strings"

/*
CONTADOR DE VOCALES
Crea un programa que cuente cuantas
vocales tiene una cadena de texto.
*/

func Execute(input string) int {
	var vocals = []string{"a", "e", "i", "o", "u"}
	counter := 0

	for _, char := range input {
		for _, c := range vocals {
			if strings.ToLower(string(char)) == c {
				counter++
				break
			}
		}
	}

	return counter
}
