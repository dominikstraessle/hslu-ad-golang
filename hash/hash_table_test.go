package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable_Put(t *testing.T) {
	table := New[string]()
	table.Put("a")
	table.Put("b")
	table.Put("c")
	table.Put("d")
	table.Put("e")
	table.Put("e")

	want := []string{
		"a", "b", "c", "d", "e",
	}
	got := table.GetAll()
	assert.Equal(t, len(want), len(got))
	for _, s := range want {
		assert.Contains(t, got, s)
	}
	got = table.GetAll()
	assert.Equal(t, len(want), len(got))
	for _, s := range want {
		assert.Contains(t, got, s)
	}
	got = table.GetAll()
	assert.Equal(t, len(want), len(got))
	for _, s := range want {
		assert.Contains(t, got, s)
	}
}

func TestTable_Contains(t *testing.T) {
	table := New[string]()
	table.Put("a")
	table.Put("b")
	table.Put("c")
	table.Put("d")
	table.Put("e")

	want := []string{
		"a", "b", "c", "d", "e",
	}
	for _, s := range want {
		assert.Truef(t, table.Contains(s), "table %v should contain %s", table.GetAll(), s)
	}

	wantNot := []string{
		"x", "123", "blub",
	}
	for _, s := range wantNot {
		assert.Falsef(t, table.Contains(s), "table %v should not contain %s", table.GetAll(), s)
	}
}

func TestTable_Remove(t *testing.T) {
	table := New[string]()
	table.Put("a")
	table.Put("b")
	table.Put("c")
	table.Put("d")
	table.Put("e")

	want := []string{
		"a", "b", "c", "d", "e",
	}
	for _, s := range want {
		assert.Truef(t, table.Contains(s), "table %v should contain %s", table.GetAll(), s)
		table.Remove(s)
		assert.Falsef(t, table.Contains(s), "table %v should not contain %s", table.GetAll(), s)
	}
}

func TestTable_Collisions(t *testing.T) {
	table := New[string]()
	table.Put("a")
	table.Put("b")
	table.Put("c")
	table.Put("d")
	table.Put("e")
	table.Put("f")
	table.Put("g")
	table.Put("h")
	table.Put("i")
	table.Put("j")
	table.Put("a")
	table.Put("b")
	table.Put("c")
	table.Put("d")
	table.Put("e")
	table.Put("f")
	table.Put("g")
	table.Put("h")
	table.Put("i")
	table.Put("j")

	want := []string{
		"a", "b", "c", "d", "e",
	}
	for _, s := range want {
		assert.Truef(t, table.Contains(s), "table %v should contain %s", table.GetAll(), s)
		table.Remove(s)
		assert.Falsef(t, table.Contains(s), "table %v should not contain %s", table.GetAll(), s)
	}
}
