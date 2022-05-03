package heap

import "testing"

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
}
