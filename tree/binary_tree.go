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
		b.switchAndRemoveRoot(t)
	} else if t < b.root.value {
		remove(b.root, b.root.left, t)
	} else {
		remove(b.root, b.root.right, t)
	}
}

func (b *BinaryTree[T]) switchAndRemoveRoot(t T) {
	// both child nodes exist -> switch and remove
	if b.root.left != nil && b.root.right != nil {
		b.root.value = getLeftMostValue(b.root, b.root.right, t)
		remove(b.root, b.root.right, t)
		return
	}

	// replace by left child node
	if b.root.left != nil {
		b.root = b.root.left
		return
	}

	// replace by right child node
	b.root = b.root.right
}

func remove[T ad.Ordered](parent *Node[T], n *Node[T], t T) {
	if n == nil {
		return
	}

	// search and remove on the left child node
	if t < n.value {
		remove(n, n.left, t)
		return
	}

	// search and remove on the right child node
	if t > n.value {
		remove(n, n.right, t)
		return
	}

	// the actual node should be replaced by the left child node
	if n.left != nil && n.right == nil {
		switchToLeftChildNode(parent, n)
		return
	}

	// the actual node should be replaced by the right child node
	if n.right != nil && n.left == nil {
		switchToRightChildNode(parent, n)
		return
	}

	// the node has two child nodes and therefore must be replaced by the lowest right-hand side value
	switchAndRemoveNode(parent, n, t)
}

func switchToRightChildNode[T ad.Ordered](parent *Node[T], n *Node[T]) {
	if parent.left == n {
		parent.left = n.right
	} else {
		parent.right = n.right
	}
}

func switchToLeftChildNode[T ad.Ordered](parent *Node[T], n *Node[T]) {
	if parent.left == n {
		parent.left = n.left
	} else {
		parent.right = n.left
	}
}

func switchAndRemoveNode[T ad.Ordered](parent *Node[T], n *Node[T], t T) {
	if n.left != nil && n.right != nil {
		n.value = getLeftMostValue(n, n.right, t)
		remove(n, n.right, t)
		return
	}
	if parent.left == n {
		parent.left = nil
		return
	}

	parent.right = nil
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
