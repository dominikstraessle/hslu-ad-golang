package sort

import (
	"ad"
)

func InsertionSort[T ad.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}

func SelectionSort[T ad.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		smallestIndex := 0
		smallestElement := temp
		for j := i; j < len(arr); j++ {
			if arr[j] < smallestElement {
				smallestIndex = j
				smallestElement = arr[j]
			}
		}
		if arr[i] > smallestElement {
			arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
		}
	}
	return arr
}
