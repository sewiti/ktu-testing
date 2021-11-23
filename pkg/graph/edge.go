package graph

import "fmt"

type Edge struct {
	Weight float64
	start  *Vertex
	end    *Vertex
}

func NewEdge(start, end *Vertex, weight float64) *Edge {
	return &Edge{
		Weight: weight,
		start:  start,
		end:    end,
	}
}

func (e *Edge) Reverse() {
	e.start, e.end = e.end, e.start
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s to %s", e.start.String(), e.end.String())
}
