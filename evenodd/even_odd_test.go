package evenodd_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/evenodd"
	"github.com/stretchr/testify/assert"
)

func Test_EvenOdd(t *testing.T) {
	testTable := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "sending 10 we should get par",
			input:    10,
			expected: "par",
		},
		{
			name:     "sending 7 we should get impar",
			input:    7,
			expected: "impar",
		},
		{
			name:     "sending -60 we should get impar",
			input:    -60,
			expected: "impar",
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := evenodd.Execute(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
