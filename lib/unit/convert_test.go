package unit

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	scenarios := []struct {
		input  string
		number float64
		err    error
	}{
		{"", 0, fmt.Errorf("convert: invalid input: ``")},
		{"123", float64(123), nil},
		{"123K", 123 * math.Pow(10, 3), nil},
		{"123M", 123 * math.Pow(10, 6), nil},
		{"123G", 123 * math.Pow(10, 9), nil},
		{"123T", 123 * math.Pow(10, 12), nil},
		{"123P", 123 * math.Pow(10, 15), nil},
		{"123E", 0, fmt.Errorf("convert: invalid unit: `E`")},
		{"1.23K", 1.23 * math.Pow(10, 3), nil},
		{"1.23M", 1.23 * math.Pow(10, 6), nil},
		{"1.23G", 1.23 * math.Pow(10, 9), nil},
		{"1.23T", 1.23 * math.Pow(10, 12), nil},
		{"1.23P", 1.23 * math.Pow(10, 15), nil},
		{"1.23E", 0, fmt.Errorf("convert: invalid unit: `E`")},
	}

	for _, scenario := range scenarios {
		number, err := convert(scenario.input)

		if scenario.err == nil {
			require.NoError(t, err)
		} else {
			require.Equal(t, scenario.err, err)
		}

		require.Equal(t, scenario.number, number)
	}
}
