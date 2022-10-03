package search

func BinarySearch(list []int, key int) int {
	l := len(list)
	if l == 0 {
		return -1
	}
	if key > list[l-1] || key < list[0] {
		return -1
	}
	return binarySearch(list, key, 0, l-1)
}

func binarySearch(list []int, key, min, max int) int {
	m := int((min + max) / 2)
	if list[m] < key {
		return binarySearch(list, key, m, max)
	} else if list[m] > key {
		return binarySearch(list, key, min, m)
	}
	return m
}
