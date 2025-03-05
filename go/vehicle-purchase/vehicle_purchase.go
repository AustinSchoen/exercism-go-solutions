package purchase

import (
	"slices"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if c := []string{"car", "truck"}; slices.Contains(c, kind) {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	}

	return option1 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	discount := 0.0

	if age < 3 {
		discount = .8
	} else if age >= 10 {
		discount = .5
	} else {
		discount = .7
	}

	return originalPrice * discount
}
