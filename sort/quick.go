package sort

import (
	"ad"
	"math/rand"
)

func Quicksort[T ad.Ordered](a []T) []T {
	if len(a) <= 1 {
		return a
	}

	return quicksort(a, 0, len(a)-1, 25)
}

func quicksort[T ad.Ordered](arr []T, left, right, threshold int) []T {
	up := left          // left border
	down := right - 1   // right border without pivot
	pivot := arr[right] // right most element as pivot
	allChecked := false

	for allChecked == false {
		for arr[up] < pivot {
			up++ // search greater (>) element from the left side
		}

		for (arr[down] > pivot) && (down > up) {
			down-- // search lower (<) element from the right side
		}

		if down > up { // only if there is no intersection
			arr[up], arr[down] = arr[down], arr[up]
			up++   // move right side
			down-- // move left side
		} else {
			allChecked = true
		}
	}

	arr[up], arr[right] = arr[right], arr[up] // pivot element to final position (arr[up])
	if left < (up - 1) {
		if len(arr[left:up]) <= threshold {
			InsertionSort(arr[left:up])
		} else {
			quicksort(arr, left, up-1, threshold) // left side
		}
	}
	if (up + 1) < right {
		if len(arr[up+1:right+1]) <= threshold {
			InsertionSort(arr[up+1 : right+1])
		} else {
			quicksort(arr, up+1, right, threshold) // right side
		}
	}
	return arr
}

func RandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return randomBytes
}

// Aufgabe i) polynomial
// Aufgabe i) O(n*ln(n))
// Aufgabe i) schneller
// Aufgabe i) solve(n*ln(n)=2) -> n = 2,345
// Aufgabe i) 4,691 * ln(4,691) -> 7.25s
