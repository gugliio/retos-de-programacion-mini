package decimaltobinary_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/decimaltobinary"
	"github.com/stretchr/testify/assert"
)

func Test_DecimalToBinary(t *testing.T) {
	testTable := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "sending 12 we should get 1100",
			input:    12,
			expected: "1100",
		},
		{
			name:     "sending 0 we should get 0",
			input:    0,
			expected: "0",
		},
		{
			name:     "sending 50 we should get 110010",
			input:    50,
			expected: "110010",
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := decimaltobinary.Execute(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
