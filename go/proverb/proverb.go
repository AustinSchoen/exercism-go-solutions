// Package proverb generate the appropriate proverb given the
package proverb

import "fmt"

// Proverb generate proverb given string
func Proverb(rhyme []string) []string {
	parts := len(rhyme)
	if parts == 0 {
		return []string{}
	}

	proverb := make([]string, parts)
	for i, word := range rhyme {
		if i == (parts - 1) {
			proverb[i] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
			continue
		}

		proverb[i] = fmt.Sprintf("For want of a %v the %v was lost.", word, rhyme[i+1])
	}

	return proverb
}
