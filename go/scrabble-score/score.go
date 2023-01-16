package scrabble

import "strings"

const onePoints   = "aeioulnrst"
const twoPoints   = "dg"
const threePoints = "bcmp"
const fourPoints  = "fhvwy"
const fivePoints  = "k"
const eightPoints = "jx"
const tenPoints   = "qz"

// This is a comment explaining this thing calculates a scrabble score
func Score(word string) int {
	var score int
	word = strings.ToLower(word)
	for _, v := range word {
		switch {
		case strings.ContainsRune(onePoints, v):
			score ++ 
		case strings.ContainsRune(twoPoints, v):
			score += 2
		case strings.ContainsRune(threePoints, v):
			score += 3
		case strings.ContainsRune(fourPoints, v):
			score += 4
		case strings.ContainsRune(fivePoints, v):
			score += 5
		case strings.ContainsRune(eightPoints, v):
			score += 8
		case strings.ContainsRune(tenPoints, v):
			score += 10
		}
	}
	return score
}