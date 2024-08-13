package factorial_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/factorial"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	testTable := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "sending 5 we should get 120",
			input:    5,
			expected: 120,
		},
		{
			name:     "sending 7 we should get 5040",
			input:    7,
			expected: 5040,
		},
		{
			name:     "sending 0 we should get 1",
			input:    0,
			expected: 1,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := factorial.Execute(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}

}
