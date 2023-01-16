package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, len(texts))
	result := FreqMap{}

	for _, v := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(v)
	}

	for i := 0; i < len(texts); i++ {
		m := <-c
		for k, v := range m {
			// fmt.Printf("Key: %q, Value: %v\n", k, v)
			result[k] = result[k] + v
		}
	}

	return result
}