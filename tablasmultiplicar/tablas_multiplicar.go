package tablasmultiplicar

import "fmt"

/*
TABLAS DE MULTIPLICAR
Imprime todas las tablas de
multiplicar de 1 al 10.
*/

func Execute() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n----------------------------------")
		fmt.Printf("\nTabla del %d", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("\nExecuting %d*%d = %d", i, j, i*j)
		}
	}
}
