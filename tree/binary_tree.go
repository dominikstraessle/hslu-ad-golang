//Package tree -> Week 02
package tree

import (
	"ad"
)

type Tree[T ad.Ordered] interface {
	Insert(T)
	Remove(T)
	Search(T) bool
	Inorder() []T
}

type BinaryTree[T ad.Ordered] struct {
	root *Node[T]
}

func (b *BinaryTree[T]) Insert(t T) {
	if b.root == nil {
		b.root = &Node[T]{
			value: t,
		}
		return
	}
	insert(b.root, t)
}

func insert[T ad.Ordered](n *Node[T], t T) {
	if t < n.value {
		if n.left == nil {
			n.left = &Node[T]{
				value: t,
			}
		} else {
			insert(n.left, t)
		}
	} else {
		if n.right == nil {
			n.right = &Node[T]{
				value: t,
			}
		} else {
			insert(n.right, t)
		}
	}
}

func (b *BinaryTree[T]) Remove(t T) {
	if b.root == nil {
		return
	}

	if b.root.value == t && b.root.left == nil && b.root.right == nil {
		b.root = nil
		return
	}
	if b.root.value == t {
		if b.root.left != nil && b.root.right != nil {
			b.root.value = getLeftMostValue(b.root, b.root.right, t)
			remove(b.root, b.root.right, t)
		} else if b.root.left != nil {
			b.root = b.root.left
		} else {
			b.root = b.root.right
		}
	} else if t < b.root.value {
		remove(b.root, b.root.left, t)
	} else {
		remove(b.root, b.root.right, t)
	}
}

func remove[T ad.Ordered](parent *Node[T], n *Node[T], t T) {
	if n == nil {
		return
	}

	if n.value == t {
		if n.left != nil && n.right == nil {
			if parent.left == n {
				parent.left = n.left
			} else {
				parent.right = n.left
			}
		} else if n.right != nil && n.left == nil {
			if parent.left == n {
				parent.left = n.right
			} else {
				parent.right = n.right
			}
		} else if n.left != nil && n.right != nil {
			n.value = getLeftMostValue(n, n.right, t)
			remove(n, n.right, t)
		} else {
			if parent.left == n {
				parent.left = nil
			} else {
				parent.right = nil
			}
		}
	} else if t < n.value {
		remove(n, n.left, t)
	} else {
		remove(n, n.right, t)
	}
}

func getLeftMostValue[T ad.Ordered](parent *Node[T], left *Node[T], t T) T {
	if left.left == nil {
		tmp := left.value
		left.value = t
		return tmp
	}

	return getLeftMostValue(parent, left.left, t)
}

func (b *BinaryTree[T]) Search(t T) bool {
	return search(b.root, t)
}

func search[T ad.Ordered](n *Node[T], t T) bool {
	if n == nil {
		return false
	}
	if t == n.value {
		return true
	}
	if t < n.value {
		return search(n.left, t)
	}
	return search(n.right, t)
}

func (b *BinaryTree[T]) Inorder() []T {
	return inorder([]T{}, b.root)
}

func inorder[T ad.Ordered](values []T, n *Node[T]) []T {
	if n == nil {
		return values
	}
	if n.left != nil {
		values = inorder(values, n.left)
	}
	values = append(values, n.value)
	if n.right != nil {
		values = inorder(values, n.right)
	}
	return values
}

type Node[T ad.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}
