package graph

import (
	"errors"
	"fmt"
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
	vertices map[int]*Vertex
	edges    []*Edge
	directed bool
}

func NewDirected() *Graph {
	return &Graph{
		vertices: make(map[int]*Vertex),
		directed: true,
	}
}

func NewUndirected() *Graph {
	return &Graph{
		vertices: make(map[int]*Vertex),
		directed: false,
	}
}

func (g *Graph) AddVertices(v ...*Vertex) error {
	for _, v := range v {
		if _, ok := g.vertices[v.Value]; ok {
			return fmt.Errorf("%w: %d", ErrVertexExists, v.Value)
		}
		g.vertices[v.Value] = v
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
	if _, ok := g.vertices[e.start.Value]; !ok {
		err := g.AddVertices(e.start)
		if err != nil {
			return err
		}
	}
	if _, ok := g.vertices[e.end.Value]; !ok {
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
	v, ok := g.vertices[start.Value]
	if !ok {
		return nil
	}
	return v.FindEdge(end)
}

func (g *Graph) GetWeight() float64 {
	weight := float64(0)
	for _, edge := range g.edges {
		weight += edge.Weight
	}
	return weight
}

func (g *Graph) Reverse() {
	for _, edge := range g.edges {
		edge.Reverse()
	}
}

func (g *Graph) GetAdjacencyMatrix() map[*Vertex]map[*Vertex]*Edge {
	adjacency := make(map[*Vertex]map[*Vertex]*Edge)
	for _, v := range g.vertices {
		for _, n := range v.GetNeighbors() {
			adjacency[v][n] = g.FindEdge(v, n)
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
