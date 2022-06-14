package sort

import (
	"ad"
	"math"
)

func InsertionSort[T ad.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex

		// iterate over the elements in the sorted space
		// if current element is smaller -> switch position
		// repeat until the element on the left is smaller than the current element
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}

		arr[iterator] = temporary
	}
	return arr
}

func ShellSort[T ad.Ordered](arr []T) []T {
	// only use gaps for slices greater than 10_000
	if len(arr) < 10_000 {
		return InsertionSort(arr)
	}

	gaps := getGaps(len(arr))
	for _, gap := range gaps {
		for currentIndex := gap; currentIndex < len(arr); currentIndex++ {
			temporary := arr[currentIndex]
			iterator := currentIndex

			// iterate over the elements in the sorted space
			// if current element is smaller -> switch position
			// repeat until the element on the left is smaller than the current element
			for ; iterator-gap >= 0 && arr[iterator-gap] > temporary; iterator -= gap {
				arr[iterator] = arr[iterator-gap]
			}

			arr[iterator] = temporary
		}
	}
	return arr
}

//getGaps returns gap sizes for the hibbard sequence
func getGaps(sliceSize int) []int {
	nGaps := int(math.Ceil(float64(sliceSize / 10_000)))
	gaps := make([]int, nGaps)
	for i := nGaps; i > 3; i-- {
		gaps = append(gaps, 2^i-1)
	}
	gaps = append(gaps, 1)
	return gaps
}

func SelectionSort[T ad.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		smallestIndex := 0
		smallestElement := temp

		// search for the smallest element in the slice
		// only search the unordered part of the slice -> j := i
		for j := i; j < len(arr); j++ {
			if arr[j] < smallestElement {
				// save the index and element of the smallest element
				smallestIndex = j
				smallestElement = arr[j]
			}
		}

		// if the found element is smaller than the current element -> switch positions
		if arr[i] > smallestElement {
			arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
		}
	}
	return arr
}

func BubbleSort[T ad.Ordered](arr []T) []T {
	// repeat bubble sort as many times as the array has elements
	for j := 0; j < len(arr); j++ {

		// count the number of bubbles for each iteration
		countBubbles := 0

		// for each iteration, iterate over the whole array (minus the last element)
		for i := 0; i < len(arr)-1; i++ {
			// if the element on the left is greater than the next -> switch them
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				countBubbles++
			}
		}

		// if there was an iteration without any bubbles -> array is sorted and can be returned
		// this optimization improves the sort algorithm in best, average and worst case significantly
		// from 65793316 to 6652 ns per operation
		// 9931x speedup
		if countBubbles == 0 {
			return arr
		}
	}
	return arr
}
