package imc_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/imc"
	"github.com/stretchr/testify/assert"
)

func Test_VowelCounter(t *testing.T) {
	testTable := []struct {
		name     string
		weight   float64
		height   float64
		expected string
	}{
		{
			name:     "sending weight 65 and height 175 expected 13.1",
			weight:   40,
			height:   1.75,
			expected: "IMC: 13.1, Bajo peso",
		},
		{
			name:     "sending weight 65 and height 175 expected 21,2",
			weight:   65,
			height:   1.75,
			expected: "IMC: 21.2, Normal",
		},
		{
			name:     "sending weight 75 and height 165 expected 27.5",
			weight:   75,
			height:   1.65,
			expected: "IMC: 27.5, Sobrepeso",
		},
		{
			name:     "sending weight 100 and height 165 expected 36.7",
			weight:   100,
			height:   1.65,
			expected: "IMC: 36.7, Obesidad",
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := imc.Execute(tc.weight, tc.height)
			assert.Equal(t, tc.expected, got)
		})
	}
}
