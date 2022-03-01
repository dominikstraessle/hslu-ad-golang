package list

import "errors"

type Data string

type Node struct {
	nextNode *Node
	data     Data
}

func (n *Node) next() *Node {
	return n.nextNode
}

func (n *Node) hasNext() bool {
	return n.nextNode != nil
}

func (n *Node) Data() Data {
	return n.data
}

type Linker interface {
	Data() Data
}

type LinkedList struct {
	head   *Node
	actual *Node
}

func NewLinkedList(data Data) List {
	return &LinkedList{
		head: &Node{
			nextNode: nil,
			data:     data,
		},
		actual: nil,
	}
}

func (l *LinkedList) Add(data Data) {
	node := &Node{
		nextNode: l.head.next(),
		data:     data,
	}

	l.head.nextNode = node
}

func (l *LinkedList) Remove(data Data) {
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

func recursiveRemove(prev *Node, actual *Node, data Data) (*Node, *Node) {
	if actual == nil {
		return prev, nil
	}

	if actual.Data() == data {
		return prev, actual
	}

	return recursiveRemove(actual, actual.next(), data)
}

func (l *LinkedList) Contains(data Data) bool {
	gotNode := recursiveContains(l.head, data)

	return gotNode != nil
}

func recursiveContains(next *Node, data Data) *Node {
	if next.Data() == data {
		return next
	}

	if next.hasNext() == false {
		return nil
	}

	return recursiveContains(next.next(), data)
}

func (l *LinkedList) Next() *Node {
	if l.actual != nil {
		result := l.actual
		l.actual = l.actual.next()
		return result
	}

	if l.head != nil {
		l.actual = l.head
		return l.head
	}

	panic(errors.New("no next element"))
}

func (l *LinkedList) Reset() {
	l.actual = nil
}

func (l *LinkedList) HasNext() bool {
	return l.actual != nil || l.head != nil
}

type List interface {
	Add(data Data)
	Remove(data Data)
	Contains(data Data) bool
	Next() *Node
	HasNext() bool
	Reset()
}
