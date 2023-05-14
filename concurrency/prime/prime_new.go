package prime

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math"
	"math/big"
	"runtime"
)

// primesConcurrent generates and prints up to rounds valid prime numbers
func primesConcurrentNew(ctx context.Context, rounds uint32) {
	// create context with cancel function
	ctx, cancelFunc := context.WithCancel(ctx)
	// channel receiving the valid prime numbers
	results := make(chan *big.Int)

	go runWorkOrchestrator(ctx, results)

	// only receive values until the required number of primes is reached
	for i := 0; i < int(rounds); i++ {
		primeNumber := <-results
		fmt.Printf("%v: %s\n", i, primeNumber.String())
	}

	// cancel the context to stop all goroutines
	cancelFunc()
}

// runWorkOrchestrator creates a running prime finder workerFn for each available cpu core until the context gets cancelled
func runWorkOrchestrator(ctx context.Context, results chan *big.Int) {
	// semaphore to control number of active workers
	sema := semaphore.NewWeighted(int64(runtime.NumCPU()))
	// the maximum possible random number -> used as limit for the random number generator
	max := getMaxRandomNumber()
	for {
		// wait and get until a semaphore is available
		err := sema.Acquire(ctx, 1)
		if err != nil {
			fmt.Println(err)
			return
		}

		go workerFn(max, ctx, results, sema)
	}
}

// workerFn tries to find a prime number before the context is cancelled otherwise cancels
func workerFn(max *big.Int, ctx context.Context, results chan *big.Int, sema *semaphore.Weighted) {
	result := make(chan *big.Int, 1)
	tryFindPrime(max, result)

	select {
	case <-ctx.Done():
		// context was cancelled
		return
	case n := <-result:
		// result was available before the context was cancelled
		if n != nil {
			results <- n
		}
	}

	// release the resource
	sema.Release(1)
}

// tryFindPrime generate a random number. Send number to chan else send nil value.
func tryFindPrime(max *big.Int, result chan *big.Int) {
	n := getRandomNumber(max)
	if n.ProbablyPrime(math.MaxInt) {
		// send prime number to channel
		result <- n
	} else {
		result <- nil
	}
}
