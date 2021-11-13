package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	// Every time we encounter a rune, we increase frequency in the map by 1
	for _, r := range s {
		m[r]++
	}
	return m
}

// Concurrent version of above
func ConcurrentFrequency(texts []string) FreqMap{
	// Create a channel to receive the results of the goroutines
	c := make(chan FreqMap)
	
	// For collecting results from Channel
	m := FreqMap{}

	// Create a function that can be invoked as a goroutine
	// Returns results to our channel
	f := func(s string) {
		c <- Frequency(s)
	}

	// Invoke goroutine for each text
	for _ , text := range texts {
		go f(text)

		// For all the runes per text
		for k, v := range <- c {
			m[k] += v
		}
	}
	return m
}