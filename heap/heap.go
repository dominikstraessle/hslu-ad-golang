package heap

import "golang.org/x/exp/constraints"

type IntegerHeap[T constraints.Ordered] interface {
	Insert(e T)
	DeleteRoot() T
	Print() []T
}

func Heapsort[T constraints.Ordered](arr []T) []T {
	h := NewFixedSizeHeap[T](len(arr))
	for _, t := range arr {
		h.Insert(t)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = h.DeleteRoot()
	}

	return arr
}

// only supports bytes
// a zero value is used for empty array entries
type FixedSizeHeap[T constraints.Ordered] struct {
	arr        []T
	lastIndex  int // unused or -1 if full
	emptyValue T
}

func (f *FixedSizeHeap[T]) Print() []T {
	return f.arr
}

func NewFixedSizeHeap[T constraints.Ordered](size int) IntegerHeap[T] {
	return &FixedSizeHeap[T]{
		arr: make([]T, size),
	}
}

func (f *FixedSizeHeap[T]) Insert(e T) {
	if f.lastIndex == -1 {
		panic("heap is full!")
	}

	i := f.lastIndex // current child index

	if f.lastIndex >= len(f.arr)-1 {
		f.lastIndex = -1 // no more elements
	} else {
		f.lastIndex++ // increase last index
	}

	f.arr[i] = e // set element at the end

	if i == 0 {
		return
	}

	f.switchParent(e, i)
}

func (f *FixedSizeHeap[T]) switchParent(e T, i int) {
	j := parentIndex(i)

	if f.arr[j] >= e {
		return // parent is gte current value
	}

	f.arr[j], f.arr[i] = f.arr[i], f.arr[j]

	f.switchParent(e, j)
	return
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return (2 * i) + 1
}

func rightChildIndex(i int) int {
	return 2 * (i + 1)
}

func (f *FixedSizeHeap[T]) DeleteRoot() T {
	e := f.arr[0]
	if e == f.emptyValue || f.lastIndex == 0 {
		return e
	}

	if f.lastIndex == -1 {
		// move root to the last used index and set to 0 (unused)
		f.arr[len(f.arr)-1], f.arr[0] = f.emptyValue, f.arr[len(f.arr)-1]
	} else {
		// move root to the last used index and set to 0 (unused)
		f.arr[f.lastIndex-1], f.arr[0] = f.emptyValue, f.arr[f.lastIndex-1]
	}

	// check and switch children
	f.switchChild(0)

	// update the last index
	if f.lastIndex == -1 {
		f.lastIndex = len(f.arr) - 1
	} else {
		f.lastIndex--
	}
	return e
}

func (f *FixedSizeHeap[T]) switchChild(i int) {
	leftChild, rightChild := leftChildIndex(i), rightChildIndex(i)

	if leftChild >= len(f.arr) {
		return
	}
	if rightChild >= len(f.arr) {
		return
	}

	if f.arr[leftChild] > f.arr[i] {
		if f.arr[leftChild] > f.arr[rightChild] { // left child is greater
			f.arr[i], f.arr[leftChild] = f.arr[leftChild], f.arr[i]
			f.switchChild(leftChild)
			return
		} else { // right child is greater
			f.arr[i], f.arr[rightChild] = f.arr[rightChild], f.arr[i]
			f.switchChild(rightChild)
			return
		}
	}
}
