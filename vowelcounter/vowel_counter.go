package vowelcounter

import "strings"

/*
CONTADOR DE VOCALES
Crea un programa que cuente cuantas
vocales tiene una cadena de texto.
*/

// First implementation

/* func Execute(input string) int {
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
} */

// Second implementation

func Execute(input string) int {
	var vocals = "aeiou"
	counter := 0

	for _, char := range input {
		charToLower := strings.ToLower(string(char))
		if strings.Contains(vocals, charToLower) {
			counter++
		}
	}

	return counter
}
