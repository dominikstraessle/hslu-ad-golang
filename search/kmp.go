package search

func initNext(pattern string) []int {
	m := len(pattern)

	// the length of the pattern plus the first index is -1
	next := make([]int, m+1)
	next[0] = -1 // first index is -1

	i := 0  // index for the loop
	j := -1 // start length with -1

	for {
		// if the pattern matches the beginning
		if (j == -1) || (pattern[i] == pattern[j]) {
			i++         // increase the index
			j++         // increase the length of the match
			next[i] = j // the current index gets the length of the match
		} else {
			j = next[j] // j resets to the next value at where j was
			// if j is 0 -> j becomes -1 again
		}

		// if i becomes greater than the length - 1 -> you are finished
		if i > (m - 1) {
			break
		}
	}

	return next
}

func kmpSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m > n {
		// cannot find a pattern longer than the text
		return -1
	}

	i, j := 0, 0
	next := initNext(pattern)

	for {
		// pass when j equals -1 or when the pattern matches the text at the given position
		// -1 means the pattern has to be searched from the beginning
		if (j == -1) || (text[i] == pattern[j]) {
			i++
			j++
		} else {
			// when the pattern is not found j becomes the value of the next array at the current position
			// if j equals 0 it becomes -1
			j = next[j]
		}

		// stop when the indexes are greater than the text or pattern length
		// j = len(pattern) means that the pattern was found
		if !(j < m) || !(i < n) {
			break
		}
	}

	// j equals the length of the pattern -> the pattern was found
	if j == m {
		// return the first index of the found pattern
		return i - m
	}

	// pattern not found
	return -1
}
