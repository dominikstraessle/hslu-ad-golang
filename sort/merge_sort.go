package sort

import (
	"golang.org/x/exp/constraints"
	"sync"
)

const ConcurrentThreshold = 500

func Mergesort[T constraints.Ordered](arr []T) []T {
	b := make([]T, len(arr)+1)

	return mergesort(arr, 0, len(arr)-1, b)
}

func mergesort[T constraints.Ordered](arr []T, left int, right int, b []T) []T {
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

func ConcurrentMergesort[T constraints.Ordered](arr []T) []T {
	b := make([]T, len(arr)+1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	concurrentMergesort(arr, 0, len(arr)-1, b, wg)
	wg.Wait()
	return arr
}

func concurrentMergesort[T constraints.Ordered](arr []T, left int, right int, b []T, quit *sync.WaitGroup) {
	var i, j, k, m int
	if right > left {
		m = (right + left) / 2

		wg := &sync.WaitGroup{}

		if m-left > ConcurrentThreshold {
			wg.Add(1)
			go concurrentMergesort(arr, left, m, b, wg)
		} else {
			mergesort(arr, left, m, b)
		}

		if right-m+1 > ConcurrentThreshold {
			wg.Add(1)
			go concurrentMergesort(arr, m+1, right, b, wg)
		} else {
			mergesort(arr, m+1, right, b)
		}

		wg.Wait()

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
	quit.Done()
}
