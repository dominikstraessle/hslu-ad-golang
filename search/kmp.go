package search

func initNext(pattern string) []int {
	m := len(pattern)

	// the length of the pattern plus the first index is -1
	next := make([]int, m+1)
	next[0] = -1 // first index is -1

	i := 0  // index for the loop
	j := -1 // start length with -1

	for true {
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
			return next
		}
	}

	return next
}
