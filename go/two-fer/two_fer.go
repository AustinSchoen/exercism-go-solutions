// Package twofer will demonstrate sharing
package twofer


// ShareWith will declare who is being shared to.
func ShareWith(s string) (response string) {
	if len(s) > 0 {
		response = "One for " + s + ", one for me."
	} else {
		response = "One for you, one for me."
	}
	return
}