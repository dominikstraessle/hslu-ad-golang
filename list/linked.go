package list

type Node[T any] struct {
	next *Node[T]
	data *T
}

type LinkedList[T any] struct {
	head *Node[T]
}

type List[T any] interface {
	Add(T)
	Remove(T)
	Contains(T)
}
