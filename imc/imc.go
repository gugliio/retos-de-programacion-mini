package imc

import (
	"fmt"
	"strconv"
)

func Execute(weight float64, height float64) (float64, error) {
	imc := weight / (height * height)
	stringRound := fmt.Sprintf("%.1f", imc)
	return strconv.ParseFloat(stringRound, 64)
}
