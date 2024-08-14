package media_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/media"
	"github.com/stretchr/testify/assert"
)

func Test_VowelCounter(t *testing.T) {
	testTable := []struct {
		name     string
		numbers  []int
		expected float64
	}{
		{
			name:     "sending weight [10 , 5] expected 7.5",
			numbers:  []int{10, 5},
			expected: 7.5,
		},
		{
			name:     "sending weight [2 , 2] expected 2",
			numbers:  []int{2, 2},
			expected: 2,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := media.Execute(tc.numbers)
			assert.Equal(t, tc.expected, got)
		})
	}
}
