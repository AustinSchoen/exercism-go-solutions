package cards

import "slices"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	for i, v := range slice {
		if i == index {
			return v
		}
	}
	return -1
}

// GetIndex retrieves the index of an item given slice.
// If the index is out of range, we want it to return -1.
func GetIndex(slice []int, index int) int {
	for i, v := range slice {
		if v == index {
			return i
		}
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	getItem := GetItem(slice, index)
	if getItem == -1 {
		return append(slice, value)
	}

	return append(append(slice[0:index], value), slice[index+1:]...)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return slices.Insert(slice, 0, values...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	return append(slice[0:index], slice[index+1:]...)
}
