package graph

import (
	"fmt"
	"strconv"
)

// Vertex represents a single point of the Graph.
type Vertex struct {
	Value int // Unique Vertex value.
	edges []*Edge
}

// NewVertex creates a new Vertex with a value.
func NewVertex(value int) *Vertex {
	return &Vertex{Value: value}
}

// AddEdges adds edges to Vertex.
//
// Returns first ErrExists if Edge is a duplicate.
func (v *Vertex) AddEdges(e ...*Edge) error {
	for _, e := range e {
		if err := v.addEdge(e); err != nil {
			return err
		}
	}
	return nil
}

func (v *Vertex) addEdge(e *Edge) error {
	for _, edge := range v.edges {
		if edge == e {
			return fmt.Errorf("edge %w: %s", ErrExists, e)
		}
	}
	v.edges = append(v.edges, e)
	return nil
}

// DeleteEdge deletes an edge from a Vertex.
//
// Returns ErrNotExists if edge doesn't exist.
func (v *Vertex) DeleteEdge(e *Edge) error {
	for i, w := range v.edges {
		if w == e {
			v.edges = append(v.edges[:i], v.edges[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("edge %w: %s", ErrNotExists, e)
}

// GetNeighbors retrieves Vertex neighbors with whom the Vertex is connected.
func (v *Vertex) GetNeighbors() []*Vertex {
	var vertices []*Vertex
	for _, e := range v.edges {
		neighbor := e.start
		if neighbor == v {
			neighbor = e.end
		}
		vertices = append(vertices, neighbor)
	}
	return vertices
}

// GetEdges retrieves all Vertex edges.
func (v *Vertex) GetEdges() []*Edge {
	return v.edges
}

// GetDegree retrieves the number of outgoing edges Vertex has.
func (v *Vertex) GetDegree() int {
	return len(v.edges)
}

// HasEdge reports whetver Vertex has an edge given.
func (v *Vertex) HasEdge(e *Edge) bool {
	for _, edge := range v.edges {
		if edge == e {
			return true
		}
	}
	return false
}

// HasNeighbor reports whetver given Vertex is it's neighbor.
func (v *Vertex) HasNeighbor(w *Vertex) bool {
	for _, e := range v.edges {
		if e.start == w || e.end == w {
			return true
		}
	}
	return false
}

// FindEdge retrieves an edge which connects with Vertex given.
func (v *Vertex) FindEdge(w *Vertex) *Edge {
	for _, e := range v.edges {
		if e.start == w || e.end == w {
			return e
		}
	}
	return nil
}

// DeleteAllEdges deletes all Vertex edges.
func (v *Vertex) DeleteAllEdges() {
	v.edges = nil
}

// String retrieves a string representation of the Vertex.
func (v *Vertex) String() string {
	return strconv.Itoa(v.Value)
}
