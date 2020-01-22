package pkg

import (
	"fmt"
	"strings"
)

// A Graph consists of a set of nodes, each of which has an ID and a name, and a set of edges. An
// edge is simply a connection between two nodes.
type Graph struct {
	nodes    []string
	matrix   [][]bool
	capacity int // number of rows / columns of matrix
}

// NumNodes returns the number of nodes in the graph.
func (g *Graph) NumNodes() int {
	return len(g.nodes)
}

// Neighbors returns the list of neighbors of a given node.
func (g *Graph) Neighbors(i int) []int {
	if len(g.matrix) <= i {
		return nil
	}

	var neighbors []int
	for j, value := range g.matrix[i] {
		if value {
			neighbors = append(neighbors, j)
		}
	}
	return neighbors
}

// PathExists checks if the given path exists in the graph. The path is specified by a list of the
// node IDs in the graph.
func (g *Graph) PathExists(path []int) bool {
	for i := range path {
		if i == 0 {
			continue
		}
		a := path[i-1]
		b := path[i]
		if !g.HasEdge(a, b) {
			return false
		}
	}
	return true
}

// NodeName returns the name for a node ID.
func (g *Graph) NodeName(n int) string {
	if g.ValidNode(n) {
		return g.nodes[n]
	}
	return fmt.Sprintf("<%d ?>", n)
}

// PrintPath prints the names of the nodes in the path, separated by sep.
func (g *Graph) PrintPath(path []int, sep string) string {
	names := make([]string, len(path))
	for i, n := range path {
		names[i] = g.NodeName(n)
	}
	return strings.Join(names, sep)
}

// AddNode adds a node to the graph and returns the node ID.
func (g *Graph) AddNode(name string) int {
	i, ok := g.LookupNode(name)
	if ok {
		return i
	}
	i = len(g.nodes)
	g.nodes = append(g.nodes, name)
	if g.capacity < len(g.nodes) {
		g.resizeMatrix()
	}
	return i
}

func (g *Graph) resizeMatrix() {
	oldCapacity := g.capacity
	g.capacity = 2 * g.NumNodes() // 2 * needed so we don't have to constantly resize
	oldMatrix := g.matrix
	g.matrix = make([][]bool, g.capacity)
	for i := range g.matrix {
		g.matrix[i] = make([]bool, g.capacity)
		if i < oldCapacity {
			copy(g.matrix[i], oldMatrix[i])
		}
	}
}

// LookupNode looks up a node by name. If the node can't be found, it returns ok == false.
func (g *Graph) LookupNode(name string) (id int, ok bool) {
	for i, n := range g.nodes {
		if n == name {
			return i, true
		}
	}
	return 0, false
}

// AddEdge adds an edge between the given nodes.
func (g *Graph) AddEdge(i, j int) {
	if !g.ValidNode(i) || !g.ValidNode(j) {
		return
	}

	if g.HasEdge(i, j) {
		return
	}
	g.matrix[i][j] = true
	g.matrix[j][i] = true
}

// AddEdge adds a directed edge between the given nodes.
func (g *Graph) AddEdgeDirected(i, j int) {
	g.matrix[i][j] = true
}

// HasEdge returns true if there is an edge between the given nodes.
func (g *Graph) HasEdge(i, j int) bool {
	if !g.ValidNode(i) || !g.ValidNode(j) {
		return false
	}
	return g.matrix[i][j]
}

// ValidNode returns true if i is the ID of a node in the graph.
func (g *Graph) ValidNode(i int) bool {
	return i >= 0 && i < len(g.nodes)
}
