package hash

import (
	"ad/list"
	"fmt"
	"hash/fnv"
)

const CAPACITY = 5

type Table[T comparable] struct {
	data [CAPACITY]list.List[T]
}

func New[T comparable]() Hash[T] {
	return &Table[T]{
		data: [CAPACITY]list.List[T]{},
	}
}

func (t *Table[T]) hash(v T) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", v)))

	hashValue := h.Sum64()

	return hashValue % (CAPACITY - 1)
}

func (t *Table[T]) Put(v T) {
	hash := t.hash(v)
	if t.data[hash] == nil {
		t.data[hash] = list.NewLinkedList[T](v)
		return
	}

	// Set semantic -> only insert if not already exists
	if t.data[hash].Contains(v) == false {
		t.data[hash].Insert(v)
	}
}

func (t *Table[T]) Remove(v T) {
	for _, l := range t.data {
		if l == nil {
			continue
		}

		for l.HasNext() {
			if l.Next().Data() == v {
				l.RemoveActual()
				l.Reset()
				break
			}
		}
	}
}

func (t *Table[T]) Contains(v T) bool {
	for _, l := range t.data {
		if l == nil {
			continue
		}

		for l.HasNext() {
			if l.Next().Data() == v {
				l.Remove(v)
				l.Reset()
				return true
			}
		}
		l.Reset()
	}
	return false
}

func (t *Table[T]) GetAll() []T {
	var all []T
	for _, l := range t.data {
		if l == nil {
			continue
		}

		for l.HasNext() {
			all = append(all, l.Next().Data())
		}
		l.Reset()
	}
	return all
}

type Hash[T comparable] interface {
	Put(T)
	Remove(T)
	Contains(T) bool
	GetAll() []T
}
