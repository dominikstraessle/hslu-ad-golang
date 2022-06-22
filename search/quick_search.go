package search

func quickSearch(text, pattern string) int {
	n, m := len(text), len(pattern)

	// pattern cannot be longer than text
	if m > n {
		return -1
	}

	asciiRange := 256 // range of ascii chars
	shift := make([]int, asciiRange)

	// init the shift array
	// value for each ascii position becomes index + length of the pattern
	for i := 0; i < asciiRange; i++ {
		shift[i] = m + i
	}

	// overwrite fields according pattern
	// only the values in the shift array which have an occurrence in the pattern ar overwritten
	// the new value becomes length of the pattern minus the current index
	// these values are not yet used by the shift pattern
	for i := 0; i < m; i++ {
		shift[pattern[i]] = m - i
	}

	i, j := 0, 0
	for {
		if text[i+j] == pattern[j] {
			j++
		} else {
			if (i + m) < n {
				i += shift[text[i+m]]
				j = 0
			} else {
				break
			}
		}

		if !(j < m) || !((i + m) <= n) {
			break
		}
	}

	if j == m {
		return i
	}

	return -1
}
