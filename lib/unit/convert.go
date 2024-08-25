package unit

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	units             = "KMGTP"
	invalidInputError = "convert: invalid input: `%v`"
	invalidUnitError  = "convert: invalid unit: `%v`"
)

func convert(s string) (float64, error) {
	if s == "" {
		return 0, fmt.Errorf(invalidInputError, s)
	}

	number, err := strconv.ParseFloat(s, 64)

	if err == nil {
		return float64(number), nil
	}

	s = strings.ToUpper(s)
	unit := string(s[len(s)-1])
	unitIndex := strings.Index(s, unit)
	unitPower := float64(strings.Index(units, unit) + 1)

	if _, err := strconv.Atoi(unit); unitPower == 0 && err != nil {
		return 0, fmt.Errorf(invalidUnitError, unit)
	}

	stringNumber := s[:unitIndex]
	number, err = strconv.ParseFloat(stringNumber, 64)

	if err != nil {
		return 0, err
	}

	return number * math.Pow(10, unitPower*3), nil
}
