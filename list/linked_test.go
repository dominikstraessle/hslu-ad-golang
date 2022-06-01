package list

import (
	"reflect"
	"testing"
)

func TestLinkedList_HasNext(t *testing.T) {
	next := &Node{}
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
				head: &Node{
					nextNode: next,
				},
				actual: next,
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
				actual: head,
			},
			want: head.next(),
		},
		{
			name: "return third",
			fields: fields{
				head:   head,
				actual: head.next(),
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
	list.Insert("other")
	list.Insert("other2")
	list.Insert(want)
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
	list.Insert(second)
	list.Insert("other data")
	list.Insert("other data")
	list.Insert("other data")
	list.Insert(third)

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

func TestIterator(t *testing.T) {
	list := NewLinkedList("6")
	list.Insert("5")
	list.Insert("4")
	list.Insert("3")
	list.Insert("2")
	list.Insert("1")

	want := []Data{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	var got []Data

	for list.HasNext() == true {
		got = append(got, list.Next().Data())
	}

	for i, data := range want {
		if got[i] != data {
			t.Errorf("got %v but want %v", got, want)
		}
	}
}
