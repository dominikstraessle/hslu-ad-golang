package heap

type IntegerHeap interface {
	Insert(e byte)
	Delete(e byte) byte
	Print() string
}

// only supports bytes
// a zero value is used for empty array entries
// duplicate values are not allowed at the moment
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

	f.arr[i] = e

	if i == 0 {
		return
	}

	f.switchParent(e, i)
}

func (f *FixedSizeHeap) switchParent(e byte, i int) {
	j := parentIndex(i)

	if f.arr[j] >= e {
		return
	}

	f.arr[j], f.arr[i] = f.arr[i], f.arr[j]

	leftChild, rightChild := leftChildIndex(j), rightChildIndex(j)

	if leftChild >= len(f.arr) {
		return
	}
	if rightChild >= len(f.arr) {
		return
	}

	if f.arr[leftChild] > f.arr[j] {
		if f.arr[leftChild] > f.arr[rightChild] {
			f.arr[j], f.arr[leftChild] = f.arr[leftChild], f.arr[j]
			f.switchParent(f.arr[j], j)
			return
		} else {
			f.arr[j], f.arr[rightChild] = f.arr[rightChild], f.arr[j]
			f.switchParent(f.arr[j], j)
			return
		}
	}

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

func (f *FixedSizeHeap) Delete(e byte) byte {
	//TODO implement me
	panic("implement me")
}
