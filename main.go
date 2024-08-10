package main

import (
	"fmt"

	"github.com/gugliio/retos-de-programacion-mini/decimaltobinary"
	"github.com/gugliio/retos-de-programacion-mini/evenodd"
	"github.com/gugliio/retos-de-programacion-mini/fizzbuzz"
	"github.com/gugliio/retos-de-programacion-mini/vowelcounter"
)

func main() {
	fmt.Println("Reto 0 - Fizzbuzz")
	fizzbuzz.Execute()

	fmt.Println("Reto 1 - Decimal to Binary")
	decimaltobinary.Execute(1)

	fmt.Println("Reto 2 - Even Odd")
	fmt.Println(evenodd.Execute(5))

	fmt.Println("Reto 3 - Vowel Counter")
	fmt.Println(vowelcounter.Execute("Hola Mundo!"))
}
