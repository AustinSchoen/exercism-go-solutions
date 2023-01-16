//Package hamming will calculate the Hamming difference between two different string representations of DNA strands.
package hamming

import "errors"

//Distance will calculate the Hamming difference between two different DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Lengths do not match")
	}

	w := len(a)
	diff := 0

	for i := 0; i < w; i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}