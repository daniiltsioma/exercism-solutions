package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	fmaps := make([]FreqMap, len(texts))
	for i := range fmaps {
		fmaps[i] = make(map[rune]int)
	}
	messages := make(chan FreqMap, 1)

	for i, text := range texts {
		go func() {
			fm := fmaps[i]
			for _, r := range text {
				fm[r]++
			}
			messages <- fm
		}()
	}

	frequencies := FreqMap{}
	for i := 0; i < len(texts); i++ {
		for k, v := range <-messages {
			frequencies[k] += v
		}
	}
	
	return frequencies
}
