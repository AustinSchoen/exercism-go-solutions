// Package gigasecond calculates when a person has been alive one gigasecond
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Second * time.Duration(math.Pow(10, 9))
	newTime := t.Add(gigasecond)
	return newTime
}