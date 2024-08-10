package imc

import (
	"fmt"
)

/*
Initial implementation

func Execute(weight float64, height float64) (float64, error) {
	imc := weight / (height * height)
	stringRound := fmt.Sprintf("%.1f", imc)
	return strconv.ParseFloat(stringRound, 64)
}
*/

func Execute(weight float64, height float64) string {
	imc := weight / (height * height)
	return validateResult(imc)
}

func validateResult(imc float64) string {
	if imc < 18.5 {
		return fmt.Sprintf("IMC: %.1f, Bajo peso", imc)
	} else if imc >= 18.5 && imc < 25 {
		return fmt.Sprintf("IMC: %.1f, Normal", imc)
	} else if imc >= 25 && imc < 30 {
		return fmt.Sprintf("IMC: %.1f, Sobrepeso", imc)
	} else {
		return fmt.Sprintf("IMC: %.1f, Obesidad", imc)
	}
}
