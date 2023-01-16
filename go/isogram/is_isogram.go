package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether the phrase is an isogram
func IsIsogram(phrase string) bool {
	normalize := func (r rune) rune {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
		} else {
			return -1
		}
		return r
	}

	phrase = strings.Map(normalize, phrase)

	for i := range phrase {
		if i != strings.LastIndexByte(phrase, phrase[i]) {
			return false
		}
	}
	return true
}