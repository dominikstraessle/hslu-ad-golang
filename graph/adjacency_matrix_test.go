package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraphA(t *testing.T) {
	nodeName := []string{
		"a", "b", "c", "d", "e",
	}

	adjacencyMatrix := [][]bool{
		{false, true, true, true, false},
		{true, false, true, false, false},
		{false, false, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	g := NewGraphA(nodeName, adjacencyMatrix)
	assert.Equal(t, 5, g.NumberOfNodes())
}

func TestNewRailwayNet1(t *testing.T) {
	nodeName := []string{
		"a", "b", "c", "d", "e",
	}

	adjacencyMatrix := [][]bool{
		{false, true, true, true, false},
		{true, false, true, false, false},
		{false, false, false, true, false},
		{false, false, false, false, false},
		{false, false, false, true, false},
	}

	g := NewRailwayNet1(nodeName, adjacencyMatrix)
	assert.Equal(t, 5, g.NumberOfNodes())

	assert.False(t, g.HasDirectPath("e", "a"))
	g.AddPath("e", "a")
	assert.True(t, g.HasDirectPath("e", "a"))
	g.RemovePath("e", "a")
	assert.False(t, g.HasDirectPath("e", "a"))

	g.AddStation("x")
	assert.False(t, g.HasDirectPath("x", "a"))
	g.AddPath("x", "a")
	assert.True(t, g.HasDirectPath("x", "a"))
}
