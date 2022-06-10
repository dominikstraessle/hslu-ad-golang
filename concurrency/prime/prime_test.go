package prime

import "testing"

func Test_prime(t *testing.T) {
	primes(10)
}

func Test_primesConcurrent(t *testing.T) {
	primesConcurrent(10)
}

func Benchmark_prime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(20)
	}
}

func Benchmark_primesConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(20)
	}
}
