package graph

import "fmt"

// Edge represents a weighted directional connection between two vertices.
type Edge struct {
	Weight float64
	start  *Vertex
	end    *Vertex
}

// NewEdge creates a new Edge. Vertices are unaffected.
func NewEdge(start, end *Vertex, weight float64) *Edge {
	return &Edge{
		Weight: weight,
		start:  start,
		end:    end,
	}
}

// Reverse reverses the Edge direction.
func (e *Edge) Reverse() {
	e.start, e.end = e.end, e.start
}

// String retrieves a string representation of the Edge.
func (e *Edge) String() string {
	return fmt.Sprintf("%s to %s", e.start.String(), e.end.String())
}
