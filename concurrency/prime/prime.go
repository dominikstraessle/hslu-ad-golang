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

func primesConcurrent(rounds uint32) {
	nWorkers := runtime.NumCPU()
	ctx := context.Background()
	pool := worker.New(nWorkers)
	sema := semaphore.NewWeighted(int64(nWorkers))
	max := getMaxRandomNumber()
	results := make(chan *big.Int)

	var numberOfPrimes uint32

	go func() {
		for atomic.LoadUint32(&numberOfPrimes) < rounds {
			err := sema.Acquire(ctx, 1)
			if err != nil {
				panic(err)
			}
			pool.Execute(func() {
				randomNumber := getRandomNumber(max)
				if randomNumber.ProbablyPrime(math.MaxInt) {
					results <- randomNumber
					atomic.AddUint32(&numberOfPrimes, 1)
				}
				sema.Release(1)
			})
		}
		close(results)
	}()

	for prime := range results {
		fmt.Printf("%v: %s\n", atomic.LoadUint32(&numberOfPrimes), prime.String())
	}

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
