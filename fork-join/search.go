package fork_join

import "context"

func SearchIndex[T comparable](arr []T, key T) int {
	for i, t := range arr {
		if t == key {
			return i
		}
	}
	return -1
}

func SearchIndexParallel[T comparable](arr []T, key T) int {
	ctx := context.Background()

	result := make(chan int)
	go searchIndexParallel(ctx, arr, key, 0, len(arr)-1, result)
	return <-result
}

func searchIndexParallel[T comparable](ctx context.Context, arr []T, key T, min, max int, result chan int) {
	if len(arr) > 1000 {
		mid := (max - min) / 2
		newResult := make(chan int, 2)
		newCtx, cancelFunc := context.WithCancel(ctx)
		defer cancelFunc()
		go searchIndexParallel(newCtx, arr, key, min, mid, newResult)
		go searchIndexParallel(newCtx, arr, key, mid, max, newResult)
		for i := 0; i < 2; i++ {
			select {
			case <-ctx.Done():
			case r := <-newResult:
				if r != -1 {
					result <- i
				}
			}
		}

	} else {
		for i := min; i < max; i++ {
			if arr[i] == key {
				result <- i
				return
			}
		}
		result <- -1
	}
}
