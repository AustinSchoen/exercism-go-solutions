//Package space age will calculate how old someone will be on various planets in the solar system.
package space

import (
	"math"
)

type Planet string

//Maybe replace with a Planet struct instead?
const earthYearInSec = 31557600.0
const mercuryYearInSec = earthYearInSec * 0.2408467
const venusYearInSec = earthYearInSec * 0.61519726
const marsYearInSec = earthYearInSec * 1.8808158
const jupiterYearInSec = earthYearInSec * 11.862615
const saturnYearInSec = earthYearInSec * 29.447498
const uranusYearInSec = earthYearInSec * 84.016846
const neptuneYearInSec = earthYearInSec * 164.79132

//Facilitate a value the test case is expecting (String of planet name)
var Planets = map[Planet]float64{
	"Earth":   earthYearInSec,
	"Mercury": mercuryYearInSec,
	"Venus":   venusYearInSec,
	"Mars":    marsYearInSec,
	"Jupiter": jupiterYearInSec,
	"Saturn":  saturnYearInSec,
	"Uranus":  uranusYearInSec,
	"Neptune": neptuneYearInSec,
}

// Age will return years in age with precision 2
func Age(t float64, planet Planet) float64 {
	bigAgeInYears := t / Planets[planet]
	spaceAgeInYears := Round(bigAgeInYears, 0.5, 2)
	return spaceAgeInYears
}

// Thanks to @DavidVaini for Q&D method: https://gist.github.com/DavidVaini/10308388
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}