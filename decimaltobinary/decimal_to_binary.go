package decimaltobinary

import (
	"fmt"
)

/*
DECIMAL A BINARIO
Crea un programa se encargue de
transformar un número decimal a binario.
*/

func Execute(decimal int) string {
	var binary string

	for decimal > 0 {
		binary = fmt.Sprintf("%d%s", decimal%2, binary)
		decimal /= 2
	}

	if binary == "" {
		return "0"
	}

	return binary
}
