//Package euclidian -> Week 01
package euclidian

func ggt(a, b int) int {
	c := 0
	if a > b {
		c = a - b
	} else {
		c = b - a
	}
	if c == b {
		return c
	}
	return ggt(b, c)
}

func ggt2(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	if a > b {
		return ggt2(a%b, b)
	}
	return ggt2(a, b%a)
}
