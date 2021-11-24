package graph

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

var (
	ErrVertexExists    = errors.New("vertex already exists")
	ErrVertexNotExists = errors.New("vertex does not exist")

	ErrEdgeExists    = errors.New("edge already exists")
	ErrEdgeNotExists = errors.New("edge does not exist")
)

type Graph struct {
	directed bool
	edges    []*Edge
	vertices []*Vertex
}

func NewDirected() *Graph {
	return &Graph{directed: true}
}

func NewUndirected() *Graph {
	return &Graph{directed: false}
}

func (g *Graph) AddVertex(v *Vertex) error {
	if g.vertexExists(v) {
		return fmt.Errorf("%w: %d", ErrVertexExists, v.Value)
	}
	g.vertices = append(g.vertices, v)
	return nil
}

func (g *Graph) AddVertices(v ...*Vertex) error {
	var err error
	for _, v := range v {
		err = g.AddVertex(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) AddEdge(e *Edge) error {
	// Check if edge exists
	for _, edge := range g.edges {
		if edge == e {
			return fmt.Errorf("%w: %s", ErrEdgeExists, e)
		}
	}

	// Ensure vertices exist
	if !g.vertexExists(e.start) {
		err := g.AddVertices(e.start)
		if err != nil {
			return err
		}
	}
	if !g.vertexExists(e.end) {
		err := g.AddVertices(e.end)
		if err != nil {
			return err
		}
	}

	// Add edge
	g.edges = append(g.edges, e)

	err := e.start.AddEdge(e)
	if err != nil {
		return err
	}
	if !g.directed { // Undirected have edge both ways
		err = e.end.AddEdge(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) DeleteEdge(e *Edge) error {
	for i, v := range g.edges {
		if v == e {
			err := e.start.DeleteEdge(e)
			if err != nil {
				return err
			}
			if !g.directed { // Undirected have edge both ways
				err := e.end.DeleteEdge(e)
				if err != nil {
					return err
				}
			}
			g.edges = append(g.edges[:i], g.edges[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%w: %s", ErrEdgeNotExists, e)
}

func (g *Graph) FindEdge(start, end *Vertex) *Edge {
	if !g.vertexExists(start) {
		return nil
	}
	return start.FindEdge(end)
}

func (g *Graph) GetWeight() float64 {
	weight := float64(0)
	for _, edge := range g.edges {
		weight += edge.Weight
	}
	return weight
}

func (g *Graph) Reverse() {
	for _, e := range g.edges {
		_ = e.start.DeleteEdge(e)
		_ = e.end.AddEdge(e)
		e.Reverse()
	}
}

func (g *Graph) GetAdjacencyMatrix() [][]float64 {
	const inf = math.MaxFloat64

	adjacency := make([][]float64, len(g.vertices))
	for i := range adjacency {
		adjacency[i] = make([]float64, len(g.vertices))
		for j := range adjacency[i] {
			adjacency[i][j] = inf
		}
	}

	indices := g.GetVerticesIndices()
	for i, v := range g.vertices {
		for _, neighbor := range v.GetNeighbors() {
			ni := indices[neighbor]
			edge := g.FindEdge(v, neighbor)
			adjacency[i][ni] = edge.Weight
		}
	}
	return adjacency
}

func (g *Graph) String() string {
	keys := make([]int, 0, len(g.vertices))
	for i := range g.vertices {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(g.vertices[k].String())
	}
	return sb.String()
}

func (g *Graph) GetVerticesIndices() map[*Vertex]int {
	indices := make(map[*Vertex]int)
	for i, v := range g.vertices {
		indices[v] = i
	}
	return indices
}

func (g *Graph) GetVertices() []*Vertex {
	return g.vertices
}

func (g *Graph) GetEdges() []*Edge {
	return g.edges
}

func (g *Graph) vertexExists(v *Vertex) bool {
	for _, vertex := range g.vertices {
		if vertex == v {
			return true
		}
	}
	return false
}
