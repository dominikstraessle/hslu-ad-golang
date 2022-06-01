package list

import "errors"

type Data string

type Node[T comparable] struct {
	nextNode *Node[T]
	data     T
}

func (n *Node[T]) next() *Node[T] {
	return n.nextNode
}

func (n *Node[T]) hasNext() bool {
	return n.nextNode != nil
}

func (n *Node[T]) Data() T {
	return n.data
}

type LinkedList[T comparable] struct {
	head     *Node[T]
	actual   *Node[T]
	previous *Node[T]
}

func NewLinkedList[T comparable](data T) List[T] {
	return &LinkedList[T]{
		head: &Node[T]{
			nextNode: nil,
			data:     data,
		},
		actual:   nil,
		previous: nil,
	}
}

func (l *LinkedList[T]) Insert(data T) {
	node := &Node[T]{
		nextNode: l.head,
		data:     data,
	}

	l.head = node
}

func (l *LinkedList[T]) Remove(data T) {
	if l.head.hasNext() == false {
		panic("cannot remove head")
	}
	if l.head.Data() == data {
		l.head = l.head.next()
	}

	prev, actual := recursiveRemove(l.head, l.head.next(), data)
	if actual == nil {
		return
	}

	prev.nextNode = actual.next()
}

func recursiveRemove[T comparable](prev *Node[T], actual *Node[T], data T) (*Node[T], *Node[T]) {
	if actual == nil {
		return prev, nil
	}

	if actual.Data() == data {
		return prev, actual
	}

	return recursiveRemove(actual, actual.next(), data)
}

func (l *LinkedList[T]) Contains(data T) bool {
	gotNode := recursiveContains[T](l.head, data)

	return gotNode != nil
}

func recursiveContains[T comparable](next *Node[T], data T) *Node[T] {
	if next.Data() == data {
		return next
	}

	if next.hasNext() == false {
		return nil
	}

	return recursiveContains(next.next(), data)
}

func (l *LinkedList[T]) Next() *Node[T] {
	if l.actual != nil {
		l.previous = l.actual
		l.actual = l.actual.next()
		return l.actual
	}

	if l.head != nil {
		l.actual = l.head
		return l.actual
	}

	panic(errors.New("no next element"))
}

func (l *LinkedList[T]) Reset() {
	l.actual = nil
}

func (l *LinkedList[T]) HasNext() bool {
	if l.actual != nil {
		return l.actual.hasNext()
	}

	return l.head != nil
}

func (l *LinkedList[T]) RemoveActual() {
	l.previous.nextNode = l.actual.next()
	l.actual = l.previous
}

type List[T comparable] interface {
	Insert(data T)
	Remove(data T)
	Contains(data T) bool
	RemoveActual()
	Next() *Node[T]
	HasNext() bool
	Reset()
}
