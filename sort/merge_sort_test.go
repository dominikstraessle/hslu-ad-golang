package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMergesort(t *testing.T) {
	x, y := getRandomizedIntSliceAndSortedSolution()
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "inorder",
			args: args{
				a: []int{0, 1, 2, 3, 4, 5},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				a: []int{5, 4, 3, 2, 1, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "sort",
			args: args{
				a: []int{2, 3, 1, 4, 5, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "big slice",
			args: args{
				a: x,
			},
			want: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Mergesort(tt.args.a))
		})
	}
}

func BenchmarkMergesortBigSlice(b *testing.B) {
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		Mergesort(x)
	}
}

func BenchmarkMergesort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		Mergesort(x)
	}
}

func BenchmarkMergesortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		Mergesort(y)
	}
}

func BenchmarkMergesortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		Mergesort(z)
	}
}

func TestConcurrentMergesort(t *testing.T) {
	x, y := getRandomizedIntSliceAndSortedSolution()
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "inorder",
			args: args{
				a: []int{0, 1, 2, 3, 4, 5},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				a: []int{5, 4, 3, 2, 1, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "sort",
			args: args{
				a: []int{2, 3, 1, 4, 5, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "big slice",
			args: args{
				a: x,
			},
			want: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ConcurrentMergesort(tt.args.a))
		})
	}
}

func BenchmarkConcurrentMergesortBigSlice(b *testing.B) {
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		ConcurrentMergesort(x)
	}
}

func BenchmarkConcurrentMergesort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		ConcurrentMergesort(x)
	}
}

func BenchmarkConcurrentMergesortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		ConcurrentMergesort(y)
	}
}

func BenchmarkConcurrentMergesortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		ConcurrentMergesort(z)
	}
}
