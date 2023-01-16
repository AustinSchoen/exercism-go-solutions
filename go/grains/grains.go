package grains

import "errors"

func Square(s int) (uint64, error) {
	var grains uint64 = 1

	if s < 1 || s > 64  {
		return 0, errors.New("Can only specify 1-64 squares")
	}

	for i := 1; i < s; i++ {
		grains = grains * 2
	}

	return grains, nil
}

func Total() uint64 {
	squares := 64

	var sum uint64 = 0

	for i := 1; i <= squares; i++ {
		count, _ :=  Square(i)
		sum += count
	}

	return uint64(sum)
}