package viral

import (
	"fmt"
)

/*
RETO VIRAL 12345678
Simula el reto viral 12345678.
Crea una función que cuente de
1 a 8, eliminando cada vez un
dígito y mostrando un espacio en
blanco en su lugar, de manera
ascendente y descendente.
*/

func Execute() {
	for i := 0; i <= 7; i++ {
		var viral string
		for j := 1; j <= 8; j++ {
			if j <= i {
				viral += " "
			} else {
				viral += fmt.Sprintf("%d", j)
			}
		}
		fmt.Println(viral)
	}

	for i := 6; i >= 0; i-- {
		var viral string
		for j := 1; j <= 8; j++ {
			if j <= i {
				viral = viral + " "
			} else {
				viral += fmt.Sprintf("%d", j)
			}
		}
		fmt.Println(viral)
	}
}
