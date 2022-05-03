package heap

type IntegerHeap interface {
	Insert(e byte)
	DeleteRoot() byte
	Print() string
}

// only supports bytes
// a zero value is used for empty array entries
type FixedSizeHeap struct {
	arr       []byte
	lastIndex int // unused or -1 if full
}

func (f *FixedSizeHeap) Print() string {
	return string(f.arr)
}

func NewFixedSizeHeap(size int) IntegerHeap {
	return &FixedSizeHeap{
		arr: make([]byte, size),
	}
}

func (f *FixedSizeHeap) Insert(e byte) {
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

func (f *FixedSizeHeap) switchParent(e byte, i int) {
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

func (f *FixedSizeHeap) DeleteRoot() byte {
	e := f.arr[0]
	if e == 0 || f.lastIndex == 0 {
		return e
	}

	if f.lastIndex == -1 {
		// move root to the last used index and set to 0 (unused)
		f.arr[len(f.arr)-1], f.arr[0] = 0, f.arr[len(f.arr)-1]
	} else {
		// move root to the last used index and set to 0 (unused)
		f.arr[f.lastIndex-1], f.arr[0] = 0, f.arr[f.lastIndex-1]
	}

	f.switchChild(0)

	if f.lastIndex == -1 {
		f.lastIndex = len(f.arr) - 1
	} else {
		f.lastIndex--
	}
	return e
}

func (f *FixedSizeHeap) switchChild(i int) {
	leftChild, rightChild := leftChildIndex(i), rightChildIndex(i)

	if leftChild >= len(f.arr) {
		return
	}
	if rightChild >= len(f.arr) {
		return
	}

	if f.arr[leftChild] > f.arr[i] {
		if f.arr[leftChild] > f.arr[rightChild] {
			f.arr[i], f.arr[leftChild] = f.arr[leftChild], f.arr[i]
			f.switchChild(leftChild)
			return
		} else {
			f.arr[i], f.arr[rightChild] = f.arr[rightChild], f.arr[i]
			f.switchChild(rightChild)
			return
		}
	}
}
