package prime

import (
	"context"
	"testing"
)

func Test_primesConcurrentNew(t *testing.T) {
	primesConcurrentNew(context.Background(), 10)
}

func Benchmark_primesConcurrentNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesConcurrentNew(context.Background(), 20)
	}
}
