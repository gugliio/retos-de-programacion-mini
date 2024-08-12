package anagrama_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/anagrama"
	"github.com/stretchr/testify/assert"
)

func Test_Anagrama(t *testing.T) {
	testTable := []struct {
		name     string
		word1    string
		word2    string
		expected bool
	}{
		{
			name:     "sending roma and amor return true",
			word1:    "roma",
			word2:    "amor",
			expected: true,
		},
		{
			name:     "sending ramon and amor return false",
			word1:    "ramon",
			word2:    "amor",
			expected: false,
		},
		{
			name:     "sending perro and gato return false",
			word1:    "perro",
			word2:    "gato",
			expected: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := anagrama.Execute(tc.word1, tc.word2)
			assert.Equal(t, tc.expected, got)
		})
	}
}
