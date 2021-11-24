package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertex(t *testing.T) {
	t.Run("should create graph vertex", func(t *testing.T) {
		v := NewVertex(5)

		assert.NotNil(t, v)
		assert.Equal(t, 5, v.Value)
		assert.Equal(t, "5", v.String())
		assert.Len(t, v.edges, 0)
	})

	t.Run("should add edges to vertex and check if it exists", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)

		e := NewEdge(v0, v1, 5)
		assert.NoError(t, v0.AddEdges(e))

		assert.True(t, v0.HasEdge(e))
		assert.False(t, v1.HasEdge(e))
		assert.Len(t, v0.GetEdges(), 1)
		assert.Equal(t, "0 to 1", v0.GetEdges()[0].String())
	})

	t.Run("should delete edges from vertex", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e02 := NewEdge(v0, v2, 0)
		assert.NoError(t, v0.AddEdges(e01, e02))

		assert.True(t, v0.HasEdge(e01))
		assert.False(t, v1.HasEdge(e01))

		assert.True(t, v0.HasEdge(e02))
		assert.False(t, v2.HasEdge(e02))

		assert.Len(t, v0.GetEdges(), 2)

		assert.Equal(t, "0 to 1", v0.GetEdges()[0].String())
		assert.Equal(t, "0 to 2", v0.GetEdges()[1].String())

		assert.NoError(t, v0.DeleteEdge(e01))

		assert.False(t, v0.HasEdge(e01))
		assert.True(t, v0.HasEdge(e02))
		assert.Equal(t, "0 to 2", v0.GetEdges()[0].String())

		assert.NoError(t, v0.DeleteEdge(e02))

		assert.False(t, v0.HasEdge(e01))
		assert.False(t, v0.HasEdge(e02))
		assert.Len(t, v0.GetEdges(), 0)
	})

	t.Run("should delete all edges from vertex", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e02 := NewEdge(v0, v2, 0)
		assert.NoError(t, v0.AddEdges(e01, e02))

		assert.True(t, v0.HasEdge(e01))
		assert.False(t, v1.HasEdge(e01))

		assert.True(t, v0.HasEdge(e02))
		assert.False(t, v2.HasEdge(e02))

		assert.Len(t, v0.GetEdges(), 2)

		v0.DeleteAllEdges()

		assert.False(t, v0.HasEdge(e01))
		assert.False(t, v1.HasEdge(e01))

		assert.False(t, v0.HasEdge(e02))
		assert.False(t, v2.HasEdge(e02))

		assert.Len(t, v0.GetEdges(), 0)
	})

	t.Run("should return vertex neighbors in case if current node is start one", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e02 := NewEdge(v0, v2, 0)
		assert.NoError(t, v0.AddEdges(e01, e02))

		assert.Len(t, v1.GetNeighbors(), 0)

		neighbors := v0.GetNeighbors()

		assert.Len(t, neighbors, 2)
		assert.Equal(t, v1, neighbors[0])
		assert.Equal(t, v2, neighbors[1])
	})

	t.Run("should return vertex neighbors in case if current node is end one", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e10 := NewEdge(v1, v0, 0)
		e20 := NewEdge(v2, v0, 0)
		assert.NoError(t, v0.AddEdges(e10, e20))

		assert.Len(t, v1.GetNeighbors(), 0)

		neighbors := v0.GetNeighbors()

		assert.Len(t, neighbors, 2)
		assert.Equal(t, v1, neighbors[0])
		assert.Equal(t, v2, neighbors[1])
	})

	t.Run("should check if vertex has specific neighbor", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		assert.NoError(t, v0.AddEdges(e01))

		assert.True(t, v0.HasNeighbor(v1))
		assert.False(t, v0.HasNeighbor(v2))
	})

	t.Run("should edge by vertex", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		assert.NoError(t, v0.AddEdges(e01))

		assert.Equal(t, e01, v0.FindEdge(v1))
		assert.Nil(t, v0.FindEdge(v2))
	})

	t.Run("should calculate vertex degree", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		assert.Equal(t, 0, v0.GetDegree())

		e01 := NewEdge(v0, v1, 0)
		assert.NoError(t, v0.AddEdges(e01))
		assert.Equal(t, 1, v0.GetDegree())

		e10 := NewEdge(v1, v0, 0)
		assert.NoError(t, v0.AddEdges(e10))
		assert.Equal(t, 2, v0.GetDegree())

		e20 := NewEdge(v2, v0, 0)
		assert.NoError(t, v0.AddEdges(e20))
		assert.Equal(t, 3, v0.GetDegree())

		assert.Len(t, v0.GetEdges(), 3)
	})
}
