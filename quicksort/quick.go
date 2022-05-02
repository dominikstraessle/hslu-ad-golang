package quicksort

import "math/rand"

func Quicksort(a []byte) []byte {
	if len(a) <= 1 {
		return a
	}

	return quicksort(a, 0, len(a)-1)
}

func quicksort(a []byte, left, right int) []byte {
	up := left        // left border
	down := right - 1 // right border without pivot
	t := a[right]     // right most element as pivot
	allChecked := false

	for allChecked == false {
		for a[up] < t {
			up++ // search greater (>=) element from the left side
		}

		for (a[down] > t) && (down > up) {
			down-- // search lower (<) element from the right side
		}

		if down > up { // only if there is no intersection
			a[up], a[down] = a[down], a[up]
			up++   // move right side
			down-- // move left side
		} else {
			allChecked = true
		}
	}

	a[up], a[right] = a[right], a[up] // pivot element to final position (a[up])
	if left < (up - 1) {
		quicksort(a, left, up-1) // left side
	}
	if (up + 1) < right {
		quicksort(a, up+1, right) // right side
	}
	return a
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
