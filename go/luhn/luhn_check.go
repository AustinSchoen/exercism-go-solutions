package luhn

import (
	"unicode"
	"unicode/utf8"
	"strings"
)

func Valid(s string) bool {
	// strip spaces
	s = stripSpaces(s)
	l := utf8.RuneCountInString(s)

	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	if l <= 1 {
		return false
	}

	digits := make([]int, l)

	//convert individual runes to integers
	for i, r := range s {
		digits[i] = int(r - '0')
	}

	//double every 2nd digit from the right
	for di := len(digits) - 2; di >= 0; di = di - 2 {
		digits[di] = digits[di] * 2
		// decrease by 9 if over 9
		if digits[di] > 9 {
			digits[di] = digits[di] - 9
		}
	}

	sum := 0
	for _, digit := range digits {
		sum += digit
	}

	if sum % 10 == 0 {
		return true
	}
	return false
}

func stripSpaces(s string) string {

	isSpace := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}
	return strings.Map(isSpace, s)
}