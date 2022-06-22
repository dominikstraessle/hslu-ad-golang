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

	// i -> text index
	// j -> pattern found index
	i, j := 0, 0
	for {
		// there is a match, increase j
		if text[i+j] == pattern[j] { // match
			j++
		} else {
			// check if it is still possible to make a match
			if (i + m) < n { // mismatch
				// shift the text index to the value from the shift array
				// the value from the shift array is the character at the current text index plus the length of the pattern
				i += shift[text[i+m]]
				// reset the pattern found index to 0
				j = 0
			} else {
				// no match is possible / no shift is possible
				break
			}
		}

		// end of the text is reached
		if !(j < m) || !((i + m) <= n) {
			break
		}
	}

	// there was a match
	if j == m {
		return i
	}

	// no match
	return -1
}
