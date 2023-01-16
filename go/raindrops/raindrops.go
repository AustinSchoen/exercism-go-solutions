package raindrops

import "strconv"

// Convert takes a number and returns Rain Speak or a number
func Convert(num int) string {
	var factors []int
	var rainSpeak string
	var conversion string

	for i := 1; i <= num; i++ {
		r := num % i
		if r == 0 {
			factors = append(factors, i)
		}
	}

	for _, v := range factors {
		switch v {
		case 3:
			rainSpeak = rainSpeak + "Pling"
		case 5:
			rainSpeak = rainSpeak + "Plang"
		case 7:
			rainSpeak = rainSpeak + "Plong"
		}

	}

	if rainSpeak == "" {
		conversion = strconv.Itoa(num)
	} else {
		conversion = rainSpeak
	}

	return conversion
}