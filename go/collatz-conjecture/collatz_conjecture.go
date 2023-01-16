package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(num int) (int, error) {
	if num == 0 {
		return 0, errors.New("Cannot divide by zero")
	} else if num < 0 {
		return 0, errors.New("Cannot divide negative numbers")
	}

	steps := 0

	for {
		if num == 1 {
			break
		}
		if num % 2 == 0 {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
		steps += 1
	}

	return steps, nil
}