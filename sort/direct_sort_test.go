package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func getRandomizedIntSliceAndSortedSolution() ([]int, []int) {
	var x []int
	size := 10000
	for i := 0; i < size; i++ {
		x = append(x, rand.Intn(size))
	}
	y := make([]int, size)
	copy(y, x)
	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})
	return x, y
}

func TestInsertionSort(t *testing.T) {
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
			if got := InsertionSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
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
			if got := SelectionSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestShellSort(t *testing.T) {
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
			if got := ShellSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
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
			if got := BubbleSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		InsertionSort(x)
	}
}

//BenchmarkInsertionSortBigSlice shows the performance loss for big slices compared to shell sort
func BenchmarkInsertionSortBigSlice(b *testing.B) {
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		InsertionSort(x)
	}
}

func BenchmarkInsertionSortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		InsertionSort(y)
	}
}

func BenchmarkInsertionSortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		InsertionSort(z)
	}
}

//BenchmarkShellSortBigSlice shows the performance gain for big slices compared to insertion sort
func BenchmarkShellSortBigSlice(b *testing.B) {
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		ShellSort(x)
	}
}

func BenchmarkShellSort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		ShellSort(x)
	}
}

func BenchmarkShellSortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		ShellSort(y)
	}
}

func BenchmarkQuicksortBigSlice(b *testing.B) {
	b.Skip()
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		Quicksort(x)
	}
}

func BenchmarkQuicksort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		Quicksort(x)
	}
}

func BenchmarkQuicksortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		Quicksort(y)
	}
}

func BenchmarkQuicksortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		Quicksort(z)
	}
}

func BenchmarkConcurrentQuicksortBigSlice(b *testing.B) {
	b.Skip()
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(x)
	}
}

func BenchmarkConcurrentQuicksort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(x)
	}
}

func BenchmarkConcurrentQuicksortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(y)
	}
}

func BenchmarkConcurrentQuicksortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(z)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		SelectionSort(x)
	}
}

func BenchmarkSelectionSortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		SelectionSort(y)
	}
}

func BenchmarkSelectionSortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		SelectionSort(z)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(x)
	}
}

func BenchmarkBubbleSortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(y)
	}
}

func BenchmarkBubbleSortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(z)
	}
}
