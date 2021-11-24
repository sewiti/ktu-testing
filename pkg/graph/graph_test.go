package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	t.Run("should add vertices to graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)

		assert.NoError(t, g.AddVertices(v0, v1))
		assert.Equal(t, "0 1", g.String())
	})

	t.Run("should add edges to undirected graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)

		e01 := NewEdge(v0, v1, 0)

		assert.NoError(t, g.AddEdges(e01))

		assert.Len(t, g.GetVertices(), 2)
		assert.Equal(t, v0, g.GetVertices()[0])
		assert.Equal(t, v1, g.GetVertices()[1])

		assert.Equal(t, "0 1", g.String())

		assert.Equal(t, 1, v0.GetDegree())
		assert.Equal(t, v1, v0.GetNeighbors()[0])
		assert.Equal(t, v1, v0.GetNeighbors()[0])

		assert.Equal(t, 1, v1.GetDegree())
		assert.Equal(t, v0, v1.GetNeighbors()[0])
		assert.Equal(t, v0, v1.GetNeighbors()[0])
	})

	t.Run("should add edges to directed graph", func(t *testing.T) {
		g := NewDirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)

		e01 := NewEdge(v0, v1, 0)

		assert.NoError(t, g.AddEdges(e01))

		assert.Equal(t, "0 1", g.String())

		assert.Equal(t, 1, v0.GetDegree())
		assert.Equal(t, v1, v0.GetNeighbors()[0])
		assert.Equal(t, v1, v0.GetNeighbors()[0])

		assert.Equal(t, 0, v1.GetDegree())
	})

	t.Run("should find edge by vertices in undirected graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 10)

		assert.NoError(t, g.AddEdges(e01))

		ge01 := g.FindEdge(v0, v1)
		ge10 := g.FindEdge(v1, v0)
		ge02 := g.FindEdge(v0, v2)
		ge20 := g.FindEdge(v2, v0)

		assert.Nil(t, ge02)
		assert.Nil(t, ge20)
		assert.Equal(t, e01, ge01)
		assert.Equal(t, e01, ge10)
		assert.Equal(t, float64(10), ge01.Weight)
	})

	t.Run("should find edge by vertices in directed graph", func(t *testing.T) {
		g := NewDirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 10)

		assert.NoError(t, g.AddEdges(e01))

		ge01 := g.FindEdge(v0, v1)
		ge10 := g.FindEdge(v1, v0)
		ge02 := g.FindEdge(v0, v2)
		ge20 := g.FindEdge(v2, v0)

		assert.Nil(t, ge02)
		assert.Nil(t, ge20)
		assert.Nil(t, ge10)
		assert.Equal(t, e01, ge01)
		assert.Equal(t, float64(10), ge01.Weight)
	})

	t.Run("should return vertex neighbors", func(t *testing.T) {
		g := NewDirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e02 := NewEdge(v0, v2, 0)

		assert.NoError(t, g.AddEdges(e01, e02))

		neighbors := v0.GetNeighbors()

		assert.Len(t, neighbors, 2)
		assert.Equal(t, v1, neighbors[0])
		assert.Equal(t, v2, neighbors[1])
	})

	t.Run("should throw an error when trying to add edge twice", func(t *testing.T) {
		g := NewDirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)

		e01 := NewEdge(v0, v1, 0)

		assert.NoError(t, g.AddEdges(e01))
		assert.Error(t, g.AddEdges(e01))
	})

	t.Run("should return the list of all added edges", func(t *testing.T) {
		g := NewDirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)

		assert.NoError(t, g.AddEdges(e01, e12))

		edges := g.GetEdges()

		assert.Len(t, edges, 2)
		assert.Equal(t, e01, edges[0])
		assert.Equal(t, e12, edges[1])
	})

	t.Run("should calculate total graph weight for default graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)
		e23 := NewEdge(v2, v3, 0)
		e03 := NewEdge(v0, v3, 0)

		assert.NoError(t, g.AddEdges(e01, e12, e23, e03))
		assert.Equal(t, float64(0), g.GetWeight())
	})

	t.Run("should calculate total graph weight for weighted graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 1)
		e12 := NewEdge(v1, v2, 2)
		e23 := NewEdge(v2, v3, 3)
		e03 := NewEdge(v0, v3, 4)

		assert.NoError(t, g.AddEdges(e01, e12, e23, e03))
		assert.Equal(t, float64(10), g.GetWeight())
	})

	t.Run("should be possible to delete edges from graph", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)
		e02 := NewEdge(v0, v2, 0)

		assert.NoError(t, g.AddEdges(e01, e12, e02))
		assert.Len(t, g.GetEdges(), 3)

		assert.NoError(t, g.DeleteEdge(e01))

		assert.Len(t, g.GetEdges(), 2)
		assert.Equal(t, e12.String(), g.GetEdges()[0].String())
		assert.Equal(t, e02.String(), g.GetEdges()[1].String())
	})

	t.Run("should should throw an error when trying to delete not existing edge", func(t *testing.T) {
		g := NewUndirected()

		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)

		assert.NoError(t, g.AddEdges(e01))
		assert.Error(t, g.DeleteEdge(e12))

	})

	t.Run("should be possible to Reverse graph", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 0)
		e02 := NewEdge(v0, v2, 0)
		e23 := NewEdge(v2, v3, 0)

		g := NewDirected()
		assert.NoError(t, g.AddEdges(e01, e02, e23))

		assert.Equal(t, "0 1 2 3", g.String())
		assert.Len(t, g.GetEdges(), 3)
		assert.Equal(t, 2, v0.GetDegree())
		assert.Equal(t, 0, v1.GetDegree())
		assert.Equal(t, 1, v2.GetDegree())
		assert.Equal(t, 0, v3.GetDegree())
		assert.Equal(t, v1, v0.GetNeighbors()[0])
		assert.Equal(t, v2, v0.GetNeighbors()[1])
		assert.Equal(t, v3, v2.GetNeighbors()[0])

		g.Reverse()

		assert.Equal(t, "0 1 2 3", g.String())
		assert.Len(t, g.GetEdges(), 3)
		assert.Equal(t, 0, v0.GetDegree())
		assert.Equal(t, 1, v1.GetDegree())
		assert.Equal(t, 1, v2.GetDegree())
		assert.Equal(t, 1, v3.GetDegree())
		assert.Equal(t, v0, v1.GetNeighbors()[0])
		assert.Equal(t, v0, v2.GetNeighbors()[0])
		assert.Equal(t, v2, v3.GetNeighbors()[0])
	})

	t.Run("should return vertices indices", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)
		e23 := NewEdge(v2, v3, 0)
		e13 := NewEdge(v1, v3, 0)

		g := NewUndirected()
		assert.NoError(t, g.AddEdges(e01, e12, e23, e13))

		indices := g.GetVerticesIndices()
		expected := map[*Vertex]int{
			v0: 0,
			v1: 1,
			v2: 2,
			v3: 3,
		}
		assert.Equal(t, expected, indices)
	})

	t.Run("should generate adjacency matrix for undirected graph", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 0)
		e12 := NewEdge(v1, v2, 0)
		e23 := NewEdge(v2, v3, 0)
		e13 := NewEdge(v1, v3, 0)

		g := NewUndirected()
		assert.NoError(t, g.AddEdges(e01, e12, e23, e13))

		const inf = math.MaxFloat64
		adjacency := g.GetAdjacencyMatrix()
		expected := [][]float64{
			{inf, 0, inf, inf},
			{0, inf, 0, 0},
			{inf, 0, inf, 0},
			{inf, 0, 0, inf},
		}
		assert.Equal(t, expected, adjacency)
	})

	t.Run("should generate adjacency matrix for directed graph", func(t *testing.T) {
		v0 := NewVertex(0)
		v1 := NewVertex(1)
		v2 := NewVertex(2)
		v3 := NewVertex(3)

		e01 := NewEdge(v0, v1, 2)
		e12 := NewEdge(v1, v2, 1)
		e23 := NewEdge(v2, v3, 5)
		e13 := NewEdge(v1, v3, 7)

		g := NewDirected()
		assert.NoError(t, g.AddEdges(e01, e12, e23, e13))

		const inf = math.MaxFloat64
		adjacency := g.GetAdjacencyMatrix()
		expected := [][]float64{
			{inf, 2, inf, inf},
			{inf, inf, 1, 7},
			{inf, inf, inf, 5},
			{inf, inf, inf, inf},
		}
		assert.Equal(t, expected, adjacency)
	})
}
