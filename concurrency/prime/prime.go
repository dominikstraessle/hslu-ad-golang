package prime

import (
	"ad/concurrency/worker"
	"context"
	"crypto/rand"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math"
	"math/big"
	"runtime"
	"sync/atomic"
)

//primesConcurrent generates and prints up to rounds valid prime numbers
func primesConcurrent(rounds uint32) {
	// number of workers
	nWorkers := runtime.NumCPU()
	ctx := context.Background()
	pool := worker.New(nWorkers)

	// semaphore to control number of active workers
	sema := semaphore.NewWeighted(int64(nWorkers))
	// the maximum possible random number -> used as limit for the random number generator
	max := getMaxRandomNumber()

	// channel receiving the valid prime numbers
	results := make(chan *big.Int)

	// the number of validated prime numbers
	var numberOfPrimes uint32

	go func() {
		// repeat as long as the wanted number of prime numbers is reached
		for atomic.LoadUint32(&numberOfPrimes) < rounds {
			// wait and get until a semaphore is available
			err := sema.Acquire(ctx, 1)
			if err != nil {
				panic(err)
			}

			// run the prime number finder in the worker pool
			pool.Execute(func() {
				randomNumber := getRandomNumber(max)
				if randomNumber.ProbablyPrime(math.MaxInt) {
					// send prime number to channel
					results <- randomNumber
					// increment the number of valid primes
					atomic.AddUint32(&numberOfPrimes, 1)
				}

				// release the resource
				sema.Release(1)
			})
		}

		// number of prime numbers is reached -> close the channel
		close(results)
	}()

	// print all prime numbers from the channel
	for prime := range results {
		fmt.Printf("%v: %s\n", atomic.LoadUint32(&numberOfPrimes), prime.String())
	}

	// stop the worker pool
	pool.Stop()
}

func primes(rounds int) {
	max := getMaxRandomNumber()

	n := 0
	for n <= rounds {
		//Generate cryptographically strong pseudo-random between 0 - max
		randomNumber := getRandomNumber(max)
		if randomNumber.ProbablyPrime(math.MaxInt) {
			fmt.Printf("%v: %s\n", n, randomNumber.String())
			n++
		}
	}
}

//getMaxRandomNumber from https://stackoverflow.com/a/45428754/7130107
func getMaxRandomNumber() *big.Int {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(1024), nil).Sub(max, big.NewInt(1))
	return max
}

func getRandomNumber(max *big.Int) *big.Int {
	randomNumber, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return randomNumber
}
