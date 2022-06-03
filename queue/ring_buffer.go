package queue

import "errors"

type RingBuffer[T any] interface {
	Offer(*T) error
	Poll() (*T, error)
}

const SIZE = 8

type foo struct {
	byte
}

type ArrayRingBuffer[T any] struct {
	head uint
	tail uint
	data [SIZE]*T
}

func (a *ArrayRingBuffer[T]) Offer(t *T) error {
	if a.head == a.tail && a.data[a.head] != nil {
		return errors.New("failed to offer to queue: queue is full")
	}
	a.data[a.head] = t
	a.head = (a.head + 1) % SIZE
	return nil
}

func (a *ArrayRingBuffer[T]) Poll() (*T, error) {
	if a.head == a.tail && a.data[a.tail] == nil {
		return nil, errors.New("failed to poll: queue is empty")
	}
	t := a.data[a.tail]
	a.data[a.tail] = nil
	a.tail = (a.tail + 1) % SIZE
	return t, nil
}
