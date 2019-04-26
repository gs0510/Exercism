package letter

//FreqMap stores the count of a character in the string
type FreqMap map[rune]int

//Frequency calculates the count of characters in a string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency calculates the count of characters in a string concurrently.
func ConcurrentFrequency(input []string) FreqMap {
	bufferSize := 10
	channel := make(chan map[rune]int, bufferSize)
	for _, element := range input {
		go func(v string) {
			channel <- Frequency(v)
		}(element)
	}

	frequency := make(map[rune]int)

	for range input {
		for key, value := range <-channel {
			frequency[key] += value
		}
	}
	return frequency
}
