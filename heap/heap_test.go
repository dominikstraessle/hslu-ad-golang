package heap

import (
	"fmt"
	"testing"
)

func TestFixedSizeHeap_Insert(t *testing.T) {
	f := NewFixedSizeHeap(6)
	f.Insert(20)
	f.Insert(10)
	f.Insert(5)
	f.Insert(12)
	f.Insert(7)
	f.Insert(50)
	want := string([]byte{50, 12, 20, 10, 7, 5})
	if f.Print() != want {
		t.Errorf("want %s got %s", want, f.Print())
	}

	f.DeleteRoot()
	want = string([]byte{20, 12, 5, 10, 7, 0})
	if f.Print() != want {
		fmt.Println(f.Print())
		t.Errorf("want %s got %s", want, f.Print())
	}

	f.DeleteRoot()
	want = string([]byte{12, 10, 5, 7, 0, 0})
	if f.Print() != want {
		fmt.Println(f.Print())
		t.Errorf("want %s got %s", want, f.Print())
	}

	f.DeleteRoot()
	want = string([]byte{10, 7, 5, 0, 0, 0})
	if f.Print() != want {
		fmt.Println(f.Print())
		t.Errorf("want %s got %s", want, f.Print())
	}
}

func TestFixedSizeHeap_InsertFull(t *testing.T) {
	t.Skip()
	f := NewFixedSizeHeap(1)
	f.Insert(1)
	f.Insert(1)
}

func TestFixedSizeHeap_InsertSameSize(t *testing.T) {
	f := NewFixedSizeHeap(10)
	f.Insert('a')
	f.Insert('a')
	f.Insert('b')
	f.Insert('a')
	f.Insert('a')
	f.Insert('x')
	f.Insert('a')
	f.Insert('c')
	f.Insert('a')
	f.Insert('x')
	fmt.Println(f.Print())
	if f.Print() != "xxbacaaaaa" {
		t.Errorf("Expected %s got %s", "xxbacaaaaa", f.Print())
	}
}

func TestFixedSizeHeap_DeleteEmpty(t *testing.T) {
	f := NewFixedSizeHeap(10)
	root := f.DeleteRoot()
	if root != 0 {
		t.Errorf("Expected 0 but got %v", root)
	}
}
