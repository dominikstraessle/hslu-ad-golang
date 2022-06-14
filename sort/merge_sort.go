package sort

import "ad"

func Mergesort[T ad.Ordered](arr []T) []T {
	b := make([]T, len(arr)+1)

	return mergesort(arr, 0, len(arr)-1, b)
}

func mergesort[T ad.Ordered](arr []T, left int, right int, b []T) []T {
	var i, j, k, m int
	if right > left {
		m = (right + left) / 2
		mergesort(arr, left, m, b)
		mergesort(arr, m+1, right, b)

		for i = left; i <= m; i++ {
			b[i] = arr[i]
		}
		for j = m; j < right; j++ {
			b[right+m-j] = arr[j+1]
		}

		i, j = left, right
		for k = left; k <= right; k++ {
			if b[i] <= b[j] {
				arr[k] = b[i]
				i++
			} else {
				arr[k] = b[j]
				j--
			}
		}

	}
	return arr
}
