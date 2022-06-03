package stack

import "errors"

type Stack[T any] interface {
	Push(*T) error
	Pop() (*T, error)
}

//SIZE of the array
const SIZE = 5

type ArrayStack[T any] struct {
	data           [SIZE]*T
	lastEmptyIndex uint
}

func NewStack[T any]() Stack[T] {
	return &ArrayStack[T]{
		data:           [SIZE]*T{},
		lastEmptyIndex: 0,
	}
}

func (a *ArrayStack[T]) Push(t *T) error {
	if int(a.lastEmptyIndex) >= len(a.data) {
		return errors.New("cannot push: stack is already full")
	}
	a.data[a.lastEmptyIndex] = t
	a.lastEmptyIndex++
	return nil
}

func (a *ArrayStack[T]) Pop() (*T, error) {
	if a.lastEmptyIndex <= 0 {
		return nil, errors.New("cannot pop: stack is empty")
	}

	a.lastEmptyIndex--
	t := a.data[a.lastEmptyIndex]
	return t, nil
}
