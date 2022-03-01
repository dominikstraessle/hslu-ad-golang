package list

import (
	"reflect"
	"testing"
)

func TestLinkedList_HasNext(t *testing.T) {
	type fields struct {
		head   *Node
		actual *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "has next",
			fields: fields{
				head:   &Node{},
				actual: nil,
			},
			want: true,
		},
		{
			name: "has no next",
			fields: fields{
				head:   nil,
				actual: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:   tt.fields.head,
				actual: tt.fields.actual,
			}
			if got := l.HasNext(); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Next(t *testing.T) {
	head := &Node{
		data: "head",
		nextNode: &Node{
			data: "second",
			nextNode: &Node{
				data: "third",
			},
		},
	}

	type fields struct {
		head   *Node
		actual *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name: "return second",
			fields: fields{
				head:   head,
				actual: head.next(),
			},
			want: head.next(),
		},
		{
			name: "return third",
			fields: fields{
				head:   head,
				actual: head.next().next(),
			},
			want: head.next().next(),
		},
		{
			name: "return head",
			fields: fields{
				head:   head,
				actual: nil,
			},
			want: head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:   tt.fields.head,
				actual: tt.fields.actual,
			}
			if got := l.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList("head")
	want := Data("second")
	list.Add("other")
	list.Add("other2")
	list.Add(want)
	if !list.Contains(want) {
		t.Logf("%v should contain element %v", list, want)
	}

	invalidData := Data("invalid")
	if list.Contains(invalidData) {
		t.Logf("%v should not contain element %v", list, invalidData)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	head := Data("head")
	second := Data("second")
	third := Data("third")
	list := NewLinkedList(head)
	list.Add(second)
	list.Add("other data")
	list.Add("other data")
	list.Add("other data")
	list.Add(third)

	list.Remove(third)
	if list.Contains(third) {
		t.Logf("%v should not contain element %v", list, third)
	}

	list.Remove(head)
	if list.Contains(head) {
		t.Logf("%v should not contain element %v", list, head)
	}

	if !list.Contains(second) {
		t.Logf("%v should contain element %v", list, second)
	}
}

func TestLinkedList_Reset(t *testing.T) {
	type fields struct {
		head   *Node
		actual *Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "reset",
			fields: fields{
				head:   &Node{},
				actual: &Node{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:   tt.fields.head,
				actual: tt.fields.actual,
			}
			l.Reset()
			if l.actual != nil {
				t.Logf("iterator should be nil after reset: %v", l)
			}
		})
	}
}
