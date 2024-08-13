package revertstring_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/revertstring"
	"github.com/stretchr/testify/assert"
)

func Test_VowelCounter(t *testing.T) {
	testTable := []struct {
		name     string
		word     string
		expected string
	}{
		{
			name:     "given the word perro return orrep",
			word:     "perro",
			expected: "orrep",
		},
		{
			name:     "given the word canciones return senoicnac",
			word:     "canciones",
			expected: "senoicnac",
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := revertstring.Execute(tc.word)
			assert.Equal(t, tc.expected, got)
		})
	}
}
