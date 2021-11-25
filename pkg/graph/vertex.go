package graph

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	Value int
	edges []*Edge
}

func NewVertex(value int) *Vertex {
	return &Vertex{Value: value}
}

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

func (v *Vertex) DeleteEdge(e *Edge) error {
	for i, w := range v.edges {
		if w == e {
			v.edges = append(v.edges[:i], v.edges[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("edge %w: %s", ErrNotExists, e)
}

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

func (v *Vertex) GetEdges() []*Edge {
	return v.edges
}

func (v *Vertex) GetDegree() int {
	return len(v.edges)
}

func (v *Vertex) HasEdge(e *Edge) bool {
	for _, edge := range v.edges {
		if edge == e {
			return true
		}
	}
	return false
}

func (v *Vertex) HasNeighbor(w *Vertex) bool {
	for _, e := range v.edges {
		if e.start == w || e.end == w {
			return true
		}
	}
	return false
}

func (v *Vertex) FindEdge(w *Vertex) *Edge {
	for _, e := range v.edges {
		if e.start == w || e.end == w {
			return e
		}
	}
	return nil
}

func (v *Vertex) DeleteAllEdges() {
	v.edges = nil
}

func (v *Vertex) String() string {
	return strconv.Itoa(v.Value)
}
