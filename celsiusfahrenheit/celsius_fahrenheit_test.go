package celsiusfahrenheit_test

import (
	"testing"

	"github.com/gugliio/retos-de-programacion-mini/celsiusfahrenheit"
	"github.com/stretchr/testify/assert"
)

func Test_Farenheit_Celius(t *testing.T) {
	testTable := []struct {
		name     string
		temp     float64
		tempType string
		expected float64
	}{
		{
			name:     "sending invalid type return 0",
			temp:     40,
			tempType: "invalid type",
			expected: 0,
		},
		{
			name:     "sending 40 farenheit return 4.4 celsius",
			temp:     77,
			tempType: "fahrenheit",
			expected: 25,
		},
		{
			name:     "sending 40 celsius return 104 farenheit",
			temp:     40,
			tempType: "celsius",
			expected: 104,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			got := celsiusfahrenheit.Execute(tc.temp, tc.tempType)
			assert.Equal(t, tc.expected, got)
		})
	}
}
