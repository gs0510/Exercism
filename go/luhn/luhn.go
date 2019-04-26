package luhn

import (
	"errors"
	"strings"
	"unicode"
)

func doubleDigit(number int) int {
	doubled := 2 * number
	if doubled > 9 {
		return doubled - 9
	}
	return doubled
}

func calculateSum(input string) (int, error) {
	sum := 0
	index := 0
	for i := len(input) - 1; 0 <= i; i-- {
		if !unicode.IsDigit(rune(input[i])) {
			return -1, errors.New("string has invalid character")
		}
		digit := int(input[i] - '0')
		if index%2 != 0 {
			sum += doubleDigit(digit)
		} else {
			sum += int(digit)
		}
		index++
	}
	return sum, nil
}

// Valid returns if a string is valid based on Luhn's Algorithm.
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	if len(input) < 2 {
		return false
	}
	sum, err := calculateSum(input)
	if err != nil {
		return false
	}
	return sum%10 == 0
}
