package sort

func directInsert(a []int) []int {
	elt := 0
	j := 0

	for i := 1; i < len(a); i++ {
		elt = a[i]
		j = i
		for (j > 0) && (a[j-1] > elt) {
			a[j] = a[j-1]
			j--
		}
		a[j] = elt
	}
	return a
}
