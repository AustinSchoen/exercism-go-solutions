// Package bob responds based on what is said to bob.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)


// Hey is a method for speaking to bob
func Hey(remark string) (response string) {

	// 1. respond with sure if remark ends in a question mark
	// 2. respond with chill out if remark is in caps [A-Z]+
	// 3. respond with calm down if question is yelled at bob
	// 4. respond with be that way if bob is addressed without something being said (empty string)
	// 5. respond with whatever to anything else

	if isYell(remark) && isQuestion(remark) {
		response = "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		response = "Sure."
	} else if isYell(remark) {
		response = "Whoa, chill out!"
	} else if isSilent(remark) {
		response = "Fine. Be that way!"
	} else {
		response = "Whatever."
	}

	return
}

func isYell(remark string) (assertion bool) {
	patternUpper, _ := regexp.Compile(`[[:upper:]]{2,}`)
	patternLower, _ := regexp.Compile(`[[:lower:]]+`)
	hasUpper := patternUpper.MatchString(remark)
	hasLower := patternLower.MatchString(remark)
	if hasLower {
		assertion = false
	} else if hasUpper {
		assertion = true
	} else {
		assertion = false
	}
	return
}

func isQuestion(remark string) (assertion bool) {
	// Questions end in a question mark - strip newlines & white space
	strippedRemark := strings.Replace(remark, " ", "", -1)
	strippedRemark = strings.Replace(strippedRemark, "\n", "", -1)

	containsMark := regexp.MustCompile(`\?+`).MatchString(strippedRemark)
	wordsBefore  := regexp.MustCompile(`([[:alnum:]]|[[:graph:]])\?`).MatchString(strippedRemark)
	wordsAfter   := regexp.MustCompile(`\?([[:alnum:]]|[[:graph:]])`).MatchString(strippedRemark)

	if (containsMark && wordsBefore) && !(wordsAfter) {
		assertion = true
	} else {
		assertion = false
	}
	return
}

func isSilent(remark string) (assertion bool) {
	pattern, _ := regexp.Compile(`^[[:^alnum:]]*$`)
	assertion = pattern.MatchString(remark)
	return
}