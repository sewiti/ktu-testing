package main

import (
	"log"

	"github.com/sewiti/ktu-testing/pkg/graph"
)

func main() {
	g := graph.NewDirected()

	v := []*graph.Vertex{
		graph.NewVertex(0),
		graph.NewVertex(1),
		graph.NewVertex(2),
	}

	err := g.AddVertices(v...)
	if err != nil {
		log.Fatalln(err)
	}

	edge := graph.NewEdge(v[0], v[1], 0)
	err = g.AddEdges(edge)
	if err != nil {
		log.Fatalln(err)
	}
}
