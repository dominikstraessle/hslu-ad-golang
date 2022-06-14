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

func ConcurrentMergesort[T ad.Ordered](arr []T) []T {
	b := make([]T, len(arr)+1)

	wait := make(chan bool, 1)
	concurrentMergesort(arr, 0, len(arr)-1, b, wait)
	<-wait
	return arr
}

func concurrentMergesort[T ad.Ordered](arr []T, left int, right int, b []T, quit chan bool) {
	var i, j, k, m int
	if right > left {
		m = (right + left) / 2

		wait := make(chan bool, 2)
		numberOfRoutines := 0

		if m-left > 500 {
			go concurrentMergesort(arr, left, m, b, wait)
			numberOfRoutines++
		} else {
			mergesort(arr, left, m, b)
		}

		if right-m+1 > 500 {
			go concurrentMergesort(arr, m+1, right, b, wait)
			numberOfRoutines++
		} else {
			mergesort(arr, m+1, right, b)
		}

		for n := 0; n < numberOfRoutines; n++ {
			select {
			case <-wait:
				continue
			}
		}

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
	quit <- true
}
