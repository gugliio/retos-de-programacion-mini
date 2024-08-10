package imc_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/imc"
	"github.com/stretchr/testify/assert"
)

func Test_VowelCounter(t *testing.T) {
	testTable := []struct {
		name        string
		weight      float64
		height      float64
		expected    float64
		expectedErr error
	}{
		{
			name:     "sending weight 65 and height 175 expected 21,2",
			weight:   65,
			height:   1.75,
			expected: 21.2,
		},
		{
			name:     "sending weight 75 and height 165 expected 22,8",
			weight:   75,
			height:   1.65,
			expected: 27.5,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got, err := imc.Execute(tc.weight, tc.height)
			assert.Equal(t, tc.expected, got)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
