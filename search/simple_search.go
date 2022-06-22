package search

func SimpleSearch(text string, pattern string) int {
	maxIndex := len(text) - len(pattern)

	for i := 0; i <= maxIndex; i++ {
		success := true
		for j := 0; j < len(pattern); j++ {
			if text[i+j] != pattern[j] {
				success = false
				break
			}
		}
		if success {
			return i
		}
	}
	return -1
}
