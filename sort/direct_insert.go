package sort

import (
	"ad"
)

func DirectInsert[T ad.Ordered](arr []T) []T {
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
