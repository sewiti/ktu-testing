package graph

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

var (
	ErrExists    = errors.New("already exists")
	ErrNotExists = errors.New("does not exist")
)

// Graph represents a set of vertices and connections between them.
type Graph struct {
	directed bool
	edges    []*Edge
	vertices []*Vertex
}

// NewDirected creates a new directed Graph.
func NewDirected() *Graph {
	return &Graph{directed: true}
}

// NewUndirected creates a new undirected Graph.
func NewUndirected() *Graph {
	return &Graph{directed: false}
}

// AddVertices adds vertices to Graph.
//
// Returns first ErrExists if Vertex is a duplicate.
func (g *Graph) AddVertices(v ...*Vertex) error {
	for _, v := range v {
		if err := g.addVertex(v); err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) addVertex(v *Vertex) error {
	if g.vertexExists(v) {
		return fmt.Errorf("vertex %w: %d", ErrExists, v.Value)
	}
	g.vertices = append(g.vertices, v)
	return nil
}

// AddEdges adds edges to Graph, also attaching itself to vertices.
//
// If unknown vertices are encountered, they are also added.
//
// Returns first ErrExists if Edge is a duplicate.
func (g *Graph) AddEdges(e ...*Edge) error {
	for _, e := range e {
		if err := g.addEdge(e); err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) addEdge(e *Edge) error {
	// Check if edge exists
	for _, edge := range g.edges {
		if edge == e {
			return fmt.Errorf("edge %w: %s", ErrExists, e)
		}
	}

	// Ensure vertices exist
	if !g.vertexExists(e.start) {
		err := g.addVertex(e.start)
		if err != nil {
			return err
		}
	}
	if !g.vertexExists(e.end) {
		err := g.addVertex(e.end)
		if err != nil {
			return err
		}
	}

	// Add edge
	g.edges = append(g.edges, e)

	err := e.start.addEdge(e)
	if err != nil {
		return err
	}
	if !g.directed { // Undirected have edge both ways
		err = e.end.addEdge(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteEdge deletes an edge, including deleting it from associated vertices.
//
// Returns ErrNotExists if edge doesn't exist.
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
	return fmt.Errorf("edge %w: %s", ErrNotExists, e)
}

// FindEdge retrieves an edge which connects the vertices given.
//
// Returns nil if doesn't exist.
func (g *Graph) FindEdge(start, end *Vertex) *Edge {
	if !g.vertexExists(start) {
		return nil
	}
	return start.FindEdge(end)
}

// GetWeight retrieves a sum of all edges weights.
func (g *Graph) GetWeight() float64 {
	weight := float64(0)
	for _, edge := range g.edges {
		weight += edge.Weight
	}
	return weight
}

// Reverse reverses all Graph's edges.
func (g *Graph) Reverse() {
	for _, e := range g.edges {
		_ = e.start.DeleteEdge(e)
		_ = e.end.addEdge(e)
		e.Reverse()
	}
}

// GetAdjacencyMatrix retrieves an adjacency matrix between each Vertex indices.
// The value is edge weight.
//
// If vertices aren't connected, value is set to math.MaxFloat64.
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

// String retrieves a string representation of the Graph: vertices values
// separated by space.
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

// GetVerticesIndices retrieves a map of the vertices indices.
//
// Key is *Vertex, value is it's index in the Graph.
func (g *Graph) GetVerticesIndices() map[*Vertex]int {
	indices := make(map[*Vertex]int)
	for i, v := range g.vertices {
		indices[v] = i
	}
	return indices
}

// GetVertices retrieves all Graph's vertices.
func (g *Graph) GetVertices() []*Vertex {
	return g.vertices
}

// GetEdges retrieves all Graph's edges.
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
