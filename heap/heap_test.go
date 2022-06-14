package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestFixedSizeHeap_Insert(t *testing.T) {
	f := NewFixedSizeHeap[byte](6)
	f.Insert(20)
	f.Insert(10)
	f.Insert(5)
	f.Insert(12)
	f.Insert(7)
	f.Insert(50)
	want := []byte{50, 12, 20, 10, 7, 5}
	assert.Equal(t, want, f.Print())

	f.DeleteRoot()
	want = []byte{20, 12, 5, 10, 7, 0}
	assert.Equal(t, want, f.Print())

	f.DeleteRoot()
	want = []byte{12, 10, 5, 7, 0, 0}
	assert.Equal(t, want, f.Print())

	f.DeleteRoot()
	want = []byte{10, 7, 5, 0, 0, 0}
	assert.Equal(t, want, f.Print())
}

func TestFixedSizeHeap_InsertFull(t *testing.T) {
	t.Skip()
	f := NewFixedSizeHeap[byte](1)
	f.Insert(1)
	f.Insert(1)
}

func TestFixedSizeHeap_InsertSameSize(t *testing.T) {
	f := NewFixedSizeHeap[byte](10)
	f.Insert('a')
	f.Insert('a')
	f.Insert('b')
	f.Insert('a')
	f.Insert('a')
	f.Insert('x')
	f.Insert('a')
	f.Insert('c')
	f.Insert('a')
	f.Insert('x')
	fmt.Println(f.Print())
	want := []byte("xxbacaaaaa")
	assert.Equal(t, want, f.Print())
}

func TestFixedSizeHeap_DeleteEmpty(t *testing.T) {
	f := NewFixedSizeHeap[int](10)
	root := f.DeleteRoot()
	if root != 0 {
		t.Errorf("Expected 0 but got %v", root)
	}
}

func TestHeapsort(t *testing.T) {
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
			//TODO: why is the test failing?
			name: "big slice",
			args: args{
				a: x,
			},
			want: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Heapsort[int](tt.args.a))
		})
	}
}

func BenchmarkHeapsortBigSlice(b *testing.B) {
	var x []int
	for i := 0; i < 100_000; i++ {
		x = append(x, rand.Intn(1_000_000))
	}
	for i := 0; i < b.N; i++ {
		Heapsort[int](x)
	}
}

func BenchmarkHeapsort(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		Heapsort[int](x)
	}
}

func BenchmarkHeapsortOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		Heapsort[int](y)
	}
}

func BenchmarkHeapsortPessimistic(b *testing.B) {
	var z []int
	for i := 10000; i > 0; i-- {
		z = append(z, i)
	}
	for i := 0; i < b.N; i++ {
		Heapsort[int](z)
	}
}

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
