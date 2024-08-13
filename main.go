package main

import (
	"fmt"

	"github.com/gugliio/retos-de-programacion-mini/anagrama"
	"github.com/gugliio/retos-de-programacion-mini/celsiusfahrenheit"
	"github.com/gugliio/retos-de-programacion-mini/decimaltobinary"
	"github.com/gugliio/retos-de-programacion-mini/evenodd"
	"github.com/gugliio/retos-de-programacion-mini/fizzbuzz"
	"github.com/gugliio/retos-de-programacion-mini/imc"
	"github.com/gugliio/retos-de-programacion-mini/revertstring"
	"github.com/gugliio/retos-de-programacion-mini/tablasmultiplicar"
	"github.com/gugliio/retos-de-programacion-mini/vowelcounter"
)

func main() {
	fmt.Println("Reto 1 - Fizzbuzz")
	fizzbuzz.Execute()

	fmt.Println("Reto 2 - Decimal to Binary")
	decimaltobinary.Execute(1)

	fmt.Println("Reto 3 - Even Odd")
	fmt.Println(evenodd.Execute(5))

	fmt.Println("Reto 4 - Vowel Counter")
	fmt.Println(vowelcounter.Execute("Hola Mundo!"))

	fmt.Println("Reto 5 - IMC")
	fmt.Println(imc.Execute(40, 1.75))

	fmt.Println("Reto 6 - Temp converter")
	fmt.Println(celsiusfahrenheit.Execute(40, "fahrenheit"))

	fmt.Println("Reto 7 - Anagrama")
	fmt.Println(anagrama.Execute("amor", "roma"))

	fmt.Println("Reto 8 - Tablas Multiplicar")
	tablasmultiplicar.Execute()

	fmt.Println("\nReto 9 - Revert string")
	fmt.Println(revertstring.Execute("Ortega"))
}
