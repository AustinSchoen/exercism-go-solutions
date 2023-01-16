//Package acronym will return an acronym from some tech jargon
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate will return the abbreviated version of some tech jargon
func Abbreviate(s string) (acronym string) {

	words := getWords(s)

	for _, word := range words {
		acronym = acronym + string(getFirstLetter(word))
	}
	return strings.ToUpper(acronym)
}

// Get all the words in the jargon
func getWords(s string) (words []string) {
	words = regexp.MustCompile(`\w+`).FindAllString(s, -1)
	return
}

// Get the first letter from a word in the jargon
func getFirstLetter(word string) (letter string) {
	for _, runeValue := range word {
		letter = string(runeValue)
		break
	}
	return
}