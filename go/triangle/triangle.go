// Determine the types of triangles
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    NaT = iota // zero unique sides
    Equ
    Iso
    Sca
)

// Determine triangle type provided the length of sides
func KindFromSides(a, b, c float64) Kind {
	// negative or zero length sides indicate NaT
	if a*b*c <= 0 {
		return NaT
	}

	s := []float64{a, b, c}
	sortSlice := sort.Float64Slice{a, b, c}

	//determine triangle equality
	sortSlice.Sort()
	if sortSlice[0] + sortSlice[1] < sortSlice[2] {
		return NaT
	}

	//determine number of unique sides
	seen := make(map[float64]struct{}, len(s))
	j := 0
	for _, v := range s {
		// If any value in s is determine to be InF or NaN, return not a triangle
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return NaT
		}

		// if number is already been seen in the holding map, continue to next iteration
		if _, ok := seen[v]; ok {
			continue
		}

		//Assign key the value of the number from slice S, set value to empty
		seen[v] = struct{}{}
		//Restructure S to include values
		s[j] = v
		j++
	}
	// retrieve only values in S restructured from previous loop
	sides := s[:j]

	switch len(sides) {
	case Equ:
		return Equ
	case Iso:
		return Iso
	case Sca:
		return Sca
	default:
		return NaT
	}
}
