package graph

import "fmt"

func ExampleNewEdge() {
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	edge := NewEdge(v0, v1, 5)

	fmt.Printf("%s\n", edge)
	// Output:
	// 0 to 1
}

func ExampleEdge_Reverse() {
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	edge := NewEdge(v0, v1, 5)

	fmt.Printf("%s\n", edge)
	edge.Reverse()
	fmt.Printf("%s\n", edge)
	// Output:
	// 0 to 1
	// 1 to 0
}

func ExampleEdge_String() {
	v0 := NewVertex(0)
	v1 := NewVertex(1)
	edge := NewEdge(v0, v1, 5)

	fmt.Printf("%s\n", edge.String())
	// Output:
	// 0 to 1
}
