// Determines if provided year is a leap year
package leap

import "math"

// Returns a boolean value if given year is a leap year
func IsLeapYear(year int) bool {
	leapYear := false
	fYear := float64(year)

	if math.Mod(fYear, 4) == 0 {
		leapYear = true
		if math.Mod(fYear, 100) == 0 {
			leapYear = false
			if math.Mod(fYear, 400) == 0 {
				leapYear = true
			}
		}
	}

	return leapYear
}
