package revertstring

/*
INVERSIÓN DE CADENAS
Crea una función que invierta
una cadena de texto.
*/

func Execute(word string) string {
	var revertString string

	for _, char := range word {
		revertString = string(char) + revertString
	}

	return revertString
}
