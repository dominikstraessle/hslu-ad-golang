package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('d')
	b.Insert('l')
	b.Insert('b')
	b.Insert('f')
	b.Insert('j')
	b.Insert('n')
	b.Insert('a')
	b.Insert('c')
	b.Insert('e')
	b.Insert('g')
	b.Insert('i')
	b.Insert('k')
	b.Insert('m')
	b.Insert('o')
	want := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")
}

// remove root with no children
func TestBinaryTree_Remove4(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Remove('f')
	b.Insert('h')
	b.Remove('h')
	assert.Equal(t, []int{}, b.Inorder())
}

// remove root with right child
func TestBinaryTree_Remove5(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('i')
	b.Remove('h')
	assert.Equal(t, []int{'i'}, b.Inorder())
}

// remove root with left child
func TestBinaryTree_Remove6(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('a')
	b.Remove('h')
	assert.Equal(t, []int{'a'}, b.Inorder())
}

// remove root with 2 child nodes
func TestBinaryTree_Remove2(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('d')
	b.Insert('l')
	b.Insert('b')
	b.Insert('f')
	b.Insert('j')
	b.Insert('n')
	b.Insert('a')
	b.Insert('c')
	b.Insert('e')
	b.Insert('g')
	b.Insert('i')
	b.Insert('k')
	b.Insert('m')
	b.Insert('o')
	want := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('b')
	want = []int{'a', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('h')
	want = []int{'a', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")
}

// remove the root with 2 child nodes
func TestBinaryTree_Remove3(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('d')
	b.Insert('l')
	b.Insert('b')
	b.Insert('f')
	b.Insert('j')
	b.Insert('n')
	b.Insert('a')
	b.Insert('c')
	b.Insert('e')
	b.Insert('g')
	b.Insert('i')
	b.Insert('k')
	b.Insert('m')
	b.Insert('o')
	want := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('h')
	want = []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")
}

func TestBinaryTree_Remove(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('d')
	b.Insert('l')
	b.Insert('b')
	b.Insert('f')
	b.Insert('j')
	b.Insert('n')
	b.Insert('a')
	b.Insert('c')
	b.Insert('e')
	b.Insert('g')
	b.Insert('i')
	b.Insert('k')
	b.Insert('m')
	b.Insert('o')
	want := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('o')
	want = []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('n')
	want = []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('c')
	want = []int{'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('b')
	want = []int{'a', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('e')
	want = []int{'a', 'd', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('f')
	want = []int{'a', 'd', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('d')
	want = []int{'a', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")

	b.Remove('h')
	want = []int{'a', 'g', 'i', 'j', 'k', 'l', 'm'}
	assert.Equalf(t, want, b.Inorder(), "Inorder()")
}

func TestBinaryTree_Search(t *testing.T) {
	b := &BinaryTree[int]{}
	b.Insert('h')
	b.Insert('d')
	b.Insert('l')
	b.Insert('b')
	b.Insert('f')
	b.Insert('j')
	assert.True(t, b.Search('h'))
	assert.True(t, b.Search('d'))
	assert.True(t, b.Search('l'))
	assert.True(t, b.Search('b'))
	assert.True(t, b.Search('f'))
	assert.True(t, b.Search('j'))
	assert.False(t, b.Search('e'))
	assert.False(t, b.Search('g'))
	assert.False(t, b.Search('i'))
	assert.False(t, b.Search('k'))
	assert.False(t, b.Search('m'))
	assert.False(t, b.Search('o'))
}

func TestBinaryTree_Inorder(t *testing.T) {
	type fields struct {
		root *Node[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "simple tree",
			fields: fields{
				root: &Node[int]{
					value: 2,
					left: &Node[int]{
						value: 1,
					},
					right: &Node[int]{
						value: 3,
					},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "full",
			fields: fields{
				root: &Node[int]{
					value: 'h',
					left: &Node[int]{
						value: 'd',
						left: &Node[int]{
							value: 'b',
							left: &Node[int]{
								value: 'a',
							},
							right: &Node[int]{
								value: 'c',
							},
						},
						right: &Node[int]{
							value: 'f',
							left: &Node[int]{
								value: 'e',
							},
							right: &Node[int]{
								value: 'g',
							},
						},
					},
					right: &Node[int]{
						value: 'l',
						left: &Node[int]{
							value: 'j',
							left: &Node[int]{
								value: 'i',
							},
							right: &Node[int]{
								value: 'k',
							},
						},
						right: &Node[int]{
							value: 'n',
							left: &Node[int]{
								value: 'm',
							},
							right: &Node[int]{
								value: 'o',
							},
						},
					},
				},
			},
			want: []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree[int]{
				root: tt.fields.root,
			}
			assert.Equalf(t, tt.want, b.Inorder(), "Inorder()")
		})
	}
}
