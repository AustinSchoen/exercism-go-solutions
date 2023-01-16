package reverse

import (
	"strings"
	"unicode/utf8"
)

func String(s string) string {
	b := strings.Builder{}
	runeCount := utf8.RuneCountInString(s)
	runes := []rune(s)

	if runeCount > 0 {
		for i := runeCount; i > 0; i-- {
			b.WriteRune(runes[i-1])
		}
	}
	return b.String()
}