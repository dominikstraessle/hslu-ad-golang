package sort

import (
	"ad"
	"math/rand"
)

const ConcurrentQuicksortThreshold = 50_000

func Quicksort[T ad.Ordered](a []T) []T {
	if len(a) <= 1 {
		return a
	}

	return quicksort(a, 0, len(a)-1, 25)
}

func quicksort[T ad.Ordered](arr []T, left, right, threshold int) []T {
	up := left        // left border
	down := right - 1 // right border without pivot
	pivot := medianOfThree(arr, left, right)
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

//ConcurrentQuicksort sorts the sub-sequences concurrently when there length is above the ConcurrentQuicksortThreshold
//I couldn't win any performance gains yet -> maybe check the implementation or the threshold again
//A problem could be, that the current implementation shares the same array but operates on different parts of it
//This array may have to be switched around in the processor for each running goroutine -> a fork-join could prevent this problem
func ConcurrentQuicksort[T ad.Ordered](a []T) []T {
	if len(a) <= 1 {
		return a
	}

	waitForChannel := make(chan bool, 1)
	concurrentQuicksort(a, 0, len(a)-1, 25, waitForChannel)
	<-waitForChannel
	return a
}

func concurrentQuicksort[T ad.Ordered](arr []T, left, right, threshold int, quitChannel chan bool) {
	up := left        // left border
	down := right - 1 // right border without pivot
	pivot := medianOfThree(arr, left, right)

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

	waitFor := 0
	waitForChannel := make(chan bool, 2)
	if left < (up - 1) {
		if len(arr[left:up]) <= threshold {
			InsertionSort(arr[left:up])
		} else if len(arr[left:up]) <= ConcurrentQuicksortThreshold {
			quicksort(arr, left, up-1, threshold)
		} else {
			waitFor++
			go concurrentQuicksort(arr, left, up-1, threshold, waitForChannel) // left side
		}
	}
	if (up + 1) < right {
		if len(arr[up+1:right+1]) <= threshold {
			InsertionSort(arr[up+1 : right+1])
		} else if len(arr[up+1:right+1]) <= ConcurrentQuicksortThreshold {
			quicksort(arr, up+1, right, threshold)
		} else {
			waitFor++
			go concurrentQuicksort(arr, up+1, right, threshold, waitForChannel) // right side
		}
	}

	for i := 0; i < waitFor; i++ {
		select {
		case <-waitForChannel:
			continue
		}
	}

	quitChannel <- true
}

func medianOfThree[T ad.Ordered](arr []T, left int, right int) T {
	// median of three
	middle := right / 2
	if arr[left] < arr[right] && arr[left] > arr[middle] {
		arr[right], arr[left] = arr[left], arr[right]
	} else if arr[middle] < arr[left] && arr[middle] > arr[right] {
		arr[right], arr[middle] = arr[middle], arr[right]
	}
	return arr[right] // right most element as pivot
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
