package vowelcounter_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/vowelcounter"
	"github.com/stretchr/testify/assert"
)

func Test_VowelCounter(t *testing.T) {
	testTable := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "sending 'Hola Mundo!' should return 4",
			input:    "Hola Mundo!",
			expected: 4,
		},
		{
			name:     "sending 'Una prueba' should return 5",
			input:    "Una prueba",
			expected: 5,
		},
		{
			name:     "sending '12345' should return 0",
			input:    "12345",
			expected: 0,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := vowelcounter.Execute(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
